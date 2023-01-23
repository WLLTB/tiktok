package service

import (
	. "tiktok/app/repository"
	"tiktok/app/schema"
	"tiktok/app/vo"
)

// SupplementTargetUserInfo currentUserId 为 0 代表这是一个没登录的用户
func SupplementTargetUserInfo(currentUserId int64, targetUserId int64) vo.User {
	targetUser := GetUserById(targetUserId)

	return vo.User{
		Id:            targetUserId,
		Name:          targetUser.Username,
		FollowCount:   CountFollowByUserId(targetUserId),
		FollowerCount: CountFollowedByFollowId(targetUserId),
		IsFollow:      currentUserId != 0 && CheckIsFollowed(currentUserId, targetUserId),
	}
}

func HandlerRegister(username string, password string) {
	// 判断当前是否有这个用户了

	// 密码加密

	var user schema.User = schema.User{
		Username: username,
		Password: password,
	}
	InsertUser(user)
}

func HandlerLogin(username string, password string) (bool, int64) {
	// 密码加密

	user, err := GetUserByUsernameAndPassword(username, password)
	if err != nil {
		return false, 0
	}
	return true, user.Id
}
