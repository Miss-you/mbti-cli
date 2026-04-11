package scoring

import (
	"testing"

	"github.com/Miss-you/mbti-cli/internal/answers"
	"github.com/Miss-you/mbti-cli/internal/questionbank"
	"github.com/stretchr/testify/require"
)

func TestScoreAggregatesSignedOptionScoresByDimension(t *testing.T) {
	result, err := Score(scoringBank(), answers.Set{
		"q01": "A",
		"q02": "D",
		"q03": "B",
		"q04": "C",
	})

	require.NoError(t, err)
	require.Equal(t, 4, result.Answered)
	require.Equal(t, 4, result.Total)
	require.Equal(t, map[questionbank.Dimension]int{
		questionbank.DimensionEI: 0,
		questionbank.DimensionSN: -1,
		questionbank.DimensionTF: 1,
		questionbank.DimensionJP: 0,
	}, result.DimensionScores)
}

func TestScoreRejectsInvalidStrictAnswersBeforeAggregation(t *testing.T) {
	result, err := Score(scoringBank(), answers.Set{
		"q01": "A",
		"q02": "D",
		"q03": "B",
	})

	require.Error(t, err)
	require.ErrorContains(t, err, "missing answer for question q04")
	require.Zero(t, result)
}

func scoringBank() questionbank.Bank {
	return questionbank.Bank{
		Questions: []questionbank.Question{
			scoringQuestion("q01", questionbank.DimensionEI, false),
			scoringQuestion("q02", questionbank.DimensionEI, true),
			scoringQuestion("q03", questionbank.DimensionSN, false),
			scoringQuestion("q04", questionbank.DimensionTF, false),
		},
	}
}

func scoringQuestion(id string, dimension questionbank.Dimension, reverse bool) questionbank.Question {
	return questionbank.Question{
		ID:        id,
		Dimension: dimension,
		Reverse:   reverse,
		Options: []questionbank.Option{
			{Code: "A", Score: 2},
			{Code: "B", Score: -1},
			{Code: "C", Score: 1},
			{Code: "D", Score: -2},
		},
	}
}
