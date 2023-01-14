package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok/app/schema"
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
	var rawVideos []schema.Video
	var videos []Video
	lastTime := c.Query("latest_time")
	rawVideos, err := service.GetVideos(lastTime, count)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, rawVideo := range rawVideos {
		video := Video{
			Id:            rawVideo.Id,
			Author:        DemoUser,
			CommentCount:  1,
			CoverUrl:      rawVideo.CoverUrl,
			IsFavorite:    true,
			FavoriteCount: 1,
		}
		videos = append(videos, video)
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: videos,
		NextTime:  time.Now().Unix(),
	})
}
