# T06 Final Implementation v1

Status: reviewed.

## Plan

Implement `internal/scoring` as a narrow aggregation package.

Public API:

```go
type Result struct {
    Answered        int
    Total           int
    DimensionScores map[questionbank.Dimension]int
}

func Score(bank questionbank.Bank, answerSet answers.Set) (Result, error)
```

Behavior:

- Strict answer validation runs before aggregation by delegating to `internal/answers.Validate`.
- Each bank question contributes exactly the selected option's signed `Score` to its dimension.
- `Question.Reverse` is not applied during scoring.
- The result initializes `EI`, `SN`, `TF`, and `JP` totals, even if a small test bank does not include every dimension.
- If the selected option cannot be resolved after validation, return a contextual scoring error instead of silently skipping it.

Review pass:

- Score: 92/100.
- No high-severity findings.
- Main risk: T06 could drift into threshold/type behavior. Mitigation: tests and implementation only assert raw dimension totals and explicitly keep threshold classification out of scope.
