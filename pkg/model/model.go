package model

type Player struct {
	PlayerID 		   string  `json:"playerId"`
	PerformanceRating  float64 `json:"performanceRating"`
	OldMMR 			   int     `json:"oldMmr"`
}

type MatchResult struct {
	TeamAPlayers []Player       `json:"team1Players"`
	TeamBPlayers []Player       `json:"team2Players"`
	TeamARounds  int            `json:"team1Rounds"`
	TeamBRounds  int            `json:"team2Rounds"`
}

type PlayerMMRChange struct {
	PlayerID  string `json:"playerId"`
	MMRChange int    `json:"mmrChange"`
	OldMMR	  int    `json:"oldMmr"`
	NewMMR    int    `json:"newMmr"`
}
