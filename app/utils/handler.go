package utils

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"tiktok/app/vo"
)

func ErrorHandler(c *gin.Context, err error) {
	log.Println(err)
	c.JSON(http.StatusOK, vo.Response{
		StatusCode: 1,
		StatusMsg:  err.Error(),
	})
}
