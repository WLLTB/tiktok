package repository

import (
	"tiktok/app/config"
)

func CountFollowByUserId(userId int64) int64 {
	var count int64
	config.Db.Table("follows").Where("user_id = ?", userId).Count(&count)
	return count
}

func CountFollowedByFollowId(followId int64) int64 {
	var count int64
	config.Db.Table("follows").Where("follower_id = ?", followId).Count(&count)
	return count
}

func CheckIsFollowed(userId int64, followId int64) bool {
	var count int64
	config.Db.Table("follows").Where("user_id = ? and follower_id = ?", userId, followId).Count(&count)
	return count != 0
}
