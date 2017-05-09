package rediss

import (
	"github.com/go-redis/redis"
	"sync"
	"log"
)

var (
	_client     *redisClient
	_clientOnce sync.Once
)

type redisClient struct {
	_redisClient *redis.Client
}

func redisConnect() *redisClient {
	_clientOnce.Do(func() {
		rClient := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})
		_client = &redisClient{rClient}
	})
	return _client
}

func (client *redisClient) sortedSetAdd(key string, score float64, object interface{}) {
	z := redis.Z{score, object}
	err := client._redisClient.ZAdd(key, z).Err()
	if err != nil {
		log.Println(err)
	}
}

func (client *redisClient) sortedSetRangeByScoreDesc(key string, min string, max string, offset int64, count int64) []string {
	zRangeBy := redis.ZRangeBy{min, max, offset, count}
	results, err := client._redisClient.ZRevRangeByScore(key, zRangeBy).Result()
	if err != nil {
		log.Println(err)
		return nil
	}
	return results

}
