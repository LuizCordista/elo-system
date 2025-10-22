package rating

type RatingCalculator interface {
	CalculateExpectedScore(teamAverageMMR, opponentAverageMMR int) float64
	CalculateRoundModifier(team1rounds, team2rounds int) float64
	CalculateIndividualPerformance(playerPerformanceRating float64) float64
	CalculateMMRChange(kFactor float64, expectedScore, roundModifier, individualPerformance float64, didWin bool) int
}
