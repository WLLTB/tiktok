package utils

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"tiktok/app/vo"
)

func ErrorHandler(c *gin.Context, errorMessage string) {
	log.Println(errorMessage)
	c.JSON(http.StatusOK, vo.Response{
		StatusCode: 1,
		StatusMsg:  errorMessage,
	})
}
