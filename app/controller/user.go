package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktok/app/service"
	"tiktok/app/utils"
	. "tiktok/app/vo"
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

}

func Login(c *gin.Context) {

}

func UserInfo(c *gin.Context) {
	token := c.Query("token")
	tokenClaim, _ := utils.VerifyToken(token)
	currentUserId := tokenClaim["userId"].(string)
	targetUserId := c.Query("user_id")

	currentUserIdInt, _ := strconv.ParseInt(currentUserId, 10, 64)
	targetUserIdInt, _ := strconv.ParseInt(targetUserId, 10, 64)
	fmt.Println(currentUserIdInt)
	fmt.Println(targetUserIdInt)

	//userInfo := service.SupplementTargetUserInfo(currentUserIdInt, targetUserIdInt)
	userInfo := service.SupplementTargetUserInfo(1, 1)
	c.JSON(http.StatusOK, UserInfoResponse{
		Response: Response{StatusCode: 0},
		User:     userInfo,
	})
}
