package answers

import (
	"encoding/json"
	"fmt"
	"strings"
)

type File struct {
	Answers map[string]string `json:"answers"`
}

type Set map[string]string

func Parse(data []byte) (Set, error) {
	var file File
	if err := json.Unmarshal(data, &file); err != nil {
		return nil, fmt.Errorf("parse answer file: %w", err)
	}

	if file.Answers == nil {
		return nil, fmt.Errorf("answers object is required")
	}

	answers := make(Set, len(file.Answers))
	for questionID, optionCode := range file.Answers {
		answers[questionID] = strings.ToUpper(strings.TrimSpace(optionCode))
	}

	return answers, nil
}
