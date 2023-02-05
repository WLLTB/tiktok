package repository

import (
	"app/config"
	"app/constant"
	. "app/model/schema"
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

func InsertComment(comment Comment) int64 {
	config.Db.Table(constant.COMMENT).Create(&comment)
	return comment.Id
}

func DeleteComment(comment Comment) {
	config.Db.Table(constant.COMMENT).Where(comment).Delete(&comment).First(&comment)
}
