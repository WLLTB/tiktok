package repository

import (
	"tiktok/app/config"
)

func GetFollowCount(userId int) int {
	var count int64
	config.Db.Table("follows").Where("user_id = ?", userId).Count(&count)
	return int(count)
}

func GetFollowedCount(followId int) int {
	var count int64
	config.Db.Table("follows").Where("follower_id = ?", followId).Count(&count)
	return int(count)
}

func CheckIsFollowed(userId int, followId int) bool {
	var count int64
	config.Db.Table("follows").Where("user_id = ? and follower_id = ?", userId, followId).Count(&count)
	return count != 0
}
