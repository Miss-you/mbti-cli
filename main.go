package main

import (
	"os"

	cmd "github.com/Miss-you/mbti-cli/cmd/mbti-cli"
	"github.com/Miss-you/mbti-cli/internal/cli"
)

func main() {
	if err := cli.Execute(cmd.NewRootCmd()); err != nil {
		os.Exit(1)
	}
}
