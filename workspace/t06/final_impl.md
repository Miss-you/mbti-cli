# T06 Final Implementation

Status: accepted.

## Accepted Plan

Add `internal/scoring` with one exported scorer:

```go
func Score(bank questionbank.Bank, answerSet answers.Set) (Result, error)
```

`Result` reports:

- `Answered`
- `Total`
- `DimensionScores`

The scorer will delegate strict answer validity to `answers.Validate`, then aggregate signed option scores per question dimension. It will not use `Question.Reverse` for any second-pass transformation because the question bank option scores are already authoritative signed values.

## Write Set

- `internal/scoring/scoring_test.go`
- `internal/scoring/scoring.go`
- `openspec/changes/core-assessment/specs/scoring-engine/spec.md`
- `openspec/changes/core-assessment/tasks.md`
- `docs/plans/2026-04-11-core-assessment-design-task.md`
- `workspace/t06/*.md`

## Review Threshold

Accepted with no high-severity findings. The package intentionally stops before threshold classification, type generation, renderer DTOs, or CLI wiring.
