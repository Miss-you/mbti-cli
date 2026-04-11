package result

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/Miss-you/mbti-cli/internal/questionbank"
	"github.com/Miss-you/mbti-cli/internal/scoring"
)

const balancedStrength = "balanced"

type Meta struct {
	Title    string `json:"title"`
	Version  string `json:"version"`
	Answered int    `json:"answered"`
	Total    int    `json:"total"`
}

type Dimension struct {
	Letter   string `json:"letter"`
	Score    int    `json:"score"`
	Strength string `json:"strength"`
	Pole     string `json:"pole"`
	Balanced bool   `json:"balanced"`
}

type Dimensions struct {
	EI Dimension `json:"EI"`
	SN Dimension `json:"SN"`
	TF Dimension `json:"TF"`
	JP Dimension `json:"JP"`
}

type Summary struct {
	Meta       Meta       `json:"meta"`
	Type       string     `json:"type"`
	Dimensions Dimensions `json:"dimensions"`
}

func NewSummary(bank questionbank.Bank, score scoring.Result, classification scoring.Classification) (Summary, error) {
	ei, err := buildDimension(classification, questionbank.DimensionEI)
	if err != nil {
		return Summary{}, err
	}
	sn, err := buildDimension(classification, questionbank.DimensionSN)
	if err != nil {
		return Summary{}, err
	}
	tf, err := buildDimension(classification, questionbank.DimensionTF)
	if err != nil {
		return Summary{}, err
	}
	jp, err := buildDimension(classification, questionbank.DimensionJP)
	if err != nil {
		return Summary{}, err
	}

	return Summary{
		Meta: Meta{
			Title:    bank.Meta.Title,
			Version:  bank.Meta.Version,
			Answered: score.Answered,
			Total:    score.Total,
		},
		Type: classification.Type,
		Dimensions: Dimensions{
			EI: ei,
			SN: sn,
			TF: tf,
			JP: jp,
		},
	}, nil
}

func RenderJSON(summary Summary) ([]byte, error) {
	data, err := json.MarshalIndent(summary, "", "  ")
	if err != nil {
		return nil, err
	}
	return append(data, '\n'), nil
}

func RenderText(summary Summary) string {
	var builder strings.Builder
	builder.WriteString(summary.Meta.Title)
	builder.WriteString(" (v")
	builder.WriteString(summary.Meta.Version)
	builder.WriteString(")\n")
	builder.WriteString("Type: ")
	builder.WriteString(summary.Type)
	builder.WriteString("\n")
	builder.WriteString("Answered: ")
	builder.WriteString(strconv.Itoa(summary.Meta.Answered))
	builder.WriteString("/")
	builder.WriteString(strconv.Itoa(summary.Meta.Total))
	builder.WriteString("\n\n")
	builder.WriteString("Dimensions:\n")
	writeTextDimension(&builder, "EI", summary.Dimensions.EI)
	writeTextDimension(&builder, "SN", summary.Dimensions.SN)
	writeTextDimension(&builder, "TF", summary.Dimensions.TF)
	writeTextDimension(&builder, "JP", summary.Dimensions.JP)
	return builder.String()
}

func buildDimension(classification scoring.Classification, dimension questionbank.Dimension) (Dimension, error) {
	classified, ok := classification.Dimensions[dimension]
	if !ok {
		return Dimension{}, fmt.Errorf("missing classification for %s", dimension)
	}

	strength := string(classified.Strength)
	pole := classified.Pole
	if classified.Balanced {
		strength = balancedStrength
		pole = balancedStrength
	}

	return Dimension{
		Letter:   classified.Letter,
		Score:    classified.Score,
		Strength: strength,
		Pole:     pole,
		Balanced: classified.Balanced,
	}, nil
}

func writeTextDimension(builder *strings.Builder, label string, dimension Dimension) {
	builder.WriteString("- ")
	builder.WriteString(label)
	builder.WriteString(": ")
	builder.WriteString(dimension.Letter)
	builder.WriteString(", score ")
	builder.WriteString(strconv.Itoa(dimension.Score))
	builder.WriteString(", strength ")
	builder.WriteString(dimension.Strength)
	builder.WriteString(", pole ")
	builder.WriteString(dimension.Pole)
	builder.WriteString("\n")
}
