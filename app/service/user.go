package service

import (
	"tiktok/app/repository"
	"tiktok/app/vo"
)

func supplementTargetUserInfo(currentUserId int64, targetUserId int64) vo.User {
	targetUser := repository.GetUserById(targetUserId)

	return vo.User{
		Id:            targetUserId,
		Name:          targetUser.Username,
		FollowCount:   repository.CountFollowByUserId(targetUserId),
		FollowerCount: repository.CountFollowedByFollowId(targetUserId),
		IsFollow:      repository.CheckIsFollowed(currentUserId, targetUserId),
	}
}
