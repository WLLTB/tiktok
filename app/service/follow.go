package service

import (
	"app/constant"
	"app/model/vo"
	. "app/repository"
	"app/utils"
	"fmt"
	"log"
	"strconv"
	"sync"
)

// FollowUser 关注用户
func FollowUser(userId int64, toUserId int64) error {
	if userId == toUserId {
		return fmt.Errorf("can not follow yourself")
	}
	if err := InsertFollow(userId, toUserId); err != nil {
		log.Printf("userId：%d follow toUserId：%d insert fail,err：%w", userId, toUserId, err)
		return err
	}
	go func() {
		err := utils.PushSet(constant.RedisSetFollowPrefix+strconv.FormatInt(userId, 10), strconv.FormatInt(toUserId, 10))
		if err != nil {
			log.Printf("userId: %d follow toUserId: %d redis关注缓存更新失败 \n", userId, toUserId)
		}
		return
	}()
	return nil
}

// CancelFollowUser 取消关注
func CancelFollowUser(userId int64, toUserId int64) error {
	if DeleteFollow(userId, toUserId) > 0 {
		go func() {
			err := utils.DeleteSet(constant.RedisSetFollowPrefix+strconv.FormatInt(userId, 10), strconv.FormatInt(toUserId, 10))
			if err != nil {
				log.Printf("userId: %d follow toUserId: %d redis取消关注缓存更新失败 \n", userId, toUserId)
			}
			return
		}()
	}
	return nil
}

// IsFollowed 查看是否关注该用户，先查redis再查数据库
func IsFollowed(userId int64, followedUserId int64) bool {
	if utils.IsSetMember(constant.RedisSetFollowPrefix+strconv.FormatInt(userId, 10), strconv.FormatInt(followedUserId, 10)) {
		return true
	}
	if CheckIsFollowed(userId, followedUserId) {
		log.Printf("用户关注缓存不一致 userId：%d followedUserId：%d\n", userId, followedUserId)
		return true
	}
	return false
}

// GetUserFollowList 获取userId用户的关注列表
func GetUserFollowList(curUserId int64, tarUserId int64) ([]vo.User, error) {
	userIdList, _ := utils.GetSet(constant.RedisSetFollowPrefix + strconv.FormatInt(tarUserId, 10))
	length := len(userIdList)
	var wg sync.WaitGroup
	wg.Add(length)

	userChan := make(chan vo.User, length)

	for _, id := range userIdList {
		id := id
		go func() {
			defer wg.Done()
			idInt, _ := strconv.ParseInt(id, 10, 64)
			user := SupplementTargetUserInfo(curUserId, idInt)
			userChan <- user
		}()
	}
	wg.Wait()
	close(userChan)

	return userChanToUserList(userChan), nil

}

func GetUserFansList(curUserId int64, tarUserId int64) ([]vo.User, error) {
	followList := GetUserFans(tarUserId)
	length := len(followList)
	var wg sync.WaitGroup
	wg.Add(length)
	userChan := make(chan vo.User, length)
	for _, follow := range followList {
		id := follow.FollowerId
		go func() {
			defer wg.Done()
			user := SupplementTargetUserInfo(curUserId, id)
			userChan <- user
		}()
	}

	wg.Wait()
	close(userChan)

	return userChanToUserList(userChan), nil

}

func GetUserFriendList(curUserId int64) ([]vo.User, error) {
	userFollowIdList, _ := utils.GetSet(constant.RedisSetFollowPrefix + strconv.FormatInt(curUserId, 10))
	length := len(userFollowIdList)
	var wg sync.WaitGroup
	wg.Add(length)

	userChan := make(chan vo.User, length)

	for _, id := range userFollowIdList {
		id := id
		go func() {
			defer wg.Done()
			idInt, _ := strconv.ParseInt(id, 10, 64)
			if IsFollowed(curUserId, idInt) {
				user := SupplementTargetUserInfo(curUserId, idInt)
				userChan <- user
			}
		}()
	}

	wg.Wait()
	close(userChan)

	return userChanToUserList(userChan), nil
}

func userChanToUserList(userChan chan vo.User) []vo.User {
	var userList []vo.User
	for user := range userChan {
		userList = append(userList, user)
	}
	return userList
}
