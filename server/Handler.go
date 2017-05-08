package server

import (
	"net/http"
	"log"
	"encoding/json"
)

func Leaderboard(w http.ResponseWriter, r *http.Request) {
	log.Println("Servving Leaderboard")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	test := "test"

	json.NewEncoder(w).Encode(&test)
}