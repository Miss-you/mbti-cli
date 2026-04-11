package questionbank

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadFileLoadsCanonicalV3QuestionBankWithSourceMetadata(t *testing.T) {
	path := filepath.Join("..", "..", "questions", "mbti-questions-v3.json")

	loaded, err := LoadFile(path)
	require.NoError(t, err)

	require.Equal(t, "AI Behavioral Style Assessment v3", loaded.Bank.Meta.Title)
	require.Equal(t, "0.3.0", loaded.Bank.Meta.Version)
	require.Len(t, loaded.Bank.Questions, 70)
	require.Equal(t, "q01", loaded.Bank.Questions[0].ID)

	require.Equal(t, path, loaded.Source.Path)
	require.Equal(t, "mbti-questions-v3.json", loaded.Source.FileName)
	require.Positive(t, loaded.Source.SizeBytes)
}

func TestLoadFileReturnsClearErrorForMissingFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "missing-bank.json")

	_, err := LoadFile(path)

	require.Error(t, err)
	require.ErrorContains(t, err, "read question bank")
	require.ErrorContains(t, err, path)
	require.True(t, errors.Is(err, os.ErrNotExist))
}

func TestLoadFileReturnsClearErrorForMalformedJSON(t *testing.T) {
	path := filepath.Join(t.TempDir(), "malformed-bank.json")
	require.NoError(t, os.WriteFile(path, []byte(`{"meta":`), 0o600))

	_, err := LoadFile(path)

	require.Error(t, err)
	require.ErrorContains(t, err, "parse question bank")
	require.ErrorContains(t, err, path)
}

func TestLoadFileRejectsEmptyPath(t *testing.T) {
	_, err := LoadFile("")

	require.Error(t, err)
	require.ErrorContains(t, err, "question bank path is required")
}
