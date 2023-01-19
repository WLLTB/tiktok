package repository

import (
	"fmt"
	"tiktok/app/config"
	"tiktok/app/schema"
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

func InsertFollow(userId int64, toUserId int64) error {
	follow := schema.Follow{
		UserId:     userId,
		FollowerId: toUserId,
	}
	result := config.Db.Table("follows").Create(&follow)
	if result.RowsAffected < 1 {
		return fmt.Errorf("插入 follow 失败: %w", result.Error)
	}
	return nil
}

func DeleteFollow(userId int64, toUserId int64) int64 {
	result := config.Db.Table("follows").Where("user_id = ?", userId).Where("follow_id = ?", toUserId).Delete(schema.Follow{})
	return result.RowsAffected
}

func GetUserFans(userId int64) []schema.Follow {
	var followList []schema.Follow
	config.Db.Table("follows").Where("follow_id = ?", userId).Find(&followList)
	return followList
}
