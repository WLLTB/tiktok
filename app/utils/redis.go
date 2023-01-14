package utils

import (
	"fmt"
	"github.com/go-redis/redis"
)

var client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

func GetValue(key string) string {
	value, err := client.Get(key).Result()
	if err != nil {
		fmt.Println(err)
	}
	return value
}

func SetValue(key string, value string) {
	err := client.Set(key, value, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
}
