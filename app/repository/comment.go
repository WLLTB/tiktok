package repository

import (
	"tiktok/app/config"
)

func CountCommentByVideoId(videoId int64) int64 {
	var count int64
	config.Db.Table("comments").Where("video_id = ?", videoId).Count(&count)
	return count
}
