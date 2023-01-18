package repository

import (
	"fmt"
	"tiktok/app/config"
	"tiktok/app/constant"
	. "tiktok/app/schema"
	"time"
)

func PageVideoListByTime(lastTime time.Time, count int64) []Video {
	var videos []Video
	config.Db.Table(constant.VIDEO).Where("publish_time < ?", lastTime).Order("publish_time desc").Limit(int(count)).Find(&videos)
	return videos
}

func InsertVideo(video Video) error {
	err := config.Db.Table(constant.VIDEO).Create(&video).Error
	if err != nil {
		return fmt.Errorf("插入 video 失败: %w", err)
	}
	return nil
}

func CountVideoById(videoId int64) int64 {
	var count int64
	config.Db.Table(constant.VIDEO).Where("id = ?", videoId).Count(&count)
	return count
}
