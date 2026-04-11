# T08 New Implementation

Status: research complete.

## Proposed Shape

Add `internal/result` with a small DTO and renderer API:

```go
type Summary struct {
    Meta       Meta       `json:"meta"`
    Type       string     `json:"type"`
    Dimensions Dimensions `json:"dimensions"`
}

func NewSummary(bank questionbank.Bank, score scoring.Result, classification scoring.Classification) (Summary, error)
func RenderJSON(summary Summary) ([]byte, error)
func RenderText(summary Summary) string
```

`Dimensions` should be a struct with `EI`, `SN`, `TF`, and `JP` fields rather than a map. That keeps JSON key order stable for golden tests and matches the fixed dimension order required by the scoring contract.

## Behavior

- Metadata comes from `bank.Meta.Title`, `bank.Meta.Version`, `score.Answered`, and `score.Total`.
- Type comes from `classification.Type`.
- Dimension rows come from `classification.Dimensions`.
- Non-balanced rows render their classifier strength as-is.
- Balanced rows render strength as the stable string `balanced`.
- JSON output is indented, stable, and newline-terminated.
- Text output is concise and factual. It describes score, strength, and pole, but does not claim a diagnosis or tell the user what their personality is.

## Boundaries

- No scoring, validation, threshold, or reverse-score behavior changes.
- No Cobra command wiring.
- No fixture expansion beyond renderer tests.
