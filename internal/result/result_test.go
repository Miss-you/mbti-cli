package result

import (
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

	require.Equal(t, `{
  "meta": {
    "title": "AI Behavioral Style Assessment v3",
    "version": "0.3.0",
    "answered": 70,
    "total": 70
  },
  "type": "EXTP",
  "dimensions": {
    "EI": {
      "letter": "E",
      "score": 8,
      "strength": "moderate_a",
      "pole": "E (Expansive)",
      "balanced": false
    },
    "SN": {
      "letter": "X",
      "score": 0,
      "strength": "balanced",
      "pole": "balanced",
      "balanced": true
    },
    "TF": {
      "letter": "T",
      "score": 1,
      "strength": "slight_a",
      "pole": "T (Analytical)",
      "balanced": false
    },
    "JP": {
      "letter": "P",
      "score": -2,
      "strength": "slight_b",
      "pole": "P (Flexible)",
      "balanced": false
    }
  }
}
`, string(got))
}

func TestRenderTextProducesReadableFactualSummary(t *testing.T) {
	summary, err := NewSummary(resultBank(), resultScore(), resultClassification())
	require.NoError(t, err)

	got := RenderText(summary)

	require.Equal(t, `AI Behavioral Style Assessment v3 (v0.3.0)
Type: EXTP
Answered: 70/70

Dimensions:
- EI: E, score 8, strength moderate_a, pole E (Expansive)
- SN: X, score 0, strength balanced, pole balanced
- TF: T, score 1, strength slight_a, pole T (Analytical)
- JP: P, score -2, strength slight_b, pole P (Flexible)
`, got)

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
