package repository

import (
	"tiktok/app/config"
	"tiktok/app/schema"
)

func CountLikedByVideoId(videoId int64) int64 {
	var count int64
	config.Db.Table("likes").Where("video_id = ?", videoId).Count(&count)
	return count
}

func GetLikeListByUserId(userId int64) []schema.Like {
	var likeList []schema.Like
	config.Db.Table("likes").Where("user_id = ?", userId).Find(&likeList)
	return likeList
}
