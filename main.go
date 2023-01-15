package main

import (
	"github.com/gin-gonic/gin"
	"tiktok/app"
	"tiktok/app/config"
)

func main() {
	r := gin.Default()
	config.InitLog()
	config.InitClient()
	config.InitGormDb()
	app.InitRouter(r)
}
