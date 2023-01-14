package service

import (
	"tiktok/app/repository"
	"tiktok/app/vo"
)

func GetAuthorInfo(userId int, authorId int) (vo.User, error) {
	author := repository.GetUserById(authorId)
	authorFollowedCount := repository.GetFollowedCount(authorId)
	authorFollowCount := repository.GetFollowCount(authorId)

	authorInfo := vo.User{
		Id:            authorId,
		Name:          author.Username,
		FollowCount:   authorFollowCount,
		FollowerCount: authorFollowedCount,
		IsFollow:      repository.CheckIsFollowed(userId, authorId),
	}
	return authorInfo, nil
}
