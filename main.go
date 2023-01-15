package main

import (
	"github.com/gin-gonic/gin"
	. "tiktok/app"
	. "tiktok/app/config"
)

func main() {
	r := gin.Default()
	InitLog()
	InitRedisClient()
	InitGormDb()
	InitOssClient()
	InitRabbitMQ()
	InitRouter(r)
}
