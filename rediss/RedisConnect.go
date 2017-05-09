package rediss

import (
	"github.com/go-redis/redis"
	"sync"
)

var (
	redisClient     *RedisClient
	redisClientOnce sync.Once
)

type RedisClient struct {
	redisClient *redis.Client
}

func Connect() *RedisClient {
	redisClientOnce.Do(func() {
		client := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})
		redisClient = &RedisClient{client}
	})
	return redisClient
}

func (rClient *RedisClient) GetClient() *redis.Client {
	return rClient.redisClient
}