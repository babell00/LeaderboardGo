package rediss

import (
	"github.com/go-redis/redis"
	"sync"
	"log"
	"os"
)

var (
	_client     *redisClient
	_redisAddress string
	_clientOnce sync.Once
)

type redisClient struct {
	redisClient *redis.Client
}

func redisConnect() *redisClient {
	_clientOnce.Do(func() {
		_redisAddress = os.Getenv("REDIS_ADDRESS")
		if _redisAddress == "" {
			_redisAddress = "localhost:6379"
		}
		log.Printf("Redis addess: %v", _redisAddress)
		rClient := redis.NewClient(&redis.Options{
			Addr:     _redisAddress,
			Password: "",
			DB:       0,
		})
		_client = &redisClient{rClient}
	})
	return _client
}

func (client *redisClient) sortedSetAdd(key string, score float64, object interface{}) {
	z := redis.Z{score, object}
	err := client.redisClient.ZAdd(key, z).Err()
	if err != nil {
		log.Println(err)
	}
}

func (client *redisClient) sortedSetRangeByScoreDesc(key string, min string, max string, offset int64, count int64) []string {
	zRangeBy := redis.ZRangeBy{min, max, offset, count}
	results, err := client.redisClient.ZRevRangeByScore(key, zRangeBy).Result()
	if err != nil {
		log.Println(err)
		return nil
	}
	return results

}
