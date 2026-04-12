package result

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Miss-you/mbti-cli/internal/questionbank"
	"github.com/Miss-you/mbti-cli/internal/scoring"
	"github.com/stretchr/testify/require"
)

func TestRenderJSONProducesStableGoldenOutput(t *testing.T) {
	summary, err := NewSummary(resultBank(), resultScore(), resultClassification())
	require.NoError(t, err)

	got, err := RenderJSON(summary)
	require.NoError(t, err)

	require.Equal(t, readGolden(t, "summary.json.golden"), string(got))
}

func TestRenderTextProducesReadableFactualSummary(t *testing.T) {
	summary, err := NewSummary(resultBank(), resultScore(), resultClassification())
	require.NoError(t, err)

	got := RenderText(summary)

	require.Equal(t, readGolden(t, "summary.txt.golden"), got)

	lower := strings.ToLower(got)
	require.NotContains(t, lower, "diagnosis")
	require.NotContains(t, lower, "personality")
	require.NotContains(t, lower, "you are")
}

func TestNewSummaryRejectsMissingDimensionClassification(t *testing.T) {
	classification := resultClassification()
	delete(classification.Dimensions, questionbank.DimensionSN)

	_, err := NewSummary(resultBank(), resultScore(), classification)

	require.Error(t, err)
	require.ErrorContains(t, err, "missing classification for SN")
}

func resultBank() questionbank.Bank {
	return questionbank.Bank{
		Meta: questionbank.Meta{
			Title:   "AI Behavioral Style Assessment v3",
			Version: "0.3.0",
		},
	}
}

func resultScore() scoring.Result {
	return scoring.Result{
		Answered: 70,
		Total:    70,
	}
}

func resultClassification() scoring.Classification {
	return scoring.Classification{
		Type: "EXTP",
		Dimensions: map[questionbank.Dimension]scoring.DimensionClassification{
			questionbank.DimensionEI: {
				Letter:   "E",
				Score:    8,
				Strength: questionbank.StrengthModerateA,
				Pole:     "E (Expansive)",
			},
			questionbank.DimensionSN: {
				Letter:   "X",
				Score:    0,
				Strength: questionbank.StrengthStrongA,
				Balanced: true,
				Pole:     "ignored by balanced renderer",
			},
			questionbank.DimensionTF: {
				Letter:   "T",
				Score:    1,
				Strength: questionbank.StrengthSlightA,
				Pole:     "T (Analytical)",
			},
			questionbank.DimensionJP: {
				Letter:   "P",
				Score:    -2,
				Strength: questionbank.StrengthSlightB,
				Pole:     "P (Flexible)",
			},
		},
	}
}

func readGolden(t *testing.T, name string) string {
	t.Helper()

	data, err := os.ReadFile(filepath.Join("testdata", name))
	require.NoError(t, err)
	return string(data)
}
