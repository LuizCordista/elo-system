package processor

import (
	"testing"

	"github.com/luizcordista/elo-system/pkg/model"
	"github.com/luizcordista/elo-system/pkg/rating"
	"github.com/stretchr/testify/assert"
)

func TestMatchProcessorImpl_ProcessMatchResult_5v5(t *testing.T) {
	calculator := rating.NewEloRatingCalculator()
	processor := NewMatchProcessor(calculator)

	match := model.MatchResult{
		// Median MMR of Team 1 is 1000
		TeamAPlayers: []model.Player{
			{PlayerID: "1", OldMMR: 1000, PerformanceRating: 1.2},
			{PlayerID: "2", OldMMR: 1100, PerformanceRating: 1.1},
			{PlayerID: "3", OldMMR: 950, PerformanceRating: 1.3},
			{PlayerID: "4", OldMMR: 1050, PerformanceRating: 1.0},
			{PlayerID: "5", OldMMR: 900, PerformanceRating: 0.9},
		},
		// Median MMR of Team 2 is 1100
		TeamBPlayers: []model.Player{
			{PlayerID: "6", OldMMR: 1200, PerformanceRating: 0.8},
			{PlayerID: "7", OldMMR: 1150, PerformanceRating: 0.9},
			{PlayerID: "8", OldMMR: 1000, PerformanceRating: 0.7},
			{PlayerID: "9", OldMMR: 1100, PerformanceRating: 1.0},
			{PlayerID: "10", OldMMR: 1050, PerformanceRating: 1.1},
		},
		TeamARounds: 13,
		TeamBRounds: 7,
	}

	expectedChanges := []model.PlayerMMRChange{
		{PlayerID: "1", MMRChange: 38, OldMMR: 1000, NewMMR: 1038},
		{PlayerID: "2", MMRChange: 35, OldMMR: 1100, NewMMR: 1135},
		{PlayerID: "3", MMRChange: 41, OldMMR: 950, NewMMR: 991},
		{PlayerID: "4", MMRChange: 32, OldMMR: 1050, NewMMR: 1082},
		{PlayerID: "5", MMRChange: 28, OldMMR: 900, NewMMR: 928},
		{PlayerID: "6", MMRChange: -39, OldMMR: 1200, NewMMR: 1161},
		{PlayerID: "7", MMRChange: -35, OldMMR: 1150, NewMMR: 1115},
		{PlayerID: "8", MMRChange: -45, OldMMR: 1000, NewMMR: 955},
		{PlayerID: "9", MMRChange: -32, OldMMR: 1100, NewMMR: 1068},
		{PlayerID: "10", MMRChange: -29, OldMMR: 1050, NewMMR: 1021},
	}

	changes := processor.ProcessMatchResult(match)

	assert.ElementsMatch(t, expectedChanges, changes)
}
