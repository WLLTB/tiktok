package repository

import (
	"fmt"
	"tiktok/app/config"
	. "tiktok/app/schema"
	"time"
)

func PageVideoListByTime(lastTime time.Time, count int64) []Video {
	var videos []Video
	config.Db.Table("videos").Where("publish_time < ?", lastTime).Order("publish_time desc").Limit(int(count)).Find(&videos)
	return videos
}

func InsertVideo(video Video) error {
	err := config.Db.Create(&video).Error
	if err != nil {
		return fmt.Errorf("插入 video 失败: %w", err)
	}
	return nil
}
