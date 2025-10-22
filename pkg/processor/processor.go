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

	team1Average := calculateTeamAverageMMR(match.Team1Players)
	team2Average := calculateTeamAverageMMR(match.Team2Players)

	expectedScoreTeam1 := mp.ratingCalculator.CalculateExpectedScore(team1Average, team2Average)
	expectedScoreTeam2 := mp.ratingCalculator.CalculateExpectedScore(team2Average, team1Average)

	roundModifier := mp.ratingCalculator.CalculateRoundModifier(match.Team1Rounds, match.Team2Rounds)

	didTeam1Win := match.Team1Rounds > match.Team2Rounds
	didTeam2Win := match.Team2Rounds > match.Team1Rounds

	results = append(results, mp.processTeamMMRChanges(match.Team1Players, expectedScoreTeam1, roundModifier, didTeam1Win)...)
	results = append(results, mp.processTeamMMRChanges(match.Team2Players, expectedScoreTeam2, roundModifier, didTeam2Win)...)

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
