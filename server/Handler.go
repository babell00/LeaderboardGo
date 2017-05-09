package server

import (
	"net/http"
	"log"
	"encoding/json"
	"github.com/babell00/LeaderboardGo/leaderboard"
	"github.com/babell00/LeaderboardGo/rediss"
	"github.com/go-redis/redis"
)

func ShowLeaderboard(w http.ResponseWriter, r *http.Request) {
	log.Println("Servving Leaderboard")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	rd := rediss.Connect()
	client := rd.GetClient()

	tt := redis.ZRangeBy{"0", "10", 0 , 6}

	results , _ := client.ZRevRangeByScore("sTest1", tt).Result()

	json.NewEncoder(w).Encode(&results)
}

func AddPlayerScore(w http.ResponseWriter, r *http.Request) {
	log.Println("Servving Leaderboard")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)

	var playerScore leaderboard.PlayerScore

	err := decoder.Decode(&playerScore)
	if err != nil {
		log.Println(err)
		return
	}

	rd := rediss.Connect()
	client := rd.GetClient()

	z := redis.Z{playerScore.Score, playerScore.PlayerName}

	client.ZAdd("sTest1", z)

	json.NewEncoder(w).Encode(&playerScore)
}
