package repository

import (
	"tiktok/app/config"
	"tiktok/app/schema"
)

func GetLikedCountByVideoId(videoId int) int {
	var count int64
	config.Db.Table("likes").Where("video_id = ?", videoId).Count(&count)
	return int(count)
}

func GetLikeListByUserId(userId int) []schema.Like {
	var likeList []schema.Like
	config.Db.Table("likes").Where("user_id = ?", userId).Find(&likeList)
	return likeList
}
