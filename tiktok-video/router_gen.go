package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	router "tiktok-video/biz/router"
)

func register(r *server.Hertz) {

	router.GeneratedRegister(r)

	customizedRegister(r)
}
