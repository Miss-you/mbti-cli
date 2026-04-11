package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestScoreCommand_RendersCanonicalAnswersAsJSON(t *testing.T) {
	stdout, stderr, err := executeRoot(t,
		"score",
		"--questions", canonicalBankPath(),
		"--answers", canonicalAnswersFile(t, "A"),
		"--format", "json",
	)
	require.NoError(t, err)
	require.Empty(t, stderr)
	require.True(t, bytes.HasSuffix([]byte(stdout), []byte("\n")))

	var got scoreJSONOutput
	require.NoError(t, json.Unmarshal([]byte(stdout), &got))
	require.Equal(t, "AI Behavioral Style Assessment v3", got.Meta.Title)
	require.Equal(t, "0.3.0", got.Meta.Version)
	require.Equal(t, 70, got.Meta.Answered)
	require.Equal(t, 70, got.Meta.Total)
	require.Equal(t, "ESTJ", got.Type)
	require.Equal(t, scoreJSONDimension{Letter: "E", Score: 36, Strength: "strong_a", Pole: "E (Expansive)"}, got.Dimensions.EI)
	require.Equal(t, scoreJSONDimension{Letter: "S", Score: 36, Strength: "strong_a", Pole: "S (Concrete)"}, got.Dimensions.SN)
	require.Equal(t, scoreJSONDimension{Letter: "T", Score: 34, Strength: "strong_a", Pole: "T (Analytical)"}, got.Dimensions.TF)
	require.Equal(t, scoreJSONDimension{Letter: "J", Score: 34, Strength: "strong_a", Pole: "J (Structured)"}, got.Dimensions.JP)
}

func TestScoreCommand_RendersTextSummary(t *testing.T) {
	stdout, stderr, err := executeRoot(t,
		"score",
		"--questions", canonicalBankPath(),
		"--answers", canonicalAnswersFile(t, "D"),
		"--format", "text",
	)
	require.NoError(t, err)
	require.Empty(t, stderr)

	require.Contains(t, stdout, "AI Behavioral Style Assessment v3 (v0.3.0)")
	require.Contains(t, stdout, "Type: INFP")
	require.Contains(t, stdout, "Answered: 70/70")
	require.Contains(t, stdout, "- EI: I, score -20, strength strong_b, pole I (Focused)")
	require.Contains(t, stdout, "- SN: N, score -36, strength strong_b, pole N (Abstract)")
	require.Contains(t, stdout, "- TF: F, score -34, strength strong_b, pole F (Empathetic)")
	require.Contains(t, stdout, "- JP: P, score -34, strength strong_b, pole P (Flexible)")

	lower := strings.ToLower(stdout)
	require.NotContains(t, lower, "diagnosis")
	require.NotContains(t, lower, "personality")
	require.NotContains(t, lower, "you are")
}

func TestScoreCommand_InvalidInputsFailWithoutStdout(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{
			name: "invalid format",
			args: []string{"score", "--questions", canonicalBankPath(), "--answers", canonicalAnswersFile(t, "A"), "--format", "yaml"},
			want: "unsupported format",
		},
		{
			name: "missing question bank path",
			args: []string{"score"},
			want: "question bank path is required",
		},
		{
			name: "missing answer file path",
			args: []string{"score", "--questions", canonicalBankPath()},
			want: "answer file path is required",
		},
		{
			name: "missing answer file",
			args: []string{"score", "--questions", canonicalBankPath(), "--answers", filepath.Join(t.TempDir(), "missing.json")},
			want: "read answer file",
		},
		{
			name: "malformed answer file",
			args: []string{"score", "--questions", canonicalBankPath(), "--answers", writeAnswerFile(t, []byte("{"), "answers.json")},
			want: "parse answer file",
		},
		{
			name: "strict validation missing answer",
			args: []string{"score", "--questions", canonicalBankPath(), "--answers", writeAnswersFile(t, map[string]string{"q01": "A"})},
			want: "missing answer for question q02",
		},
		{
			name: "strict validation unknown question",
			args: []string{"score", "--questions", canonicalBankPath(), "--answers", writeAnswersFile(t, map[string]string{"q999": "A"})},
			want: "unknown question id q999",
		},
		{
			name: "strict validation invalid option",
			args: []string{"score", "--questions", canonicalBankPath(), "--answers", writeAnswersFile(t, map[string]string{"q01": "Z"})},
			want: "question q01 option Z is not valid",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout, stderr, err := executeRoot(t, tt.args...)

			require.Error(t, err)
			require.ErrorContains(t, err, tt.want)
			require.Contains(t, stderr, tt.want)
			require.Empty(t, stdout)
		})
	}
}

type scoreJSONOutput struct {
	Meta struct {
		Title    string `json:"title"`
		Version  string `json:"version"`
		Answered int    `json:"answered"`
		Total    int    `json:"total"`
	} `json:"meta"`
	Type       string `json:"type"`
	Dimensions struct {
		EI scoreJSONDimension `json:"EI"`
		SN scoreJSONDimension `json:"SN"`
		TF scoreJSONDimension `json:"TF"`
		JP scoreJSONDimension `json:"JP"`
	} `json:"dimensions"`
}

type scoreJSONDimension struct {
	Letter   string `json:"letter"`
	Score    int    `json:"score"`
	Strength string `json:"strength"`
	Pole     string `json:"pole"`
	Balanced bool   `json:"balanced"`
}

func canonicalAnswersFile(t *testing.T, optionCode string) string {
	t.Helper()

	answers := make(map[string]string, 70)
	for i := 1; i <= 70; i++ {
		answers[fmt.Sprintf("q%02d", i)] = optionCode
	}
	return writeAnswersFile(t, answers)
}

func writeAnswersFile(t *testing.T, answers map[string]string) string {
	t.Helper()

	data, err := json.Marshal(map[string]map[string]string{
		"answers": answers,
	})
	require.NoError(t, err)
	return writeAnswerFile(t, data, "answers.json")
}

func writeAnswerFile(t *testing.T, data []byte, name string) string {
	t.Helper()

	path := filepath.Join(t.TempDir(), name)
	require.NoError(t, os.WriteFile(path, data, 0o600))
	return path
}
