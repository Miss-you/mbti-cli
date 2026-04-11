# T01 Test Strategy

## What To Prove

T01 must prove that the Go model can unmarshal the canonical v3 question bank without losing fields needed by later tasks.

## Focused Test

Run:

```bash
go test ./internal/questionbank
```

The test will:

- read `questions/mbti-questions-v3.json`
- unmarshal into `questionbank.Bank`
- marshal the model back to JSON and compare it to the canonical raw JSON structure
- assert meta title/version/total
- assert all 70 questions are present
- assert all four dimension metadata entries are addressable
- assert all scoring thresholds unmarshal as exact two-int ranges
- assert every question has unmarshaled id, dimension, reverse, localized scenario, and four options
- assert every option has unmarshaled code, localized label, and signed score
- assert representative first and last question values match the canonical v3 file

## Out Of Scope

- Loader behavior for arbitrary paths
- Schema validation failures
- Answer parsing or validation
- Score aggregation
- Threshold classification
- Result rendering
- CLI commands

## Delivery Gates

Run fresh before closing:

```bash
make fmt
make test
make lint
make build
```

If `golangci-lint` is unavailable locally, record the skipped `make lint` gate in the task board notes and `todo.md`.
