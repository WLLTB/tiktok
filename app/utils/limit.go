package utils

import (
	"context"
	"github.com/go-redis/redis/v8"
	. "tiktok/app/constant"
	"time"
)

// TokenBucketLimiter 令牌桶限流器
type TokenBucketLimiter struct {
	capacity int           // 容量
	rate     int           // 发放令牌速率/秒
	client   *redis.Client // Redis客户端
	script   *redis.Script // TryAcquire脚本
}

func NewTokenBucketLimiter(client *redis.Client, capacity, rate int) *TokenBucketLimiter {
	return &TokenBucketLimiter{
		capacity: capacity,
		rate:     rate,
		client:   client,
		script:   redis.NewScript(TokenBucketLimiterTryAcquireRedisScript),
	}
}

func (l *TokenBucketLimiter) TryAcquire(ctx context.Context, resource string) error {
	// 当前时间
	now := time.Now().Unix()
	success, err := l.script.Run(ctx, l.client, []string{resource}, l.capacity, l.rate, now).Bool()
	if err != nil {
		return err
	}
	// 若到达窗口请求上限，请求失败
	if !success {
		return ErrAcquireFailed
	}
	return nil
}
