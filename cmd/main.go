package main

import (
	"fmt"

	"github.com/luizcordista/elo-system/pkg/model"
	"github.com/luizcordista/elo-system/pkg/processor"
)

func main() {
	match := model.MatchResult{
		Team1Players: []model.Player{
			{PlayerID: "player1", PerformanceRating: 1.0, OldMMR: 1500},
			{PlayerID: "player2", PerformanceRating: 1.0, OldMMR: 1500},
		},
		Team2Players: []model.Player{
			{PlayerID: "player3", PerformanceRating: 1.0, OldMMR: 1500},
			{PlayerID: "player4", PerformanceRating: 1.0, OldMMR: 1500},
		},
		Team1Rounds: 8,
		Team2Rounds: 5,
	}

	results := processor.ProcessMatchResult(match)

	for _, r := range results {
		fmt.Printf("Player %s MMR Change: %d\n", r.PlayerID, r.MMRChange)
	}
}
