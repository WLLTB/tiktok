package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktok/app/config"
	"tiktok/app/schema"
	"tiktok/app/utils"
	. "tiktok/app/vo"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	videoId := c.PostForm("video_id")
	videoIdInt, _ := strconv.ParseInt(videoId, 10, 64)
	actionType := c.PostForm("action_type")
	actionTypeInt, _ := strconv.ParseInt(actionType, 10, 64)
	token := c.PostForm("token")
	userId, _ := utils.VerifyToken(token)

	// 判断video是否存在
	videoOne := schema.Video{}
	config.Db.Model(&schema.Video{}).Select("Id").Where("Id = ?", videoIdInt).First(&videoOne)

	if videoOne.Id == 0 {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "Video is not exist",
		})
	}
	if actionTypeInt == 1 {
		likeOne := schema.Like{UserId: userId, VideoId: videoIdInt}
		config.Db.Model(&schema.Like{}).Create(&likeOne)
		c.JSON(http.StatusOK, Response{
			StatusCode: 0,
			StatusMsg:  "Like successful",
		})
	}
	if actionTypeInt == 2 {
		likeOne := schema.Like{UserId: userId, VideoId: videoIdInt}
		config.Db.Model(&schema.Like{}).Where(likeOne).Delete(&likeOne)
		c.JSON(http.StatusOK, Response{
			StatusCode: 0,
			StatusMsg:  "Unlike successful",
		})
	}
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
