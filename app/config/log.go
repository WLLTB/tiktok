package config

import (
	"log"
	"time"
)

func InitLog() {
	log.SetFlags(log.Ltime)
	log.SetPrefix("[GIN_LOG] " + time.Now().Format("2006/01/02 - 15:04:05") + " | ")
}
