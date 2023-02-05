package router

import (
	. "app/constant"
	"app/controller"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func customizedRegister(r *server.Hertz) {
	r.GET(FeedPath, controller.Feed)
}
