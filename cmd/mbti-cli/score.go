package cli

import (
	"fmt"
	"os"

	"github.com/Miss-you/mbti-cli/internal/answers"
	"github.com/Miss-you/mbti-cli/internal/questionbank"
	"github.com/Miss-you/mbti-cli/internal/result"
	"github.com/Miss-you/mbti-cli/internal/scoring"
	"github.com/spf13/cobra"
)

const (
	scoreFormatJSON = "json"
	scoreFormatText = "text"
)

type scoreOptions struct {
	bankPath    string
	answersPath string
	format      string
}

func newScoreCmd() *cobra.Command {
	opts := scoreOptions{
		format: scoreFormatText,
	}

	cmd := &cobra.Command{
		Use:   "score",
		Short: "Score an MBTI answer file.",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runScore(cmd, opts)
		},
	}

	cmd.Flags().StringVar(&opts.bankPath, "questions", "", "Path to the question bank JSON file.")
	cmd.Flags().StringVar(&opts.answersPath, "answers", "", "Path to the answer JSON file.")
	cmd.Flags().StringVar(&opts.format, "format", opts.format, "Output format: text or json.")

	return cmd
}

func runScore(cmd *cobra.Command, opts scoreOptions) error {
	if err := validateScoreOptions(opts); err != nil {
		return err
	}

	loaded, err := questionbank.LoadFile(opts.bankPath)
	if err != nil {
		return err
	}
	if err := questionbank.Validate(loaded.Bank); err != nil {
		return fmt.Errorf("validate question bank %q: %w", opts.bankPath, err)
	}

	answerData, err := os.ReadFile(opts.answersPath)
	if err != nil {
		return fmt.Errorf("read answer file %q: %w", opts.answersPath, err)
	}
	answerSet, err := answers.Parse(answerData)
	if err != nil {
		return fmt.Errorf("parse answer file %q: %w", opts.answersPath, err)
	}

	score, err := scoring.Score(loaded.Bank, answerSet)
	if err != nil {
		return err
	}
	classification, err := scoring.Classify(loaded.Bank, score)
	if err != nil {
		return err
	}
	summary, err := result.NewSummary(loaded.Bank, score, classification)
	if err != nil {
		return err
	}

	switch opts.format {
	case scoreFormatJSON:
		data, err := result.RenderJSON(summary)
		if err != nil {
			return err
		}
		_, err = cmd.OutOrStdout().Write(data)
		return err
	case scoreFormatText:
		_, err := fmt.Fprint(cmd.OutOrStdout(), result.RenderText(summary))
		return err
	default:
		return fmt.Errorf("unsupported format %q", opts.format)
	}
}

func validateScoreOptions(opts scoreOptions) error {
	switch opts.format {
	case scoreFormatJSON, scoreFormatText:
	default:
		return fmt.Errorf("unsupported format %q", opts.format)
	}

	if opts.bankPath == "" {
		return fmt.Errorf("question bank path is required")
	}
	if opts.answersPath == "" {
		return fmt.Errorf("answer file path is required")
	}

	return nil
}
