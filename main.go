package main

import (
	"github.com/gin-gonic/gin"
	"tiktok/app"
	"tiktok/app/config"
	"tiktok/app/utils"
)

func main() {
	r := gin.Default()
	config.InitLog()
	utils.InitClient()
	config.InitGormDb()
	app.InitRouter(r)
}
