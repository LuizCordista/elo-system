package processor

import (
	"github.com/luizcordista/elo-system/pkg/model"
	"github.com/luizcordista/elo-system/pkg/rating"
)

type MatchProcessorImpl struct {
	ratingCalculator rating.RatingCalculator
}

func NewMatchProcessor(ratingCalculator rating.RatingCalculator) MatchProcessor {
	return &MatchProcessorImpl{
		ratingCalculator: ratingCalculator,
	}
}

func (mp *MatchProcessorImpl) ProcessMatchResult(match model.MatchResult) []model.PlayerMMRChange {
	results := make([]model.PlayerMMRChange, 0)

	teamAAverage := calculateTeamAverageMMR(match.TeamAPlayers)
	teamBAverage := calculateTeamAverageMMR(match.TeamBPlayers)

	expectedScoreTeamA := mp.ratingCalculator.CalculateExpectedScore(teamAAverage, teamBAverage)
	expectedScoreTeamB := mp.ratingCalculator.CalculateExpectedScore(teamBAverage, teamAAverage)

	roundModifier := mp.ratingCalculator.CalculateRoundModifier(match.TeamARounds, match.TeamBRounds)

	didTeamAWin := match.TeamARounds > match.TeamBRounds
	didTeamBWin := match.TeamBRounds > match.TeamARounds

	results = append(results, mp.processTeamMMRChanges(match.TeamAPlayers, expectedScoreTeamA, roundModifier, didTeamAWin)...)
	results = append(results, mp.processTeamMMRChanges(match.TeamBPlayers, expectedScoreTeamB, roundModifier, didTeamBWin)...)

	return results
}

func (mp *MatchProcessorImpl) processTeamMMRChanges(players []model.Player, expectedScore float64, roundModifier float64, didWin bool) []model.PlayerMMRChange {
	results := make([]model.PlayerMMRChange, 0)

	for _, p := range players {
		ind := mp.ratingCalculator.CalculateIndividualPerformance(p.PerformanceRating)
		change := mp.ratingCalculator.CalculateMMRChange(
			rating.DefaultKFactor,
			expectedScore,
			roundModifier,
			ind,
			didWin,
		)
		results = append(results, model.PlayerMMRChange{
			PlayerID:  p.PlayerID,
			MMRChange: change,
			OldMMR: p.OldMMR,
			NewMMR: p.OldMMR + change,
		})
	}

	return results
}

func calculateTeamAverageMMR(players []model.Player) int {
	total := 0
	for _, p := range players {
		total += p.OldMMR
	}
	return total / len(players)
}
