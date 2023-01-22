package repository

import (
	"tiktok/app/config"
	"tiktok/app/constant"
	. "tiktok/app/schema"
)

func CountCommentByVideoId(videoId int64) int64 {
	var count int64
	config.Db.Table(constant.COMMENT).Where("video_id = ?", videoId).Count(&count)
	return count
}

func GetCommentList(videoId int64) []Comment {
	var comments []Comment
	config.Db.Table(constant.COMMENT).Where("video_id = ?", videoId).Find(&comments)
	return comments
}
