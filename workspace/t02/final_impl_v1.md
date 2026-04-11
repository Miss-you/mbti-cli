# T02 Final Implementation v1

## Product Writes

- Add `internal/questionbank/loader.go`.
- Add `internal/questionbank/loader_test.go`.

## API

```go
type Source struct {
	Path      string
	FileName  string
	SizeBytes int64
}

type LoadedBank struct {
	Bank   Bank
	Source Source
}

func LoadFile(path string) (LoadedBank, error)
```

## Behavior

`LoadFile` will:

1. Reject an empty path with a clear error.
2. Read the file from the supplied path.
3. Unmarshal the file into `Bank`.
4. Return `LoadedBank` with the typed bank and source metadata.
5. Wrap read errors with the path and preserve `errors.Is(err, os.ErrNotExist)`.
6. Wrap JSON parse errors with the path and parse context.

## Tests First

Add loader unit tests before implementation:

- canonical v3 file loads into `LoadedBank`
- `Source.Path`, `Source.FileName`, and `Source.SizeBytes` are populated
- missing file returns an error containing operation and path context
- malformed file returns an error containing parse context and path
- empty path returns a clear required-path error

## OpenSpec

Extend `core-assessment` with a loader requirement under the question bank model spec. Add task items for failing loader tests, implementation, and focused package test.

## Review Pass

Plan review score: 91/100.

No high-severity issues found. The plan stays within T02 and leaves schema validation for T03.
