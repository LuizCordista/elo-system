package rating

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const floatDelta = 1e-9

func TestCalculateExpectedScore_EqualRatingsReturnsPointFive(t *testing.T) {
	ratingCalculator := NewEloRatingCalculator()
	exp := ratingCalculator.CalculateExpectedScore(1500, 1500)
	require.InDelta(t, 0.5, exp, floatDelta, "expected 0.5 for equal ratings")
}

func TestCalculateExpectedScore_HigherRatingReturnsMoreThanPointFive(t *testing.T) {
	ratingCalculator := NewEloRatingCalculator()
	exp := ratingCalculator.CalculateExpectedScore(1600, 1500)
	require.Greater(t, exp, 0.5, "expected > 0.5 for higher rating")
}

func TestCalculateExpectedScore_LowerRatingReturnsLessThanPointFive(t *testing.T) {
	ratingCalculator := NewEloRatingCalculator()
	exp := ratingCalculator.CalculateExpectedScore(1400, 1500)
	require.Less(t, exp, 0.5, "expected < 0.5 for lower rating")
}

func TestCalculatedRoundModifier_ReturnsCorrectValues(t *testing.T) {
	ratingCalculator := NewEloRatingCalculator()
	mod := ratingCalculator.CalculateRoundModifier(10, 10)
	require.InDelta(t, 1.0, mod, floatDelta, "expected modifier 1.0 when rounds are equal")

	expected := 1.1538461538
	mod = ratingCalculator.CalculateRoundModifier(13, 9)
	require.InDelta(t, expected, mod, floatDelta, "unexpected modifier for 13 vs 11")

	maxExpected := 1.5
	mod = ratingCalculator.CalculateRoundModifier(13, 0)
	require.InDelta(t, maxExpected, mod, floatDelta, "unexpected modifier for max round difference")
}

func TestCalculatingIndividualPerfomance_ReturnsCorrectValues(t *testing.T) {
	ratingCalculator := NewEloRatingCalculator()
	performanceInsideBounds := ratingCalculator.CalculateIndividualPerformance(1.1)
	require.InDelta(t, 1.1, performanceInsideBounds, floatDelta, "unexpected performance for score 1.1")

	performanceLowerThanBounds := ratingCalculator.CalculateIndividualPerformance(0.3)
	require.InDelta(t, 0.7, performanceLowerThanBounds, floatDelta, "unexpected performance for score 0.3")

	performanceHigherThanBounds := ratingCalculator.CalculateIndividualPerformance(1.7)
	require.InDelta(t, 1.3, performanceHigherThanBounds, floatDelta, "unexpected performance for score 1.7")
}
