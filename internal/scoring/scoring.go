package scoring

import (
	"fmt"

	"github.com/Miss-you/mbti-cli/internal/answers"
	"github.com/Miss-you/mbti-cli/internal/questionbank"
)

type Result struct {
	Answered        int
	Total           int
	DimensionScores map[questionbank.Dimension]int
}

func Score(bank questionbank.Bank, answerSet answers.Set) (Result, error) {
	if err := answers.Validate(bank, answerSet); err != nil {
		return Result{}, err
	}

	result := Result{
		Answered:        len(answerSet),
		Total:           len(bank.Questions),
		DimensionScores: initializedDimensionScores(),
	}

	for _, question := range bank.Questions {
		optionCode := answerSet[question.ID]
		option, ok := selectedOption(question.Options, optionCode)
		if !ok {
			return Result{}, fmt.Errorf("score answers: question %s option %s is not defined", question.ID, optionCode)
		}
		result.DimensionScores[question.Dimension] += option.Score
	}

	return result, nil
}

func initializedDimensionScores() map[questionbank.Dimension]int {
	return map[questionbank.Dimension]int{
		questionbank.DimensionEI: 0,
		questionbank.DimensionSN: 0,
		questionbank.DimensionTF: 0,
		questionbank.DimensionJP: 0,
	}
}

func selectedOption(options []questionbank.Option, code string) (questionbank.Option, bool) {
	for _, option := range options {
		if option.Code == code {
			return option, true
		}
	}
	return questionbank.Option{}, false
}
