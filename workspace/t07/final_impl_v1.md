# T07 Final Implementation v1

Status: revised for review.

## Plan

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
- Positive scores map to `bank.Meta.Dimensions[dim].PoleA` and the first letter of the dimension name: `E`, `S`, `T`, `J`.
- Negative scores map to `bank.Meta.Dimensions[dim].PoleB` and the second letter of the dimension name: `I`, `N`, `F`, `P`.
- Score `0` maps to pole `balanced`, letter `X`, `Balanced=true`, and no threshold bucket.
- Type generation is deterministic in `EI`, `SN`, `TF`, `JP` order, using `X` in any balanced position.
- Missing scores for supported dimensions default to `0`, matching the existing raw scorer's initialized dimension behavior.
- Missing threshold ranges or uncovered non-zero scores return errors instead of silently inventing a strength.
- Missing dimension metadata for a classified dimension returns an error rather than fabricating pole labels.

## Tests

- Table test all canonical threshold boundary values and representative interior values.
- Table test positive, negative, and zero letter/pole behavior.
- Test pole labels come from question bank dimension metadata.
- Test fixed type order with an intentionally shuffled input map.
- Test a partially balanced type string such as `EXTP`.
- Test zero-score result emits `X`, `balanced`, and `Balanced=true`.
- Test missing threshold and uncovered score errors.

## Write Set

- `internal/scoring/classifier_test.go`
- `internal/scoring/classifier.go`
- `openspec/changes/core-assessment/specs/scoring-classifier/spec.md`
- `openspec/changes/core-assessment/tasks.md`
- `docs/plans/2026-04-11-core-assessment-design-task.md`
- `workspace/t07/*.md`
