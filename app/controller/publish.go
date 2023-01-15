package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok/app/utils"
	. "tiktok/app/vo"
	"time"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

// Publish 将用户投稿的视频上传到阿里云 OSS，将对应的地址 URL 通过消息队列存储到数据库中，加快响应速度
func Publish(c *gin.Context) {
	token := c.PostForm("token")

	if _, exist := usersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	userId := "sas"

	file, err := c.FormFile("data")
	url, err := utils.OssUpload(file, time.Now().Format("2006-01-02 15:04:05")+"_"+userId)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	// 队列更新数据库

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  url + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
