package main

import (
	"github.com/gin-gonic/gin"
	config "tiktok/app/config"
)

func main() {
	r := gin.Default()

	config.InitGormDb()
	config.InitRouter(r)

	r.Run(":9999")
}
