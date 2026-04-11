# T07 Final Implementation

Status: accepted.

## Accepted Plan

Add threshold and type classification to `internal/scoring` while preserving the raw `Score` API.

Public API:

```go
type DimensionClassification struct {
    Score    int
    Strength questionbank.Strength
    Balanced bool
    Pole     string
    Letter   string
}

type Classification struct {
    Type       string
    Dimensions map[questionbank.Dimension]DimensionClassification
}

func Classify(bank questionbank.Bank, result Result) (Classification, error)
```

Behavior:

- Classification uses `bank.Meta.Scoring.Thresholds`.
- Positive scores map to `bank.Meta.Dimensions[dim].PoleA` and the first dimension letter.
- Negative scores map to `bank.Meta.Dimensions[dim].PoleB` and the second dimension letter.
- Score `0` maps to pole `balanced`, letter `X`, `Balanced=true`, and no threshold bucket.
- Type generation is deterministic in `EI`, `SN`, `TF`, `JP` order, using `X` for any balanced dimension.
- Missing scores for supported dimensions default to `0`, matching the existing scorer initialization.
- Missing threshold ranges, missing dimension metadata, or uncovered non-zero scores return contextual errors.

## Write Set

- `internal/scoring/classifier_test.go`
- `internal/scoring/classifier.go`
- `openspec/changes/core-assessment/specs/scoring-classifier/spec.md`
- `openspec/changes/core-assessment/tasks.md`
- `docs/plans/2026-04-11-core-assessment-design-task.md`
- `workspace/t07/*.md`

## Review

- First review: 84/100, two high-severity issues fixed.
- Second review: 97/100, no high-severity issues.
