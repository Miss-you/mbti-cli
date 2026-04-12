package cli

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func cliTestdataPath(name string) string {
	return filepath.Join("testdata", name)
}

func readCLIGolden(t *testing.T, name string) string {
	t.Helper()

	data, err := os.ReadFile(cliTestdataPath(name))
	require.NoError(t, err)
	return string(data)
}
