package service

import (
	"tiktok/app/repository"
	"tiktok/app/vo"
)

func supplementAuthorInfo(userId int64, authorId int64) vo.User {
	author := repository.GetUserById(authorId)

	return vo.User{
		Id:            authorId,
		Name:          author.Username,
		FollowCount:   repository.CountFollowByUserId(authorId),
		FollowerCount: repository.CountFollowedByFollowId(authorId),
		IsFollow:      repository.CheckIsFollowed(userId, authorId),
	}
}
