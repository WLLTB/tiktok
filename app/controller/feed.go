package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok/app/constant"
	"tiktok/app/service"
	"tiktok/app/utils"
	. "tiktok/app/vo"
	"time"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

// Feed 处理视频流
func Feed(c *gin.Context) {
	lastTime := c.Query(constant.LastTime)
	token := c.Query(constant.TOKEN)
	userId, err := utils.VerifyToken(token)
	// 因为没强制登录，所以非法token就当没登录，给 0
	if err != nil {
		userId = 0
	}
	videoList, err := service.SupplementFeedVideoList(userId, lastTime, constant.VideoCount)
	if err != nil {
		utils.ErrorHandler(c, constant.ServerError)
		return
	}

	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: videoList,
		NextTime:  time.Now().Unix(),
	})
}
