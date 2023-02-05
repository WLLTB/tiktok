package main

import (
	. "app/config"
	"app/constant"
	"app/router"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	InitLog()
	InitRedisClient()
	InitGormDb()
	InitOssClient()
	InitRabbitMQ()

	h := server.Default(server.WithHostPorts(constant.PORT))
	router.Register(h)
	h.Spin()
}
