package answers

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Miss-you/mbti-cli/internal/questionbank"
	"github.com/stretchr/testify/require"
)

func TestValidateAcceptsCompleteStrictAnswerSet(t *testing.T) {
	err := Validate(answerValidationBank(), Set{
		"q01": "A",
		"q02": "D",
		"q03": "B",
	})

	require.NoError(t, err)
}

func TestValidateRejectsUnknownInvalidAndMissingAnswers(t *testing.T) {
	answerSet, err := Parse([]byte(`{
		"answers": {
			"q01": "z",
			"unknown": "A"
		}
	}`))
	require.NoError(t, err)

	err = Validate(answerValidationBank(), answerSet)

	require.Error(t, err)
	require.ErrorContains(t, err, "answer validation failed")
	require.ErrorContains(t, err, "unknown question id unknown")
	require.ErrorContains(t, err, "question q01 option Z is not valid")
	require.ErrorContains(t, err, "missing answer for question q02")
	require.ErrorContains(t, err, "missing answer for question q03")
}

func TestValidateRejectsInvalidAnswerFixture(t *testing.T) {
	data, err := os.ReadFile(filepath.Join("testdata", "invalid-answers.json"))
	require.NoError(t, err)
	answerSet, err := Parse(data)
	require.NoError(t, err)

	err = Validate(answerValidationBank(), answerSet)

	require.Error(t, err)
	require.ErrorContains(t, err, "unknown question id unknown")
	require.ErrorContains(t, err, "question q01 option Z is not valid")
	require.ErrorContains(t, err, "missing answer for question q02")
	require.ErrorContains(t, err, "missing answer for question q03")
}

func TestValidateReportsDeterministicAggregatedIssues(t *testing.T) {
	bank := answerValidationBank()
	answerSet := Set{
		"q03":   "x",
		"later": "A",
		"alpha": "B",
	}
	wantIssues := []string{
		"unknown question id alpha",
		"unknown question id later",
		"question q03 option x is not valid",
		"missing answer for question q01",
		"missing answer for question q02",
	}

	for range 5 {
		err := Validate(bank, answerSet)

		require.Error(t, err)
		var validationErr *ValidationError
		require.ErrorAs(t, err, &validationErr)
		require.Equal(t, wantIssues, validationErr.Issues)
	}
}

func answerValidationBank() questionbank.Bank {
	return questionbank.Bank{
		Questions: []questionbank.Question{
			answerValidationQuestion("q01"),
			answerValidationQuestion("q02"),
			answerValidationQuestion("q03"),
		},
	}
}

func answerValidationQuestion(id string) questionbank.Question {
	return questionbank.Question{
		ID: id,
		Options: []questionbank.Option{
			{Code: "A"},
			{Code: "B"},
			{Code: "C"},
			{Code: "D"},
		},
	}
}
