package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	. "tiktok/app"
	. "tiktok/app/config"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	r := gin.Default()
	InitLog()
	InitRedisClient()
	InitGormDb()
	InitOssClient()
	InitRabbitMQ()
	InitRouter(r)
}
