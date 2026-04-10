package cli

import (
	"bytes"
	"testing"

	"github.com/Miss-you/mbti-cli/internal/version"
	"github.com/stretchr/testify/require"
)

func TestNewRootCmd_ExecutesAndPrintsHelp(t *testing.T) {
	cmd := NewRootCmd()
	buf := &bytes.Buffer{}
	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.SetArgs([]string{})

	err := cmd.Execute()
	require.NoError(t, err)
	require.Contains(t, buf.String(), "Usage:")
	require.Contains(t, buf.String(), "mbti-cli")
}

func TestVersionCommand_PrintsVersionInfo(t *testing.T) {
	cmd := NewRootCmd()
	buf := &bytes.Buffer{}
	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.SetArgs([]string{"version"})

	err := cmd.Execute()
	require.NoError(t, err)
	require.Equal(t, version.String()+string('\n'), buf.String())
}
