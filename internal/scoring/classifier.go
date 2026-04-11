package scoring

import (
	"fmt"
	"strings"

	"github.com/Miss-you/mbti-cli/internal/questionbank"
)

const balancedPole = "balanced"

var classificationDimensionOrder = []questionbank.Dimension{
	questionbank.DimensionEI,
	questionbank.DimensionSN,
	questionbank.DimensionTF,
	questionbank.DimensionJP,
}

var classificationStrengthOrder = []questionbank.Strength{
	questionbank.StrengthStrongA,
	questionbank.StrengthModerateA,
	questionbank.StrengthSlightA,
	questionbank.StrengthSlightB,
	questionbank.StrengthModerateB,
	questionbank.StrengthStrongB,
}

type DimensionClassification struct {
	Score    int
	Strength questionbank.Strength
	Balanced bool
	Pole     string
	Letter   string
}

type Classification struct {
	Type       string
	Dimensions map[questionbank.Dimension]DimensionClassification
}

func Classify(bank questionbank.Bank, result Result) (Classification, error) {
	classification := Classification{
		Dimensions: make(map[questionbank.Dimension]DimensionClassification, len(classificationDimensionOrder)),
	}

	var typeBuilder strings.Builder
	for _, dimension := range classificationDimensionOrder {
		score := result.DimensionScores[dimension]
		dimensionClassification, err := classifyDimension(bank, dimension, score)
		if err != nil {
			return Classification{}, err
		}
		classification.Dimensions[dimension] = dimensionClassification
		typeBuilder.WriteString(dimensionClassification.Letter)
	}

	classification.Type = typeBuilder.String()
	return classification, nil
}

func classifyDimension(bank questionbank.Bank, dimension questionbank.Dimension, score int) (DimensionClassification, error) {
	if score == 0 {
		return DimensionClassification{
			Score:    score,
			Balanced: true,
			Pole:     balancedPole,
			Letter:   "X",
		}, nil
	}

	metadata, ok := bank.Meta.Dimensions[dimension]
	if !ok {
		return DimensionClassification{}, fmt.Errorf("missing dimension metadata for %s", dimension)
	}

	strength, err := classifyStrength(bank.Meta.Scoring.Thresholds, dimension, score)
	if err != nil {
		return DimensionClassification{}, err
	}

	if score > 0 {
		return DimensionClassification{
			Score:    score,
			Strength: strength,
			Pole:     metadata.PoleA,
			Letter:   dimensionLetter(dimension, 0),
		}, nil
	}

	return DimensionClassification{
		Score:    score,
		Strength: strength,
		Pole:     metadata.PoleB,
		Letter:   dimensionLetter(dimension, 1),
	}, nil
}

func classifyStrength(thresholds map[questionbank.Strength]questionbank.ThresholdRange, dimension questionbank.Dimension, score int) (questionbank.Strength, error) {
	for _, strength := range classificationStrengthOrder {
		threshold, ok := thresholds[strength]
		if !ok {
			continue
		}
		if score >= threshold[0] && score <= threshold[1] {
			return strength, nil
		}
	}

	return "", fmt.Errorf("classify %s score %d: no threshold range", dimension, score)
}

func dimensionLetter(dimension questionbank.Dimension, index int) string {
	letters := string(dimension)
	return letters[index : index+1]
}
