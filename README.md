# LeaderboardGo
 Simple leaderboard implementation using Golang and Redis

 Application required to have Redis running on localhost:6379, this can be changed in redis/RedisConnect.go.

## Install

    git clone https://github.com/babell00/LeaderboardGo.git
    cd LeaderboardGo
    go get
    go run main.go

## Add player's score to leaderboard
 You can add player's score to leaderbaord by making POST request to http://localhost:8080/playerscore/ 
 
    curl -H "Content-Type: application.json" -X POST -d '{"game_name": "my_super_game", "player_name": "tom", "score": 50}' http://localhost:8080/playerscore

## Getting game's leadeboard(top 10)
 To get leaderboard you need to make GET request to http://localhost:8080/leaderboard/{game_name}
  
    curl -H "Content-Type: application.json" -X GET http://localhost:8080/leaderboard/my_super_game
