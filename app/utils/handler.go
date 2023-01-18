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
	c.Abort()
}

func SuccessHandler(c *gin.Context, successMessage string) {
	log.Println(successMessage)
	c.JSON(http.StatusOK, vo.Response{
		StatusCode: 0,
		StatusMsg:  successMessage,
	})
	c.Abort()
}
