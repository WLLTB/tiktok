package controller

import (
	"app/constant"
	"app/model/vo"
	"app/service"
	"app/utils"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"time"
)

type FeedResponse struct {
	vo.Response
	VideoList []vo.Video `json:"video_list,omitempty"`
	NextTime  int64      `json:"next_time,omitempty"`
}

// Feed 处理视频流
func Feed(ctx context.Context, c *app.RequestContext) {
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
		Response:  vo.Response{StatusCode: 0},
		VideoList: videoList,
		NextTime:  time.Now().Unix(),
	})
}
