package repository

import (
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

func GetLikeVideoList(userId int64) []Video {
	var videoIds []int
	config.Db.Table(constant.LIKE).Where("user_id = ?", userId).Pluck(constant.VideoID, &videoIds)

	var videos []Video
	config.Db.Model(&Video{}).Where("id in (?)", videoIds).Find(&videos)
	return videos
}

func InsertVideo(video Video) error {
	err := config.Db.Table(constant.VIDEO).Create(&video)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func CountVideoById(videoId int64) int64 {
	var count int64
	config.Db.Table(constant.VIDEO).Where("id = ?", videoId).Count(&count)
	return count
}

func GetVideoById(videoId int64) (Video, error) {
	var video Video
	err := config.Db.Table(constant.VIDEO).Where("id = ?", videoId).First(&video)
	if err.Error != nil {
		return Video{}, err.Error
	}
	return video, nil
}

func GetPublishVideoList(userId int64) []Video {
	var videos []Video
	config.Db.Table(constant.VIDEO).Where("user_id = ?", userId).Find(&videos)
	return videos
}
