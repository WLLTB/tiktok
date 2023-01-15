package config

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	. "tiktok/app/constant"
	. "tiktok/app/utils"
)

var Db *gorm.DB

func InitGormDb() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", MysqlUsername, MysqlPassword, MysqlHost, MysqlPort, MysqlDbname)
	Db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	log.Print("数据库连接成功")
}

func InitLog() {
	log.SetPrefix("[GIN_LOG] ")
}

// InitRedisClient 初始化 Redis 客户端
func InitRedisClient() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     RedisAddress,
		Password: RedisPassword,
		DB:       RedisDb,
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		log.Panic("连接 Redis 失败: %w", err)
	}
	log.Print("连接 Redis 成功")
}

// InitOssClient 初始化 Oss 客户端
func InitOssClient() {
	var err error
	OssClient, err = oss.New(OssEndpoint, OssAccessKeyId, OssAccessKeySecret)
	if err != nil {
		log.Panic("创建 OSS 客户端失败: %w", err)
	}
}
