package rediss

import "sync"

var (
	_service *RedisService
	_serviceOnce sync.Once
)

type RedisService struct {
	redisClient *redisClient
}

func GetRedisService() *RedisService {
	_serviceOnce.Do(func() {
		redisClient := redisConnect()
		_service = &RedisService{redisClient}
	})
	return _service
}

func (service *RedisService) AddToSortedSet(key string, score float64, object interface{}){
	service.redisClient.sortedSetAdd(key, score, object)
}

func (service *RedisService) GetTop10Player(key string) []string {
	return service.redisClient.sortedSetRangeByScoreDesc(key, "0", "1000", 0 , 10)
}