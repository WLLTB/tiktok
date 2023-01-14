package config

import (
	"log"
)

func InitLog() {
	log.SetPrefix("[GIN_LOG] ")
}
