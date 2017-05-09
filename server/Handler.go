package server

import (
	"net/http"
	"encoding/json"
	"github.com/babell00/LeaderboardGo/rediss"
	"github.com/babell00/LeaderboardGo/leaderboard"
	"github.com/gorilla/mux"
	"log"
)

func GetLeaderboard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	pathParam := vars["game_name"]

	redisService := rediss.GetRedisService()

	results := redisService.GetTop10Player(pathParam)

	json.NewEncoder(w).Encode(&results)
}

func AddPlayerScore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)

	var playerScore leaderboard.PlayerScore

	err := decoder.Decode(&playerScore)
	if err != nil {
		log.Println(err)
		return
	}
	redisService := rediss.GetRedisService()
	redisService.AddToSortedSet(playerScore.GameName, playerScore.Score, playerScore)

	json.NewEncoder(w).Encode(&playerScore)
}
