package cli

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"

	"github.com/Miss-you/mbti-cli/internal/questionbank"
	"github.com/spf13/cobra"
)

const (
	questionsFormatJSON = "json"
	questionsFormatText = "text"
	questionsLangZH     = "zh"
	questionsLangEN     = "en"
)

type questionsOptions struct {
	bankPath string
	format   string
	lang     string
	count    int
	seed     int64
	seedSet  bool
}

type questionSetOutput struct {
	Meta      questionSetMeta  `json:"meta"`
	Questions []questionOutput `json:"questions"`
}

type questionSetMeta struct {
	Title    string `json:"title"`
	Version  string `json:"version"`
	Language string `json:"language"`
	Count    int    `json:"count"`
	Total    int    `json:"total"`
	Source   string `json:"source"`
}

type questionOutput struct {
	ID       string         `json:"id"`
	Scenario string         `json:"scenario"`
	Options  []optionOutput `json:"options"`
}

type optionOutput struct {
	Code  string `json:"code"`
	Label string `json:"label"`
}

func newQuestionsCmd() *cobra.Command {
	opts := questionsOptions{
		format: questionsFormatText,
		lang:   questionsLangZH,
	}

	cmd := &cobra.Command{
		Use:   "questions",
		Short: "Export a selected MBTI question set.",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.seedSet = cmd.Flags().Changed("seed")
			return runQuestions(cmd, opts)
		},
	}

	cmd.Flags().StringVar(&opts.bankPath, "questions", "", "Path to the question bank JSON file.")
	cmd.Flags().StringVar(&opts.format, "format", opts.format, "Output format: text or json.")
	cmd.Flags().StringVar(&opts.lang, "lang", opts.lang, "Question language: zh or en.")
	cmd.Flags().IntVar(&opts.count, "count", 0, "Number of questions to export. Use 0 for all questions.")
	cmd.Flags().Int64Var(&opts.seed, "seed", 0, "Deterministic shuffle seed applied before count selection.")

	return cmd
}

func runQuestions(cmd *cobra.Command, opts questionsOptions) error {
	if err := validateQuestionsOptions(opts); err != nil {
		return err
	}

	loaded, err := questionbank.LoadFile(opts.bankPath)
	if err != nil {
		return err
	}
	if err := questionbank.Validate(loaded.Bank); err != nil {
		return fmt.Errorf("validate question bank %q: %w", opts.bankPath, err)
	}

	selected, err := selectQuestions(loaded.Bank.Questions, opts.count, opts.seed, opts.seedSet)
	if err != nil {
		return err
	}
	output := buildQuestionSetOutput(loaded, selected, opts.lang)

	switch opts.format {
	case questionsFormatJSON:
		data, err := json.MarshalIndent(output, "", "  ")
		if err != nil {
			return err
		}
		data = append(data, '\n')
		_, err = cmd.OutOrStdout().Write(data)
		return err
	case questionsFormatText:
		_, err := fmt.Fprint(cmd.OutOrStdout(), renderQuestionSetText(output))
		return err
	default:
		return fmt.Errorf("unsupported format %q", opts.format)
	}
}

func validateQuestionsOptions(opts questionsOptions) error {
	switch opts.format {
	case questionsFormatJSON, questionsFormatText:
	default:
		return fmt.Errorf("unsupported format %q", opts.format)
	}

	switch opts.lang {
	case questionsLangZH, questionsLangEN:
	default:
		return fmt.Errorf("unsupported language %q", opts.lang)
	}

	if opts.count < 0 {
		return fmt.Errorf("count must be non-negative")
	}

	return nil
}

func selectQuestions(questions []questionbank.Question, count int, seed int64, seedSet bool) ([]questionbank.Question, error) {
	if count > len(questions) {
		return nil, fmt.Errorf("count %d exceeds available questions %d", count, len(questions))
	}

	selected := append([]questionbank.Question(nil), questions...)
	if seedSet {
		random := rand.New(rand.NewSource(seed))
		random.Shuffle(len(selected), func(i, j int) {
			selected[i], selected[j] = selected[j], selected[i]
		})
	}

	if count == 0 {
		return selected, nil
	}

	return selected[:count], nil
}

func buildQuestionSetOutput(loaded questionbank.LoadedBank, selected []questionbank.Question, lang string) questionSetOutput {
	questions := make([]questionOutput, 0, len(selected))
	for _, question := range selected {
		questions = append(questions, buildQuestionOutput(question, lang))
	}

	return questionSetOutput{
		Meta: questionSetMeta{
			Title:    loaded.Bank.Meta.Title,
			Version:  loaded.Bank.Meta.Version,
			Language: lang,
			Count:    len(questions),
			Total:    len(loaded.Bank.Questions),
			Source:   loaded.Source.Path,
		},
		Questions: questions,
	}
}

func buildQuestionOutput(question questionbank.Question, lang string) questionOutput {
	options := make([]optionOutput, 0, len(question.Options))
	for _, option := range question.Options {
		options = append(options, optionOutput{
			Code:  option.Code,
			Label: localizedText(option.Label, lang),
		})
	}

	return questionOutput{
		ID:       question.ID,
		Scenario: localizedText(question.Scenario, lang),
		Options:  options,
	}
}

func localizedText(text questionbank.LocalizedText, lang string) string {
	if lang == questionsLangEN {
		return text.EN
	}
	return text.ZH
}

func renderQuestionSetText(output questionSetOutput) string {
	var builder strings.Builder
	builder.WriteString(output.Meta.Title)
	builder.WriteString(" (v")
	builder.WriteString(output.Meta.Version)
	builder.WriteString(")\n")
	builder.WriteString("Questions: ")
	fmt.Fprintf(&builder, "%d/%d\n", output.Meta.Count, output.Meta.Total)
	builder.WriteString("Language: ")
	builder.WriteString(output.Meta.Language)
	builder.WriteString("\n\n")

	for _, question := range output.Questions {
		builder.WriteString(question.ID)
		builder.WriteString(". ")
		builder.WriteString(question.Scenario)
		builder.WriteString("\n")
		for _, option := range question.Options {
			builder.WriteString("  ")
			builder.WriteString(option.Code)
			builder.WriteString(". ")
			builder.WriteString(option.Label)
			builder.WriteString("\n")
		}
	}

	return builder.String()
}
