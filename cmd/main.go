package main

import (
	"fmt"

	"github.com/luizcordista/elo-system/pkg/model"
	"github.com/luizcordista/elo-system/pkg/processor"
	"github.com/luizcordista/elo-system/pkg/rating"
)

func main() {
	ratingCalculator := rating.NewEloRatingCalculator()
	matchProcessor := processor.NewMatchProcessor(ratingCalculator)

	match := model.MatchResult{
		Team1Players: []model.Player{
			{PlayerID: "player1", PerformanceRating: 1.0, OldMMR: 1500},
			{PlayerID: "player2", PerformanceRating: 1.0, OldMMR: 1500},
		},
		Team2Players: []model.Player{
			{PlayerID: "player3", PerformanceRating: 1.0, OldMMR: 1200},
			{PlayerID: "player4", PerformanceRating: 1.3, OldMMR: 1500},
		},
		Team1Rounds: 13,
		Team2Rounds: 11,
	}

	results := matchProcessor.ProcessMatchResult(match)

	for _, r := range results {
		fmt.Printf("Player %s MMR Change: %d\n", r.PlayerID, r.MMRChange)
	}
}
