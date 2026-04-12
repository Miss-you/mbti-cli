package questionbank

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateAcceptsCanonicalV3QuestionBank(t *testing.T) {
	bank := loadCanonicalV3Bank(t)

	require.NoError(t, Validate(bank))
}

func TestValidateAcceptsQuestionBankFixture(t *testing.T) {
	bank := loadQuestionBankFixture(t, "valid-bank.json")

	require.NoError(t, Validate(bank))
}

func TestValidateRejectsInvalidQuestionBankFixtures(t *testing.T) {
	tests := []struct {
		name string
		file string
		want string
	}{
		{
			name: "meta total mismatch",
			file: "invalid-meta-total.json",
			want: "meta.total",
		},
		{
			name: "duplicate question id",
			file: "invalid-duplicate-question-id.json",
			want: "duplicate question id q01",
		},
		{
			name: "threshold overlap",
			file: "invalid-threshold-overlap.json",
			want: "threshold moderate_a overlaps threshold slight_a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bank := loadQuestionBankFixture(t, tt.file)

			err := Validate(bank)

			require.Error(t, err)
			require.ErrorContains(t, err, tt.want)
		})
	}
}

func TestValidateRejectsInvalidQuestionBankSchemas(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Bank)
		want   string
	}{
		{
			name: "meta total mismatch",
			mutate: func(bank *Bank) {
				bank.Meta.Total = len(bank.Questions) + 1
			},
			want: "meta.total",
		},
		{
			name: "missing dimension metadata",
			mutate: func(bank *Bank) {
				delete(bank.Meta.Dimensions, DimensionEI)
			},
			want: "missing dimension metadata for EI",
		},
		{
			name: "unknown dimension metadata",
			mutate: func(bank *Bank) {
				bank.Meta.Dimensions[Dimension("XX")] = DimensionMeta{Count: 1}
			},
			want: "unknown dimension metadata XX",
		},
		{
			name: "dimension count mismatch",
			mutate: func(bank *Bank) {
				meta := bank.Meta.Dimensions[DimensionEI]
				meta.Count++
				bank.Meta.Dimensions[DimensionEI] = meta
			},
			want: "dimension EI count",
		},
		{
			name: "empty question id",
			mutate: func(bank *Bank) {
				bank.Questions[0].ID = ""
			},
			want: "questions[0].id is required",
		},
		{
			name: "duplicate question id",
			mutate: func(bank *Bank) {
				bank.Questions[1].ID = bank.Questions[0].ID
			},
			want: "duplicate question id q01",
		},
		{
			name: "unsupported question dimension",
			mutate: func(bank *Bank) {
				bank.Questions[0].Dimension = Dimension("XX")
			},
			want: "questions[0].dimension",
		},
		{
			name: "missing localized scenario text",
			mutate: func(bank *Bank) {
				bank.Questions[0].Scenario.ZH = ""
			},
			want: "questions[0].scenario.zh is required",
		},
		{
			name: "option count is not four",
			mutate: func(bank *Bank) {
				bank.Questions[0].Options = bank.Questions[0].Options[:3]
			},
			want: "questions[0].options must contain exactly 4 options",
		},
		{
			name: "duplicate option code",
			mutate: func(bank *Bank) {
				bank.Questions[0].Options[1].Code = "A"
			},
			want: "questions[0].options missing code B",
		},
		{
			name: "unknown option code",
			mutate: func(bank *Bank) {
				bank.Questions[0].Options[0].Code = "Z"
			},
			want: "questions[0].options[0].code",
		},
		{
			name: "invalid option score",
			mutate: func(bank *Bank) {
				bank.Questions[0].Options[0].Score = 0
			},
			want: "questions[0].options[0].score",
		},
		{
			name: "missing option label",
			mutate: func(bank *Bank) {
				bank.Questions[0].Options[0].Label.EN = ""
			},
			want: "questions[0].options[0].label.en is required",
		},
		{
			name: "missing threshold",
			mutate: func(bank *Bank) {
				delete(bank.Meta.Scoring.Thresholds, StrengthSlightA)
			},
			want: "missing threshold slight_a",
		},
		{
			name: "unknown threshold",
			mutate: func(bank *Bank) {
				bank.Meta.Scoring.Thresholds[Strength("extra")] = ThresholdRange{0, 1}
			},
			want: "unknown threshold extra",
		},
		{
			name: "unordered threshold range",
			mutate: func(bank *Bank) {
				bank.Meta.Scoring.Thresholds[StrengthStrongA] = ThresholdRange{99, 13}
			},
			want: "threshold strong_a range",
		},
		{
			name: "overlapping threshold ranges",
			mutate: func(bank *Bank) {
				bank.Meta.Scoring.Thresholds[StrengthModerateA] = ThresholdRange{4, 12}
			},
			want: "threshold moderate_a overlaps threshold slight_a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bank := loadCanonicalV3Bank(t)
			tt.mutate(&bank)

			err := Validate(bank)

			require.Error(t, err)
			require.ErrorContains(t, err, tt.want)
			var validationErr *ValidationError
			require.ErrorAs(t, err, &validationErr)
			require.NotEmpty(t, validationErr.Issues)
		})
	}
}

func loadCanonicalV3Bank(t *testing.T) Bank {
	t.Helper()

	data, err := os.ReadFile(filepath.Join("..", "..", "questions", "mbti-questions-v3.json"))
	require.NoError(t, err)

	return unmarshalBank(t, data)
}

func loadQuestionBankFixture(t *testing.T, name string) Bank {
	t.Helper()

	data, err := os.ReadFile(filepath.Join("testdata", name))
	require.NoError(t, err)

	return unmarshalBank(t, data)
}

func unmarshalBank(t *testing.T, data []byte) Bank {
	t.Helper()

	var bank Bank
	require.NoError(t, json.Unmarshal(data, &bank))
	return bank
}
