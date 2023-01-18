package service

import (
	"tiktok/app/repository"
	"tiktok/app/vo"
)

// SupplementTargetUserInfo currentUserId 为 0 代表这是一个没登录的用户
func SupplementTargetUserInfo(currentUserId int64, targetUserId int64) vo.User {
	targetUser := repository.GetUserById(targetUserId)

	return vo.User{
		Id:            targetUserId,
		Name:          targetUser.Username,
		FollowCount:   repository.CountFollowByUserId(targetUserId),
		FollowerCount: repository.CountFollowedByFollowId(targetUserId),
		IsFollow:      currentUserId != 0 && repository.CheckIsFollowed(currentUserId, targetUserId),
	}
}
