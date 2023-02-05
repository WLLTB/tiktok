package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
)

func Register(r *server.Hertz) {

	GeneratedRegister(r)

	customizedRegister(r)
}
