package model

type Player struct {
	PlayerID 		   string  `json:"playerId"`
	PerformanceRating  float64 `json:"performanceRating"`
	OldMMR 			   int     `json:"oldMmr"`
}

type MatchResult struct {
	Team1Players []Player `json:"team1Players"`
	Team2Players []Player `json:"team2Players"`
	Team1Rounds  int            `json:"team1Rounds"`
	Team2Rounds  int            `json:"team2Rounds"`
}

type PlayerMMRChange struct {
	PlayerID  string `json:"playerId"`
	MMRChange int    `json:"mmrChange"`
}
