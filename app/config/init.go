package config

import (
	"fmt"
	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	. "tiktok/app/utils"
)

var Db *gorm.DB

func InitGormDb() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", MYSQL_USERNAME, MYSQL_PASSWORD, MYSQL_HOST, MYSQL_PORT, MYSQL_DBNAME)
	Db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	log.Print("数据库连接成功")
}

func InitLog() {
	log.SetPrefix("[GIN_LOG] ")
}

// InitClient 初始化 Redis 客户端
func InitClient() {
	Client = redis.NewClient(&redis.Options{
		Addr:     REDIS_ADDRESS,
		Password: REDIS_PASSWORD,
		DB:       REDIS_DB,
	})
	_, err := Client.Ping().Result()
	if err != nil {
		log.Panic("连接 Redis 失败: %w", err)
	}
	log.Print("连接 Redis 成功")
}
