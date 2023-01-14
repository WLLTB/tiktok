package utils

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
)

var client *redis.Client

// InitClient 初始化 Redis 客户端
func InitClient() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := client.Ping().Result()
	if err != nil {
		log.Panic("连接 Redis 失败: %w", err)
	}
	log.Print("连接 Redis 成功")
}

// GetValue 获取 Redis 中的值
func GetValue(key string) (string, error) {
	value, err := client.Get(key).Result()
	if err != nil {
		return "", fmt.Errorf("获取 key %s 失败: %w", key, err)
	}
	return value, nil
}

// SetValue 设置 Redis 中的值
func SetValue(key string, value string, expire time.Duration) error {
	err := client.Set(key, value, expire).Err()
	if err != nil {
		return fmt.Errorf("保存 key %s 失败: %w", key, err)
	}
	return nil
}
