package version

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestString_UsesDefaultBuildMetadata(t *testing.T) {
	require.Equal(t, "mbti-cli version dev (commit none, date unknown)", String())
}

func TestString_UsesInjectedBuildMetadata(t *testing.T) {
	originalVersion := Version
	originalCommit := Commit
	originalDate := Date
	t.Cleanup(func() {
		Version = originalVersion
		Commit = originalCommit
		Date = originalDate
	})

	Version = "1.2.3"
	Commit = "abc1234"
	Date = "2026-04-11"

	require.Equal(t, "mbti-cli version 1.2.3 (commit abc1234, date 2026-04-11)", String())
}
