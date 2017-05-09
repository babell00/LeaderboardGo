package rediss

import (
	"sync"
	"github.com/babell00/LeaderboardGo/leaderboard"
	"encoding/json"
	"log"
)

var (
	_service     *RedisService
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

func (service *RedisService) AddToSortedSet(key string, score float64, playerScore leaderboard.PlayerScore) {
	playerScoreJson, err := json.Marshal(playerScore)
	if err != nil {
		log.Println(err)
		return
	}
	service.redisClient.sortedSetAdd(key, score, playerScoreJson)
}

func (service *RedisService) GetTop10Player(key string) []leaderboard.PlayerScore {
	results := service.redisClient.sortedSetRangeByScoreDesc(key, "0", "1000", 0, 10)
	return convertResultArrayToPlayerScoreArray(results)
}

func convertResultArrayToPlayerScoreArray(results []string) []leaderboard.PlayerScore{
	playerScoreArray := make([]leaderboard.PlayerScore, len(results), cap(results))
	for k, v := range results {
		var playerScore leaderboard.PlayerScore
		err := json.Unmarshal([]byte(v), &playerScore)
		if err != nil {
			log.Println(err)
			continue
		}
		playerScoreArray[k] = playerScore
	}
	return playerScoreArray
}
