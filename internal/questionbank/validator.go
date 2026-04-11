package questionbank

import (
	"fmt"
	"sort"
	"strings"
)

var supportedDimensions = []Dimension{
	DimensionEI,
	DimensionSN,
	DimensionTF,
	DimensionJP,
}

var supportedStrengths = []Strength{
	StrengthStrongA,
	StrengthModerateA,
	StrengthSlightA,
	StrengthSlightB,
	StrengthModerateB,
	StrengthStrongB,
}

type ValidationError struct {
	Issues []string
}

func (err *ValidationError) Error() string {
	if err == nil || len(err.Issues) == 0 {
		return "question bank validation failed"
	}

	return "question bank validation failed: " + strings.Join(err.Issues, "; ")
}

func Validate(bank Bank) error {
	var issues []string

	if bank.Meta.Total != len(bank.Questions) {
		issues = append(issues, fmt.Sprintf("meta.total=%d does not match questions length %d", bank.Meta.Total, len(bank.Questions)))
	}

	issues = append(issues, validateDimensionMetadata(bank.Meta.Dimensions)...)

	questionCounts := make(map[Dimension]int, len(supportedDimensions))
	seenQuestionIDs := make(map[string]int, len(bank.Questions))
	for i, question := range bank.Questions {
		issues = append(issues, validateQuestion(i, question, seenQuestionIDs)...)
		if isSupportedDimension(question.Dimension) {
			questionCounts[question.Dimension]++
		}
	}

	for _, dimension := range supportedDimensions {
		meta, ok := bank.Meta.Dimensions[dimension]
		if !ok {
			continue
		}
		if meta.Count != questionCounts[dimension] {
			issues = append(issues, fmt.Sprintf("dimension %s count=%d does not match questions count %d", dimension, meta.Count, questionCounts[dimension]))
		}
	}

	issues = append(issues, validateThresholds(bank.Meta.Scoring.Thresholds)...)

	if len(issues) > 0 {
		return &ValidationError{Issues: issues}
	}

	return nil
}

func validateDimensionMetadata(dimensions map[Dimension]DimensionMeta) []string {
	var issues []string

	for _, dimension := range supportedDimensions {
		if _, ok := dimensions[dimension]; !ok {
			issues = append(issues, fmt.Sprintf("missing dimension metadata for %s", dimension))
		}
	}

	for _, dimension := range sortedDimensionKeys(dimensions) {
		if !isSupportedDimension(dimension) {
			issues = append(issues, fmt.Sprintf("unknown dimension metadata %s", dimension))
		}
	}

	return issues
}

func validateQuestion(index int, question Question, seenQuestionIDs map[string]int) []string {
	var issues []string

	if strings.TrimSpace(question.ID) == "" {
		issues = append(issues, fmt.Sprintf("questions[%d].id is required", index))
	} else if firstIndex, ok := seenQuestionIDs[question.ID]; ok {
		issues = append(issues, fmt.Sprintf("duplicate question id %s at questions[%d] and questions[%d]", question.ID, firstIndex, index))
	} else {
		seenQuestionIDs[question.ID] = index
	}

	if !isSupportedDimension(question.Dimension) {
		issues = append(issues, fmt.Sprintf("questions[%d].dimension %q is not supported", index, question.Dimension))
	}

	if strings.TrimSpace(question.Scenario.ZH) == "" {
		issues = append(issues, fmt.Sprintf("questions[%d].scenario.zh is required", index))
	}
	if strings.TrimSpace(question.Scenario.EN) == "" {
		issues = append(issues, fmt.Sprintf("questions[%d].scenario.en is required", index))
	}

	issues = append(issues, validateOptions(index, question.Options)...)

	return issues
}

func validateOptions(questionIndex int, options []Option) []string {
	var issues []string

	if len(options) != 4 {
		issues = append(issues, fmt.Sprintf("questions[%d].options must contain exactly 4 options", questionIndex))
	}

	seenCodes := make(map[string]int, len(options))
	for i, option := range options {
		if !isSupportedOptionCode(option.Code) {
			issues = append(issues, fmt.Sprintf("questions[%d].options[%d].code %q is not supported", questionIndex, i, option.Code))
		} else if firstIndex, ok := seenCodes[option.Code]; ok {
			issues = append(issues, fmt.Sprintf("duplicate option code %s at questions[%d].options[%d] and questions[%d].options[%d]", option.Code, questionIndex, firstIndex, questionIndex, i))
		} else {
			seenCodes[option.Code] = i
		}

		if strings.TrimSpace(option.Label.ZH) == "" {
			issues = append(issues, fmt.Sprintf("questions[%d].options[%d].label.zh is required", questionIndex, i))
		}
		if strings.TrimSpace(option.Label.EN) == "" {
			issues = append(issues, fmt.Sprintf("questions[%d].options[%d].label.en is required", questionIndex, i))
		}
		if !isSupportedOptionScore(option.Score) {
			issues = append(issues, fmt.Sprintf("questions[%d].options[%d].score %d is not supported", questionIndex, i, option.Score))
		}
	}

	for _, code := range []string{"A", "B", "C", "D"} {
		if _, ok := seenCodes[code]; !ok {
			issues = append(issues, fmt.Sprintf("questions[%d].options missing code %s", questionIndex, code))
		}
	}

	return issues
}

func validateThresholds(thresholds map[Strength]ThresholdRange) []string {
	var issues []string

	for _, strength := range supportedStrengths {
		threshold, ok := thresholds[strength]
		if !ok {
			issues = append(issues, fmt.Sprintf("missing threshold %s", strength))
			continue
		}
		if threshold[0] > threshold[1] {
			issues = append(issues, fmt.Sprintf("threshold %s range [%d,%d] is unordered", strength, threshold[0], threshold[1]))
		}
	}

	for _, strength := range sortedStrengthKeys(thresholds) {
		if !isSupportedStrength(strength) {
			issues = append(issues, fmt.Sprintf("unknown threshold %s", strength))
		}
	}

	for i, strength := range supportedStrengths {
		left, ok := thresholds[strength]
		if !ok || left[0] > left[1] {
			continue
		}
		for _, other := range supportedStrengths[i+1:] {
			right, ok := thresholds[other]
			if !ok || right[0] > right[1] {
				continue
			}
			if rangesOverlap(left, right) {
				issues = append(issues, fmt.Sprintf("threshold %s overlaps threshold %s", strength, other))
			}
		}
	}

	return issues
}

func sortedDimensionKeys(dimensions map[Dimension]DimensionMeta) []Dimension {
	keys := make([]Dimension, 0, len(dimensions))
	for dimension := range dimensions {
		keys = append(keys, dimension)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	return keys
}

func sortedStrengthKeys(thresholds map[Strength]ThresholdRange) []Strength {
	keys := make([]Strength, 0, len(thresholds))
	for strength := range thresholds {
		keys = append(keys, strength)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	return keys
}

func isSupportedDimension(dimension Dimension) bool {
	for _, supported := range supportedDimensions {
		if dimension == supported {
			return true
		}
	}
	return false
}

func isSupportedStrength(strength Strength) bool {
	for _, supported := range supportedStrengths {
		if strength == supported {
			return true
		}
	}
	return false
}

func isSupportedOptionCode(code string) bool {
	switch code {
	case "A", "B", "C", "D":
		return true
	default:
		return false
	}
}

func isSupportedOptionScore(score int) bool {
	switch score {
	case -2, -1, 1, 2:
		return true
	default:
		return false
	}
}

func rangesOverlap(left ThresholdRange, right ThresholdRange) bool {
	return left[0] <= right[1] && right[0] <= left[1]
}
