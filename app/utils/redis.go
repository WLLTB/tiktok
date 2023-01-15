package utils

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

// GetValue 获取 Redis 中的值
func GetValue(key string) (string, error) {
	value, err := RedisClient.Get(key).Result()
	if err != nil {
		return "", fmt.Errorf("获取 key %s 失败: %w", key, err)
	}
	return value, nil
}

// SetValue 设置 Redis 中的值
func SetValue(key string, value string, expire time.Duration) error {
	err := RedisClient.Set(key, value, expire).Err()
	if err != nil {
		return fmt.Errorf("保存 key %s 失败: %w", key, err)
	}
	return nil
}
