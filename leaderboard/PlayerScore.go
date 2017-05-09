package leaderboard

type PlayerScore struct {
	GameName string `json:"game_name"`
	PlayerName string `json:"player_name"`
	Score float64	`json:"score"`
}
