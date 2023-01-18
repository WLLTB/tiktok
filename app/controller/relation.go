package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktok/app/constant"
	"tiktok/app/service"
	"tiktok/app/utils"
	. "tiktok/app/vo"
)

type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	// 获取query
	token := c.Query("token")
	toUserId := c.Query("to_user_id")
	actionType := c.Query("action_type")

	// 参数校验
	toUserIdInt, err := strconv.ParseInt(toUserId, 10, 64)
	if err != nil {
		utils.ErrorHandler(c, "query参数 to_user_id 错误")
	}

	actionTypeInt, err := strconv.ParseInt(actionType, 10, 16)
	if err != nil || (actionTypeInt != 1 && actionTypeInt != 2) {
		utils.ErrorHandler(c, "query参数 action_type 错误")
	}

	userId, _ := utils.VerifyToken(token)

	// 判断关注or取关
	if actionTypeInt == 1 {
		// 关注用户
		err := service.FollowUser(userId, toUserIdInt)
		if err != nil {
			utils.ErrorHandler(c, err.Error())
		} else {
			utils.SuccessHandler(c, constant.FollowSuccess)
		}
	} else if actionTypeInt == 2 {
		// 取消关注
		_ = service.CancelFollowUser(userId, toUserIdInt)
		utils.SuccessHandler(c, constant.CancelFollowSuccess)
	}

	//if _, exist := usersLoginInfo[token]; exist {
	//	c.JSON(http.StatusOK, Response{StatusCode: 0})
	//} else {
	//	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	//}
}

// FollowList 获取某个用户的关注列表
func FollowList(c *gin.Context) {
	// 参数获取&校验
	token := c.Query("token")
	_, _ = utils.VerifyToken(token)
	userId := c.Query("user_id")
	userIdInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		utils.ErrorHandler(c, "query参数 user_id 错误")
	}

	// 获取用户列表
	userList, err := service.GetUserFollowList(userIdInt)
	if err != nil {
		utils.ErrorHandler(c, constant.ServerError)
	}

	// 返回集合
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		UserList: userList,
	})

	//c.JSON(http.StatusOK, UserListResponse{
	//	Response: Response{
	//		StatusCode: 0,
	//	},
	//	UserList: []User{DemoUser},
	//})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: []User{DemoUser},
	})
}

// FriendList all users have same friend list
func FriendList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: []User{DemoUser},
	})
}
