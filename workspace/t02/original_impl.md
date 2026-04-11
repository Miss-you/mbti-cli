# T02 Original Implementation

## Current Behavior

T01 has added representation-only question bank model types in `internal/questionbank`.

Current facts:

- `Bank` mirrors the canonical v3 JSON shape.
- `model_test.go` proves `questions/mbti-questions-v3.json` can unmarshal without field loss.
- There is no package API that reads a bank from a filesystem path.
- Callers must currently combine `os.ReadFile` and `json.Unmarshal` themselves.
- The active OpenSpec change only covers the model and explicitly leaves loading to later tasks.

## Gap For T02

T02 needs a loader that:

- accepts a file path
- reads the JSON from disk
- unmarshals it into `Bank`
- returns source metadata useful for diagnostics and later CLI/result paths
- reports missing and malformed files with clear, wrapped errors

## Out Of Scope

- Schema validation
- Dimension/count/threshold checks
- Answer parsing or validation
- Scoring
- CLI command wiring
