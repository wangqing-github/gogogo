package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

var client *redis.Client
var ctx = context.Background()

func InitRedis() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		PoolSize: 20,
	})
	// 测试连接
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("无法连接到Redis: %v", err)
	}
	fmt.Println("redis ping，连接redis成功:", pong)
}

func Keys() []string {
	keys, err := client.Keys(ctx, "*").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(keys)
	return keys
}

func SetValue(key string, value any) {
	fmt.Println("client信息为：", client)
	err := client.Set(ctx, key, value, 0).Err()
	if err != nil {
		log.Fatalf("redis设置值失败：key:%s,value:%s", key, value)
	}
}

func GetByKey(key string) string {
	result, err := client.Get(ctx, key).Result()
	if err != nil {
		log.Fatalln("redis获取值失败：key", key)
	}
	return result
}

func SAdd(key string, value string) {
	err := client.SAdd(ctx, key, value).Err()
	if err != nil {
		log.Fatalf("Failed to set slice in Redis: %v", err)
	}
}

func SMembers(key string) (result []string) {
	result, err := client.SMembers(ctx, key).Result()
	if err != nil {
		log.Fatalf("Failed to set slice in Redis: %v", err)
	}
	return
}

func ZAdd(key string, member string, value float64) {
	err := client.ZAdd(ctx, key, &redis.Z{
		Score:  value,
		Member: member,
	}).Err()
	if err != nil {
		log.Fatalf("Failed to set slice in Redis: %v", err)
	}
}

func ZScore(key string, member string) {
	err := client.ZScore(ctx, key, member).Err()
	if err != nil {
		log.Fatalf("Failed to set slice in Redis: %v", err)
	}
}
