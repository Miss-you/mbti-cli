package answers

import (
	"fmt"
	"sort"
	"strings"

	"github.com/Miss-you/mbti-cli/internal/questionbank"
)

type ValidationError struct {
	Issues []string
}

func (err *ValidationError) Error() string {
	if err == nil || len(err.Issues) == 0 {
		return "answer validation failed"
	}

	return "answer validation failed: " + strings.Join(err.Issues, "; ")
}

func Validate(bank questionbank.Bank, answerSet Set) error {
	questionOptions := make(map[string]map[string]struct{}, len(bank.Questions))
	questionOrder := make([]string, 0, len(bank.Questions))
	for _, question := range bank.Questions {
		questionOptions[question.ID] = optionCodes(question.Options)
		questionOrder = append(questionOrder, question.ID)
	}

	var issues []string
	for _, questionID := range sortedAnswerIDs(answerSet) {
		optionCode := answerSet[questionID]
		options, ok := questionOptions[questionID]
		if !ok {
			issues = append(issues, fmt.Sprintf("unknown question id %s", questionID))
			continue
		}
		if _, ok := options[optionCode]; !ok {
			issues = append(issues, fmt.Sprintf("question %s option %s is not valid", questionID, optionCode))
		}
	}

	for _, questionID := range questionOrder {
		if _, ok := answerSet[questionID]; !ok {
			issues = append(issues, fmt.Sprintf("missing answer for question %s", questionID))
		}
	}

	if len(issues) > 0 {
		return &ValidationError{Issues: issues}
	}

	return nil
}

func optionCodes(options []questionbank.Option) map[string]struct{} {
	codes := make(map[string]struct{}, len(options))
	for _, option := range options {
		codes[option.Code] = struct{}{}
	}
	return codes
}

func sortedAnswerIDs(answerSet Set) []string {
	ids := make([]string, 0, len(answerSet))
	for id := range answerSet {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	return ids
}
