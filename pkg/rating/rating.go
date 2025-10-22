package rating

import "math"

const (
	DefaultKFactor = 40.0
	MaxRating      = 1.3
	MinRating      = 0.7
	MaxRounds      = 13.0
)

type EloRatingCalculator struct{}

func NewEloRatingCalculator() RatingCalculator {
	return &EloRatingCalculator{}
}

func (e *EloRatingCalculator) CalculateExpectedScore(teamAverageMMR, opponentAverageMMR int) float64 {
	diff := float64(opponentAverageMMR - teamAverageMMR)
	return 1.0 / (1.0 + math.Pow(10.0, diff/400.0))
}

func (e *EloRatingCalculator) CalculateRoundModifier(teamARounds, teamBRounds int) float64 {
	roundDifference := math.Abs(float64(teamARounds - teamBRounds))
	return 1.0 + (0.5 * (roundDifference / MaxRounds))
}

func (e *EloRatingCalculator) CalculateIndividualPerformance(playerPerformanceRating float64) float64 {
	if playerPerformanceRating > MaxRating {
		return MaxRating
	}
	if playerPerformanceRating < MinRating {
		return MinRating
	}
	return playerPerformanceRating
}

func (e *EloRatingCalculator) CalculateMMRChange(kFactor float64, expectedScore, roundModifier, individualPerformance float64, didWin bool) int {
	var mmrChange float64
	if didWin {
		delta := 1.0 - expectedScore
		mmrChange = kFactor * delta * roundModifier * individualPerformance
	} else {
		delta := 0.0 - expectedScore
		mmrChange = kFactor * delta * roundModifier / individualPerformance
	}
	return int(math.Round(mmrChange))
}
