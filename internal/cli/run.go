package cli

import "github.com/spf13/cobra"

func Execute(cmd *cobra.Command) error {
	return cmd.Execute()
}
