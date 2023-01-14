package repository

import (
	"tiktok/app/config"
)

func GetCommentCountByVideoId(videoId int) int {
	var count int64
	config.Db.Table("comments").Where("video_id = ?", videoId).Count(&count)
	return int(count)
}
