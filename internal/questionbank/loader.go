package questionbank

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Source struct {
	Path      string
	FileName  string
	SizeBytes int64
}

type LoadedBank struct {
	Bank   Bank
	Source Source
}

func LoadFile(path string) (LoadedBank, error) {
	if path == "" {
		return LoadedBank{}, fmt.Errorf("question bank path is required")
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return LoadedBank{}, fmt.Errorf("read question bank %q: %w", path, err)
	}

	var bank Bank
	if err := json.Unmarshal(data, &bank); err != nil {
		return LoadedBank{}, fmt.Errorf("parse question bank %q: %w", path, err)
	}

	return LoadedBank{
		Bank: bank,
		Source: Source{
			Path:      path,
			FileName:  filepath.Base(path),
			SizeBytes: int64(len(data)),
		},
	}, nil
}
