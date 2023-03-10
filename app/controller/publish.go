package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktok/app/constant"
	"tiktok/app/repository"
	"tiktok/app/schema"
	"tiktok/app/service"
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
	token := c.PostForm(constant.TOKEN)
	userId, _ := utils.VerifyToken(token)

	title := c.PostForm(constant.TITLE)
	// 考虑限制上传时间间隔
	file, err := c.FormFile(constant.DATA)
	if err != nil {
		utils.ErrorHandler(c, constant.ServerError)
		return
	}

	fileName := time.Now().Format("2006-01-02 15:04:05") + "_" + strconv.FormatInt(userId, 10) + constant.VideoFormat

	playUrl, err := utils.OssUpload(file, fileName)
	coverUrl := playUrl + constant.CoverSuffix

	if err != nil {
		utils.ErrorHandler(c, constant.ServerError)
		return
	}

	// 需要重构成队列更新数据库
	_ = repository.InsertVideo(schema.Video{
		AuthorId:    userId,
		PlayUrl:     playUrl,
		CoverUrl:    coverUrl,
		PublishTime: time.Now(),
		Title:       title,
	})
	utils.SuccessHandler(c, constant.ActionSuccess)
}

// PublishList 用户投稿列表
func PublishList(c *gin.Context) {
	userId, _ := strconv.ParseInt(c.Query(constant.UserId), 10, 64)
	token := c.Query(constant.TOKEN)
	currentUserId, _ := utils.VerifyToken(token)

	videoList, err := service.SupplementPublishVideoList(userId, currentUserId)
	if err != nil {
		utils.ErrorHandler(c, constant.ServerError)
		return
	}

	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: videoList,
	})
}
