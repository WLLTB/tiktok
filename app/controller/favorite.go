package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktok/app/constant"
	"tiktok/app/service"
	"tiktok/app/utils"
	. "tiktok/app/vo"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	videoId, err := strconv.ParseInt(c.Query(constant.VideoID), 10, 64)
	if err != nil {
		utils.ErrorHandler(c, constant.InvalidVideoID)
		return
	}

	actionType, err := strconv.ParseInt(c.Query(constant.ActionType), 10, 64)
	if err != nil {
		utils.ErrorHandler(c, constant.InvalidActionType)
		return
	}

	token := c.Query(constant.TOKEN)
	userId, _ := utils.VerifyToken(token)

	service.HandlerFavoriteAction(userId, videoId, actionType, c)
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
