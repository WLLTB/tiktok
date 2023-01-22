package repository

import (
	"tiktok/app/config"
	"tiktok/app/constant"
	"tiktok/app/schema"
)

func CountLikedByVideoId(videoId int64) int64 {
	var count int64
	config.Db.Table(constant.LIKE).Where("video_id = ?", videoId).Count(&count)
	return count
}

func GetLikeListByUserId(userId int64) []schema.Like {
	var likeList []schema.Like
	config.Db.Table(constant.LIKE).Where("user_id = ?", userId).Find(&likeList)
	return likeList
}

func CreateLike(like *schema.Like) {
	config.Db.Table(constant.LIKE).Create(like)
}

func DeleteLike(like *schema.Like) {
	config.Db.Table(constant.LIKE).Where(like).Delete(like)
}
