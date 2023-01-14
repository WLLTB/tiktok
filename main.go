package main

import (
	"github.com/gin-gonic/gin"
	config "tiktok/app/config"
	"tiktok/app/utils"
)

func main() {
	r := gin.Default()
	config.InitLog()
	utils.InitClient()
	config.InitGormDb()
	config.InitRouter(r)
}
