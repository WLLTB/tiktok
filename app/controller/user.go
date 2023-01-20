package controller

import (
	"net/http"
	"strconv"
	"tiktok/app/constant"
	"tiktok/app/service"
	"tiktok/app/utils"
	. "tiktok/app/vo"

	"github.com/gin-gonic/gin"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

var userIdSequence int

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserInfoResponse struct {
	Response
	User User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query(constant.USERNAME)
	password := c.Query(constant.PASSWORD)
	service.HandlerRegister(username, password)
}

func Login(c *gin.Context) {
	username := c.Query(constant.USERNAME)
	password := c.Query(constant.PASSWORD)

	hasUser, userId := service.HandlerLogin(username, password)
	if !hasUser {
		utils.ErrorHandler(c, "Username Or Password Error")
		return
	}

	token, _ := utils.GenerateToken(userId)
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: 0},
		UserId:   1,
		Token:    token,
	})
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")
	userId, _ := utils.VerifyToken(token)
	targetUserId := c.Query("user_id")

	// 类型转换
	targetUserIdInt, _ := strconv.ParseInt(targetUserId, 10, 64)

	userInfo := service.SupplementTargetUserInfo(userId, targetUserIdInt)
	c.JSON(http.StatusOK, UserInfoResponse{
		Response: Response{StatusCode: 0},
		User:     userInfo,
	})
}
