package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktok/app/constant"
	"tiktok/app/repository"
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
	// 处理重复请求、是否有此记录、封装友好提示
	videoId, _ := strconv.ParseInt(c.Query(constant.VideoID), 10, 64)
	actionType, _ := strconv.ParseInt(c.Query(constant.ActionType), 10, 64)
	commentId, _ := strconv.ParseInt(c.Query(constant.CommentId), 10, 64)
	commentText := c.Query(constant.CommentText)
	userId, _ := utils.VerifyToken(c.Query(constant.TOKEN))

	// 判断video是否存在
	currentVideo, err := repository.GetVideoById(videoId)
	if err != nil || currentVideo.Id == 0 {
		utils.ErrorHandler(c, constant.VideoNotExist)
		return
	}

	var comment schema.Comment
	commentActionResponse := CommentActionResponse{
		Response: Response{StatusCode: 0, StatusMsg: constant.ActionSuccess},
	}

	switch actionType {
	case 1:
		comment = schema.Comment{UserId: userId, VideoId: currentVideo.Id, CommentText: commentText, CreateDate: time.Now()}
		commentActionResponse.Comment = HandlerCommentAction(actionType, comment, userId, currentVideo.AuthorId)
		c.JSON(http.StatusOK, commentActionResponse)
	case 2:
		comment = schema.Comment{Id: commentId}
		HandlerCommentAction(actionType, comment, userId, currentVideo.AuthorId)
		c.JSON(http.StatusOK, commentActionResponse)
	default:
		utils.ErrorHandler(c, constant.InvalidActionType)
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	token := c.Query(constant.TOKEN)
	userId, _ := utils.VerifyToken(token)
	videoIdInt, _ := strconv.ParseInt(c.Query(constant.VideoID), 10, 64)

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
