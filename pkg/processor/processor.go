package processor

import (
	"github.com/luizcordista/elo-system/pkg/model"
	"github.com/luizcordista/elo-system/pkg/rating"
)

func ProcessMatchResult(match model.MatchResult) []model.PlayerMMRChange {
	results := make([]model.PlayerMMRChange, 0)

	team1Average := calculateTeamAverageMMR(match.Team1Players)
	team2Average := calculateTeamAverageMMR(match.Team2Players)

	expectedScoreTeam1 := rating.CalculateExpectedScore(team1Average, team2Average)
	expectedScoreTeam2 := rating.CalculateExpectedScore(team2Average, team1Average)

	roundModifier := rating.CalculateRoundModifier(match.Team1Rounds, match.Team2Rounds)

	didTeam1Win := match.Team1Rounds > match.Team2Rounds
	didTeam2Win := match.Team2Rounds > match.Team1Rounds

	results = append(results, processTeamMMRChanges(match.Team1Players, expectedScoreTeam1, roundModifier, didTeam1Win)...)
	results = append(results, processTeamMMRChanges(match.Team2Players, expectedScoreTeam2, roundModifier, didTeam2Win)...)

	return results
}

func processTeamMMRChanges(players []model.Player, expectedScore float64, roundModifier float64, didWin bool) []model.PlayerMMRChange {
	results := make([]model.PlayerMMRChange, 0)

	for _, p := range players {
		ind := rating.CalculateIndividualPerformance(p.PerformanceRating)
		change := rating.CalculateMMRChange(
			rating.DefaultKFactor,
			expectedScore,
			roundModifier,
			ind,
			didWin,
		)
		results = append(results, model.PlayerMMRChange{
			PlayerID:  p.PlayerID,
			MMRChange: change,
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
