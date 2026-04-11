# T07 New Implementation

Status: research complete.

## Proposed Shape

Extend `internal/scoring` with classifier types and one narrow exported API:

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

`Classify` should:

1. Read threshold ranges from `bank.Meta.Scoring.Thresholds`.
2. Classify non-zero scores by finding the threshold range containing that score.
3. Treat score `0` as balanced without consulting thresholds.
4. Generate type letters in fixed order: `EI`, `SN`, `TF`, `JP`.
5. Use `bank.Meta.Dimensions[dim].PoleA` and the first dimension letter for positive scores.
6. Use `bank.Meta.Dimensions[dim].PoleB` and the second dimension letter for negative scores.
7. Use pole `balanced`, letter `X`, and `Balanced=true` for zero scores.
8. Return contextual errors when required thresholds are missing or a non-zero score matches no threshold.

## Boundaries

- No raw score aggregation changes.
- No JSON/text rendering DTOs. T08 owns renderer output shape.
- No Cobra command wiring. T10 owns CLI scoring.
