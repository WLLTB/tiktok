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
	token := c.Query("token")
	toUserId := c.Query("to_user_id")
	actionType := c.Query("action_type")

	toUserIdInt, err := strconv.ParseInt(toUserId, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "query参数 to_user_id 错误",
		})
	}

	actionTypeInt, err := strconv.ParseInt(actionType, 10, 16)
	if err != nil || (actionTypeInt != 1 && actionTypeInt != 2) {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "query参数 action_type 错误",
		})
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  constant.InvalidMessage,
		})
	}

	if actionTypeInt == 1 {
		err := service.FollowUser(userId, toUserIdInt)
		if err != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, Response{
				StatusCode: 0,
				StatusMsg:  "follow success",
			})
		}
	} else if actionTypeInt == 2 {
		_ = service.CancelFollowUser(userId, toUserIdInt)
		c.JSON(http.StatusOK, Response{
			StatusCode: 0,
			StatusMsg:  "cancel follow success",
		})
	}

	//if _, exist := usersLoginInfo[token]; exist {
	//	c.JSON(http.StatusOK, Response{StatusCode: 0})
	//} else {
	//	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	//}
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: []User{DemoUser},
	})
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
