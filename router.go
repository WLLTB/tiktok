package main

import (
	"github.com/gin-gonic/gin"
	"tiktok/controller"
)

func initRouter(r *gin.Engine) {
	apiRouter := r.Group("/douyin")

	apiRouter.GET("/feed/", controller.Feed)
}
