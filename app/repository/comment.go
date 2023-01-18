package repository

import (
	"tiktok/app/config"
	. "tiktok/app/schema"
)

func CountCommentByVideoId(videoId int64) int64 {
	var count int64
	config.Db.Table("comments").Where("video_id = ?", videoId).Count(&count)
	return count
}

func GetCommentList(videoId int64) []Comment {
	var comments []Comment
	config.Db.Model(&Comment{}).Where("video_id = ?", videoId).Find(&comments)
	return comments
}
