package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	handler "tiktok-video/biz/handler"
)

func customizedRegister(r *server.Hertz) {
	r.GET("/ping", handler.Ping)
}
