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

// PushSet 向 Set 中添加元素
func PushSet(key string, value string) error {
	err := RedisClient.SAdd(key, value).Err()
	if err != nil {
		return fmt.Errorf("在 %s 集合中保存 %s 失败: %w", key, value, err)
	}
	return nil
}

// GetSet 获取 Set 中的所有元素
func GetSet(key string) ([]string, error) {
	var set []string
	set, err := RedisClient.SMembers(key).Result()
	if err != nil {
		return set, fmt.Errorf("获取 %s 集合失败: %w", key, err)
	}
	return set, nil
}

// HasSetValue 判断set中是否有某个值
func HasSetValue(key, value string) (bool, error) {
	has, err := RedisClient.SIsMember(key, value).Result()
	if err != nil {
		return false, fmt.Errorf("判断 key %s 中是否有 %s 失败: %w", key, value, err)
	}
	return has, nil
}
