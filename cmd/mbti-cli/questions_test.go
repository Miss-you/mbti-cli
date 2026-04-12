package cli

import (
	"bytes"
	"encoding/json"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQuestionsCommand_RendersCanonicalBankAsJSON(t *testing.T) {
	stdout, stderr, err := executeRoot(t,
		"questions",
		"--questions", canonicalBankPath(),
		"--format", "json",
	)
	require.NoError(t, err)
	require.Empty(t, stderr)
	require.True(t, bytes.HasSuffix([]byte(stdout), []byte("\n")))

	var got questionsJSONOutput
	require.NoError(t, json.Unmarshal([]byte(stdout), &got))
	require.Equal(t, "AI Behavioral Style Assessment v3", got.Meta.Title)
	require.Equal(t, "0.3.0", got.Meta.Version)
	require.Equal(t, "zh", got.Meta.Language)
	require.Equal(t, 70, got.Meta.Count)
	require.Equal(t, 70, got.Meta.Total)
	require.Equal(t, canonicalBankPath(), got.Meta.Source)
	require.Len(t, got.Questions, 70)
	require.Equal(t, "q01", got.Questions[0].ID)
	require.Contains(t, got.Questions[0].Scenario, "用户让你帮忙")
	require.Len(t, got.Questions[0].Options, 4)
	require.Equal(t, "A", got.Questions[0].Options[0].Code)
	require.Contains(t, got.Questions[0].Options[0].Label, "主动标注文化梗")

	require.NotContains(t, stdout, `"score"`)
	require.NotContains(t, stdout, `"reverse"`)
	require.NotContains(t, stdout, `"thresholds"`)
	require.NotContains(t, stdout, `"dimensions"`)
}

func TestQuestionsCommand_RendersLocalizedTextWithCount(t *testing.T) {
	stdout, stderr, err := executeRoot(t,
		"questions",
		"--questions", canonicalBankPath(),
		"--format", "text",
		"--count", "1",
		"--lang", "en",
	)
	require.NoError(t, err)
	require.Empty(t, stderr)

	require.Equal(t, readCLIGolden(t, "questions-one-en.txt.golden"), stdout)
}

func TestQuestionsCommand_SeededSelectionIsDeterministic(t *testing.T) {
	firstStdout, firstStderr, err := executeRoot(t,
		"questions",
		"--questions", canonicalBankPath(),
		"--format", "json",
		"--count", "3",
		"--seed", "7",
	)
	require.NoError(t, err)
	require.Empty(t, firstStderr)

	secondStdout, secondStderr, err := executeRoot(t,
		"questions",
		"--questions", canonicalBankPath(),
		"--format", "json",
		"--count", "3",
		"--seed", "7",
	)
	require.NoError(t, err)
	require.Empty(t, secondStderr)

	firstIDs := questionIDs(t, firstStdout)
	secondIDs := questionIDs(t, secondStdout)
	require.Equal(t, firstIDs, secondIDs)
	require.Len(t, firstIDs, 3)
	require.NotEqual(t, []string{"q01", "q02", "q03"}, firstIDs)
}

func TestQuestionsCommand_SeededSelectionDefaultsToAllQuestions(t *testing.T) {
	stdout, stderr, err := executeRoot(t,
		"questions",
		"--questions", canonicalBankPath(),
		"--format", "json",
		"--seed", "7",
	)
	require.NoError(t, err)
	require.Empty(t, stderr)

	ids := questionIDs(t, stdout)
	require.Len(t, ids, 70)
	require.NotEqual(t, "q01", ids[0])
}

func TestQuestionsCommand_InvalidInputsFailWithoutStdout(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{
			name: "invalid format",
			args: []string{"questions", "--questions", canonicalBankPath(), "--format", "yaml"},
			want: "unsupported format",
		},
		{
			name: "invalid language",
			args: []string{"questions", "--questions", canonicalBankPath(), "--lang", "fr"},
			want: "unsupported language",
		},
		{
			name: "negative count",
			args: []string{"questions", "--questions", canonicalBankPath(), "--count", "-1"},
			want: "count must be non-negative",
		},
		{
			name: "count too large",
			args: []string{"questions", "--questions", canonicalBankPath(), "--count", "71"},
			want: "count 71 exceeds available questions 70",
		},
		{
			name: "missing path",
			args: []string{"questions"},
			want: "question bank path is required",
		},
		{
			name: "missing file",
			args: []string{"questions", "--questions", filepath.Join(t.TempDir(), "missing.json")},
			want: "read question bank",
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

type questionsJSONOutput struct {
	Meta struct {
		Title    string `json:"title"`
		Version  string `json:"version"`
		Language string `json:"language"`
		Count    int    `json:"count"`
		Total    int    `json:"total"`
		Source   string `json:"source"`
	} `json:"meta"`
	Questions []struct {
		ID       string `json:"id"`
		Scenario string `json:"scenario"`
		Options  []struct {
			Code  string `json:"code"`
			Label string `json:"label"`
		} `json:"options"`
	} `json:"questions"`
}

func executeRoot(t *testing.T, args ...string) (string, string, error) {
	t.Helper()

	cmd := NewRootCmd()
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd.SetOut(stdout)
	cmd.SetErr(stderr)
	cmd.SetArgs(args)

	err := cmd.Execute()
	return stdout.String(), stderr.String(), err
}

func canonicalBankPath() string {
	return filepath.Join("..", "..", "questions", "mbti-questions-v3.json")
}

func questionIDs(t *testing.T, data string) []string {
	t.Helper()

	var output questionsJSONOutput
	require.NoError(t, json.Unmarshal([]byte(data), &output))

	ids := make([]string, 0, len(output.Questions))
	for _, question := range output.Questions {
		ids = append(ids, question.ID)
	}
	return ids
}
