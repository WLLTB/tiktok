package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktok/app/config"
	"tiktok/app/constant"
	"tiktok/app/schema"
	. "tiktok/app/service"
	"tiktok/app/utils"
	. "tiktok/app/vo"
	"time"
)

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	videoId := c.PostForm("video_id")
	videoIdInt, _ := strconv.ParseInt(videoId, 10, 64)
	actionType := c.PostForm("action_type")
	actionTypeInt, _ := strconv.ParseInt(actionType, 10, 64)
	commentText := c.PostForm("comment_text")
	commentId := c.PostForm("comment_id")
	commentIdInt, _ := strconv.ParseInt(commentId, 10, 64)
	token := c.PostForm("token")
	userId, _ := utils.VerifyToken(token)

	// 判断video是否存在
	videoOne := schema.Video{}
	config.Db.Model(&schema.Video{}).Select("Id").Where("Id = ?", videoIdInt).First(&videoOne)

	if videoOne.Id == 0 {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Video is not exist"})
	} else {
		if actionTypeInt == 1 {
			commentOne := schema.Comment{UserId: userId, VideoId: videoIdInt, CommentText: commentText, CreateDate: time.Now()}
			config.Db.Model(&schema.Comment{}).Create(&commentOne)

			c.JSON(http.StatusOK, CommentActionResponse{
				Response: Response{StatusCode: 0, StatusMsg: "Comment successful"},
				Comment: Comment{
					Id:         commentOne.Id,
					User:       SupplementTargetUserInfo(userId, userId),
					Content:    commentText,
					CreateDate: commentOne.CreateDate.Format("2006-01-02 15:04:05"),
				},
			})
		}
		if actionTypeInt == 2 {
			commentOne := schema.Comment{UserId: userId, VideoId: videoIdInt, CommentText: commentText, Id: commentIdInt}
			config.Db.Model(&schema.Comment{}).Where(commentOne).Delete(&commentOne).First(&commentOne)
			c.JSON(http.StatusOK, CommentActionResponse{
				Response: Response{StatusCode: 0, StatusMsg: "Delete successful"},
				Comment: Comment{
					Id:         commentOne.Id,
					User:       SupplementTargetUserInfo(userId, userId),
					Content:    commentText,
					CreateDate: commentOne.CreateDate.Format("2006-01-02 15:04:05"),
				},
			})
		}
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	token := c.Query("token")
	userId, _ := utils.VerifyToken(token)
	videoId := c.Query("video_id")
	videoIdInt, _ := strconv.ParseInt(videoId, 10, 64)

	commentList, err := SupplementCommentList(userId, videoIdInt)
	if err != nil {
		utils.ErrorHandler(c, constant.ServerError)
		return
	}

	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Response{StatusCode: 0},
		CommentList: commentList,
	})
}
