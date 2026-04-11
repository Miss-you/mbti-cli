package scoring

import (
	"path/filepath"
	"testing"

	"github.com/Miss-you/mbti-cli/internal/questionbank"
	"github.com/stretchr/testify/require"
)

func TestClassifyLabelsThresholdBoundaries(t *testing.T) {
	tests := []struct {
		name       string
		score      int
		strength   questionbank.Strength
		wantPole   string
		wantLetter string
	}{
		{name: "strong A lower", score: 13, strength: questionbank.StrengthStrongA, wantPole: "E (Expansive)", wantLetter: "E"},
		{name: "strong A upper", score: 999, strength: questionbank.StrengthStrongA, wantPole: "E (Expansive)", wantLetter: "E"},
		{name: "moderate A lower", score: 5, strength: questionbank.StrengthModerateA, wantPole: "E (Expansive)", wantLetter: "E"},
		{name: "moderate A upper", score: 12, strength: questionbank.StrengthModerateA, wantPole: "E (Expansive)", wantLetter: "E"},
		{name: "slight A lower", score: 1, strength: questionbank.StrengthSlightA, wantPole: "E (Expansive)", wantLetter: "E"},
		{name: "slight A upper", score: 4, strength: questionbank.StrengthSlightA, wantPole: "E (Expansive)", wantLetter: "E"},
		{name: "slight B upper", score: -1, strength: questionbank.StrengthSlightB, wantPole: "I (Focused)", wantLetter: "I"},
		{name: "slight B lower", score: -4, strength: questionbank.StrengthSlightB, wantPole: "I (Focused)", wantLetter: "I"},
		{name: "moderate B upper", score: -5, strength: questionbank.StrengthModerateB, wantPole: "I (Focused)", wantLetter: "I"},
		{name: "moderate B lower", score: -12, strength: questionbank.StrengthModerateB, wantPole: "I (Focused)", wantLetter: "I"},
		{name: "strong B upper", score: -13, strength: questionbank.StrengthStrongB, wantPole: "I (Focused)", wantLetter: "I"},
		{name: "strong B lower", score: -999, strength: questionbank.StrengthStrongB, wantPole: "I (Focused)", wantLetter: "I"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			classification, err := Classify(classifierBank(), Result{
				DimensionScores: map[questionbank.Dimension]int{
					questionbank.DimensionEI: tc.score,
				},
			})

			require.NoError(t, err)
			got := classification.Dimensions[questionbank.DimensionEI]
			require.Equal(t, tc.score, got.Score)
			require.Equal(t, tc.strength, got.Strength)
			require.False(t, got.Balanced)
			require.Equal(t, tc.wantPole, got.Pole)
			require.Equal(t, tc.wantLetter, got.Letter)
		})
	}
}

func TestClassifyBuildsTypeInFixedDimensionOrder(t *testing.T) {
	classification, err := Classify(classifierBank(), Result{
		DimensionScores: map[questionbank.Dimension]int{
			questionbank.DimensionJP: -2,
			questionbank.DimensionTF: 1,
			questionbank.DimensionEI: 2,
			questionbank.DimensionSN: 0,
		},
	})

	require.NoError(t, err)
	require.Equal(t, "EXTP", classification.Type)
	require.Equal(t, "balanced", classification.Dimensions[questionbank.DimensionSN].Pole)
	require.True(t, classification.Dimensions[questionbank.DimensionSN].Balanced)
	require.Equal(t, "X", classification.Dimensions[questionbank.DimensionSN].Letter)
	require.Equal(t, "T (Analytical)", classification.Dimensions[questionbank.DimensionTF].Pole)
	require.Equal(t, "P (Flexible)", classification.Dimensions[questionbank.DimensionJP].Pole)
}

func TestClassifyDefaultsMissingDimensionScoresToBalanced(t *testing.T) {
	classification, err := Classify(classifierBank(), Result{
		DimensionScores: map[questionbank.Dimension]int{
			questionbank.DimensionEI: 1,
		},
	})

	require.NoError(t, err)
	require.Equal(t, "EXXX", classification.Type)
	require.Equal(t, DimensionClassification{
		Score:    0,
		Balanced: true,
		Pole:     "balanced",
		Letter:   "X",
	}, classification.Dimensions[questionbank.DimensionSN])
}

func TestClassifyUsesCanonicalV3BankMetadata(t *testing.T) {
	loaded, err := questionbank.LoadFile(filepath.Join("..", "..", "questions", "mbti-questions-v3.json"))
	require.NoError(t, err)
	require.NoError(t, questionbank.Validate(loaded.Bank))

	classification, err := Classify(loaded.Bank, Result{
		DimensionScores: map[questionbank.Dimension]int{
			questionbank.DimensionEI: 13,
			questionbank.DimensionSN: -5,
			questionbank.DimensionTF: 0,
			questionbank.DimensionJP: 4,
		},
	})

	require.NoError(t, err)
	require.Equal(t, "ENXJ", classification.Type)
	require.Equal(t, questionbank.StrengthStrongA, classification.Dimensions[questionbank.DimensionEI].Strength)
	require.Equal(t, "E (Expansive)", classification.Dimensions[questionbank.DimensionEI].Pole)
	require.Equal(t, questionbank.StrengthModerateB, classification.Dimensions[questionbank.DimensionSN].Strength)
	require.Equal(t, "N (Abstract)", classification.Dimensions[questionbank.DimensionSN].Pole)
	require.True(t, classification.Dimensions[questionbank.DimensionTF].Balanced)
	require.Equal(t, "balanced", classification.Dimensions[questionbank.DimensionTF].Pole)
	require.Equal(t, questionbank.StrengthSlightA, classification.Dimensions[questionbank.DimensionJP].Strength)
	require.Equal(t, "J (Structured)", classification.Dimensions[questionbank.DimensionJP].Pole)
}

func TestClassifyReportsInvalidClassifierInputs(t *testing.T) {
	t.Run("uncovered non-zero score", func(t *testing.T) {
		_, err := Classify(classifierBank(), Result{
			DimensionScores: map[questionbank.Dimension]int{
				questionbank.DimensionEI: 1000,
			},
		})

		require.Error(t, err)
		require.ErrorContains(t, err, "classify EI score 1000")
		require.ErrorContains(t, err, "no threshold range")
	})

	t.Run("missing dimension metadata", func(t *testing.T) {
		bank := classifierBank()
		delete(bank.Meta.Dimensions, questionbank.DimensionEI)

		_, err := Classify(bank, Result{
			DimensionScores: map[questionbank.Dimension]int{
				questionbank.DimensionEI: 1,
			},
		})

		require.Error(t, err)
		require.ErrorContains(t, err, "missing dimension metadata for EI")
	})

	t.Run("missing threshold coverage", func(t *testing.T) {
		bank := classifierBank()
		delete(bank.Meta.Scoring.Thresholds, questionbank.StrengthSlightA)

		_, err := Classify(bank, Result{
			DimensionScores: map[questionbank.Dimension]int{
				questionbank.DimensionEI: 1,
			},
		})

		require.Error(t, err)
		require.ErrorContains(t, err, "classify EI score 1")
		require.ErrorContains(t, err, "no threshold range")
	})
}

func classifierBank() questionbank.Bank {
	return questionbank.Bank{
		Meta: questionbank.Meta{
			Dimensions: map[questionbank.Dimension]questionbank.DimensionMeta{
				questionbank.DimensionEI: {PoleA: "E (Expansive)", PoleB: "I (Focused)"},
				questionbank.DimensionSN: {PoleA: "S (Concrete)", PoleB: "N (Abstract)"},
				questionbank.DimensionTF: {PoleA: "T (Analytical)", PoleB: "F (Empathetic)"},
				questionbank.DimensionJP: {PoleA: "J (Structured)", PoleB: "P (Flexible)"},
			},
			Scoring: questionbank.ScoringMeta{
				Thresholds: map[questionbank.Strength]questionbank.ThresholdRange{
					questionbank.StrengthStrongA:   {13, 999},
					questionbank.StrengthModerateA: {5, 12},
					questionbank.StrengthSlightA:   {1, 4},
					questionbank.StrengthSlightB:   {-4, -1},
					questionbank.StrengthModerateB: {-12, -5},
					questionbank.StrengthStrongB:   {-999, -13},
				},
			},
		},
	}
}
