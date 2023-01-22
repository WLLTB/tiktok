package repository

import (
	"fmt"
	"tiktok/app/config"
	. "tiktok/app/constant"
	"tiktok/app/schema"
)

func CountFollowByUserId(userId int64) int64 {
	var count int64
	config.Db.Table(FOLLOW).Where("user_id = ?", userId).Count(&count)
	return count
}

func CountFollowedByFollowId(followId int64) int64 {
	var count int64
	config.Db.Table(FOLLOW).Where("follower_id = ?", followId).Count(&count)
	return count
}

func CheckIsFollowed(userId int64, followId int64) bool {
	var count int64
	config.Db.Table(FOLLOW).Where("user_id = ? and follower_id = ?", userId, followId).Count(&count)
	return count != 0
}

func InsertFollow(userId int64, toUserId int64) error {
	follow := schema.Follow{
		UserId:     userId,
		FollowerId: toUserId,
	}
	result := config.Db.Table(FOLLOW).Create(&follow)
	if result.RowsAffected < 1 {
		return fmt.Errorf("插入 follow 失败: %w", result.Error)
	}
	return nil
}

func DeleteFollow(userId int64, toUserId int64) int64 {
	result := config.Db.Table(FOLLOW).Where("user_id = ? and follower_id = ?", userId, toUserId).Delete(schema.Follow{})
	return result.RowsAffected
}

func GetUserFans(userId int64) []schema.Follow {
	var followList []schema.Follow
	config.Db.Table(FOLLOW).Where("follower_id = ?", userId).Find(&followList)
	return followList
}
