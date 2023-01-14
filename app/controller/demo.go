package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok/app/repository"
	"tiktok/app/schema"
)

type demo struct {
	Response
	UserList []schema.User `json:"user_id"`
}

func GetTableUserList(c *gin.Context) {
	tableUsers, _ := repository.GetTableUserList()
	c.JSON(http.StatusOK, demo{
		Response: Response{StatusCode: 0},
		UserList: tableUsers,
	})
}
