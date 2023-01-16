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
	lastTime := c.Query("latest_time")
	videoList, err := service.SupplementVideoList(1, lastTime, constant.VideoCount)

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
