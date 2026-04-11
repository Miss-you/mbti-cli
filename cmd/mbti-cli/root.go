package cli

import "github.com/spf13/cobra"

func NewRootCmd() *cobra.Command {
	root := &cobra.Command{
		Use:          "mbti-cli",
		Short:        "mbti-cli provides a small CLI for MBTI coding workflows.",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	root.AddCommand(newVersionCmd(), newQuestionsCmd())
	return root
}
