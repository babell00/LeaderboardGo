package main

import (
	"log"
	"github.com/babell00/LeaderboardGo/server"
)

func main() {
	log.Println("Starting application")

	server.NewServer(8080)
}