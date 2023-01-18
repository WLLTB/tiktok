package service

import (
	"log"
	"strconv"
	"tiktok/app/constant"
	"tiktok/app/repository"
	"tiktok/app/utils"
)

func FollowUser(userId int64, toUserId int64) error {
	if err := repository.InsertFollow(userId, toUserId); err != nil {
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

func CancelFollowUser(userId int64, toUserId int64) error {
	if repository.DeleteFollow(userId, toUserId) > 0 {
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
