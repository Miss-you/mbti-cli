# T02 Test Strategy

## What To Prove

T02 must prove the question bank package can load a bank from a filesystem path and return source metadata without doing schema validation.

## Focused Test

Run:

```bash
go test -count=1 ./internal/questionbank
```

The loader tests will:

- load `questions/mbti-questions-v3.json`
- assert the returned `Bank` has representative typed data
- assert source metadata contains the input path, base filename, and positive byte size
- assert a missing file returns an error with read/path context and still wraps `os.ErrNotExist`
- assert a malformed JSON file returns an error with parse/path context
- assert an empty path returns a required-path error

## Out Of Scope

- Schema validation failures
- Dimension count checks
- Threshold validation
- Answer parsing or validation
- Scoring
- CLI command smoke tests

## Delivery Gates

Run fresh before closing:

```bash
make fmt
make test
make lint
make build
openspec validate core-assessment
```

If `golangci-lint` is unavailable locally, record the skipped `make lint` gate in the task board notes and `todo.md`.
