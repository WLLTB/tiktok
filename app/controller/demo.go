package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok/app/common"
	"tiktok/app/repository"
	"tiktok/app/schema"
)

type demo struct {
	common.Response
	UserList []schema.User `json:"user_id"`
}

func GetTableUserList(c *gin.Context) {
	tableUsers, _ := repository.GetTableUserList()
	c.JSON(http.StatusOK, demo{
		Response: common.Response{StatusCode: 0},
		UserList: tableUsers,
	})
}
