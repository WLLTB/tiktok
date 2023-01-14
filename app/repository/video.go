package repository

import (
	"tiktok/app/config"
	. "tiktok/app/schema"
	"time"
)

func PageVideoListByTime(lastTime time.Time, count int) []Video {
	var videos []Video
	config.Db.Table("videos").Where("publish_time < ?", lastTime).Order("publish_time desc").Limit(count).Find(&videos)
	return videos
}
