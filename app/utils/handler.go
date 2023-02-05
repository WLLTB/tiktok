package utils

import (
	"app/model/vo"
	"github.com/cloudwego/hertz/pkg/app"
	"log"
	"net/http"
)

func ErrorHandler(c *app.RequestContext, errorMessage string) {
	log.Println(errorMessage)
	c.JSON(http.StatusOK, vo.Response{
		StatusCode: 1,
		StatusMsg:  errorMessage,
	})
	c.Abort()
}

func SuccessHandler(c *app.RequestContext, successMessage string) {
	log.Println(successMessage)
	c.JSON(http.StatusOK, vo.Response{
		StatusCode: 0,
		StatusMsg:  successMessage,
	})
	c.Abort()
}
