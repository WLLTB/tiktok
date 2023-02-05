package service

import (
	"app/model/vo"
	. "app/repository"
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
