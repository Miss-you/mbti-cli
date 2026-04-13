# T01 Test Strategy

## What To Prove

T01 must prove that the Go model can unmarshal the canonical v3 question bank without losing representative fields needed by later tasks, while keeping validation, loading, scoring, rendering, and CLI behavior outside the model task.

## Focused Test

Run:

```bash
go test -v ./internal/questionbank -run TestQuestionbank -count=1 -ginkgo.v
go test -count=1 ./internal/questionbank
```

The test will:

- read `questions/mbti-questions-v3.json`
- unmarshal directly into `questionbank.Bank`
- assert representative meta title/title_zh/version/total
- assert a representative dimension metadata entry is available by typed dimension key
- assert representative scoring threshold ranges unmarshal as exact two-int ranges
- assert all 70 questions are present without requiring exhaustive per-question checks
- assert one canonical first question preserves id, dimension, reverse flag, localized scenario, option code, localized option label, and signed score
- assert one reversed question preserves the `reverse` flag
- avoid full JSON round-trip equality, schema-invalid fixtures, loader behavior, validation behavior, scoring behavior, rendering behavior, and CLI behavior in T01 BDD

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
