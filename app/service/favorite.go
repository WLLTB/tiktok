package service

import (
	"github.com/gin-gonic/gin"
	"tiktok/app/constant"
	"tiktok/app/repository"
	"tiktok/app/schema"
	"tiktok/app/utils"
)

func HandlerFavoriteAction(userId int64, videoId int64, actionType int64, c *gin.Context) {
	if !checkVideoExist(videoId) {
		utils.ErrorHandler(c, constant.VideoNotExist)
		return
	}

	like := schema.Like{UserId: userId, VideoId: videoId}
	switch actionType {
	case 1:
		repository.CreateLike(&like)
		utils.SuccessHandler(c, constant.LikeSuccess)
	case 2:
		repository.DeleteLike(&like)
		utils.SuccessHandler(c, constant.UnLikeSuccess)
	default:
		utils.SuccessHandler(c, constant.ActionFail)
	}
}

func checkVideoExist(videoId int64) bool {
	matchedVideoCount := repository.CountVideoById(videoId)
	return matchedVideoCount != 0
}
