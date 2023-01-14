package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok/app/service"
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
	const count = 10
	lastTime := c.Query("latest_time")
	videoList, err := service.ConvertVideoList(1, lastTime, count)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: videoList,
		NextTime:  time.Now().Unix(),
	})
}
