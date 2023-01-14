package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "tiktok/app/vo"
)

var tempChat = map[string][]Message{}

var messageIdSequence = int(1)

type ChatResponse struct {
	Response
	MessageList []Message `json:"message_list"`
}

// MessageAction no practical effect, just check if token is valid
func MessageAction(c *gin.Context) {
	//token := c.Query("token")
	//toUserId := c.Query("to_user_id")
	//content := c.Query("content")
	//
	//if user, exist := usersLoginInfo[token]; exist {
	//	userIdB, _ := strconv.Atoi(toUserId)
	//	chatKey := genChatKey(user.Id, userIdB)
	//
	//	//atomic.Addint(&messageIdSequence, 1)
	//	curMessage := Message{
	//		Id:         messageIdSequence,
	//		Content:    content,
	//		CreateTime: time.Now().Format(time.Kitchen),
	//	}
	//
	//	if messages, exist := tempChat[chatKey]; exist {
	//		tempChat[chatKey] = append(messages, curMessage)
	//	} else {
	//		tempChat[chatKey] = []Message{curMessage}
	//	}
	//	c.JSON(http.StatusOK, Response{StatusCode: 0})
	//} else {
	//	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	//}
}

// MessageChat all users have same follow list
func MessageChat(c *gin.Context) {
	//token := c.Query("token")
	//toUserId := c.Query("to_user_id")
	//
	//if user, exist := usersLoginInfo[token]; exist {
	//	userIdB, _ := strconv.Atoi(toUserId)
	//	chatKey := genChatKey(user.Id, int(userIdB))
	//
	//	c.JSON(http.StatusOK, ChatResponse{Response: Response{StatusCode: 0}, MessageList: tempChat[chatKey]})
	//} else {
	//	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	//}
}

func genChatKey(userIdA int, userIdB int) string {
	if userIdA > userIdB {
		return fmt.Sprintf("%d_%d", userIdB, userIdA)
	}
	return fmt.Sprintf("%d_%d", userIdA, userIdB)
}
