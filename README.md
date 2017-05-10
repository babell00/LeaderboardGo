# LeaderboardGo
 Simple leaderboard implementation using Golang and Redis

 Application required to have Redis running on localhost:6379, this can be changed in redis/RedisConnect.go.

## Install

    git clone https://github.com/babell00/LeaderboardGo.git
    cd LeaderboardGo
    go get
    go run main.go

## Add player's score to leaderboard
 You can add player's score to leaderbaord by making POST request to <br />
 http://localhost:8080/playerscore/ <br />
 
    curl -H "Content-Type: application.json" -X POST -d '{"game_name": "my_super_game", "player_name": "tom", "score": 50}' http://localhost:8080/playerscore

## Getting game's leadeboard
 To get leaderboard you need to make GET request to  <br />
 http://localhost:8080/leaderboard/{game_name}?size={how_many_top_player_you_want_to_get} <br />
 If size is not provided request will return top 10
  
    curl -H "Content-Type: application.json" -X GET http://localhost:8080/leaderboard/my_super_game?size=10

## Run with Docker
    docker build -t leaderboard .
    docker container run --name lb -p 8080:8080 -e REDIS_ADDRESS="192.168.2.11:6379" leaderbaord
 If you will not provide environment variable REDIS_ADDRESS, application will look for redis in localhost