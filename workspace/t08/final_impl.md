# T08 Final Implementation

Status: accepted.

## Plan

Add `internal/result` as the single owner of result DTOs and renderer formatting.

Public API:

```go
type Meta struct {
    Title    string `json:"title"`
    Version  string `json:"version"`
    Answered int    `json:"answered"`
    Total    int    `json:"total"`
}

type Dimension struct {
    Letter   string `json:"letter"`
    Score    int    `json:"score"`
    Strength string `json:"strength"`
    Pole     string `json:"pole"`
    Balanced bool   `json:"balanced"`
}

type Dimensions struct {
    EI Dimension `json:"EI"`
    SN Dimension `json:"SN"`
    TF Dimension `json:"TF"`
    JP Dimension `json:"JP"`
}

type Summary struct {
    Meta       Meta       `json:"meta"`
    Type       string     `json:"type"`
    Dimensions Dimensions `json:"dimensions"`
}

func NewSummary(bank questionbank.Bank, score scoring.Result, classification scoring.Classification) (Summary, error)
func RenderJSON(summary Summary) ([]byte, error)
func RenderText(summary Summary) string
```

Implementation rules:

- Build dimensions in fixed `EI`, `SN`, `TF`, `JP` order.
- Return a contextual error if any required dimension classification is missing.
- Convert balanced dimensions to `Strength: "balanced"` for output stability.
- Keep `Balanced` in JSON for machine consumers instead of relying on `letter == "X"` inference.
- Use `json.MarshalIndent(summary, "", "  ")` and append a trailing newline.
- Render text as a factual summary with one row per dimension.

## Write Set

- `internal/result/result_test.go`
- `internal/result/result.go`
- `openspec/changes/core-assessment/specs/result-renderers/spec.md`
- `openspec/changes/core-assessment/tasks.md`
- `docs/plans/2026-04-11-core-assessment-design-task.md`
- `workspace/t08/*.md`

## Review

- First review: 82/100, two high-severity issues fixed.
- Second review: 97/100, no high-severity issues.
- Code review: no must-fix issues; accepted residual trust boundary that `internal/result` expects scorer/classifier-produced inputs.

## Verification

- First gate: `go test -count=1 ./internal/result`
- Renderer golden tests must assert exact JSON and text strings for a representative classified result:
  - meta title/version/answered/total
  - type string
  - dimensions in `EI`, `SN`, `TF`, `JP` order
  - a balanced dimension rendered with `letter: "X"`, `strength: "balanced"`, `pole: "balanced"`, and `balanced: true`
  - text output without `diagnosis`, `personality`, or `you are`
- Related gate: `go test -count=1 ./...`
- Repo gates: `make fmt`, `make test`, `make lint`, `make build`
- OpenSpec: `openspec validate core-assessment`
- No CLI smoke test for T08 because Cobra commands are out of scope.

## Golden Sample

JSON renderer output for the canonical mixed case must be exactly:

```json
{
  "meta": {
    "title": "AI Behavioral Style Assessment v3",
    "version": "0.3.0",
    "answered": 70,
    "total": 70
  },
  "type": "EXTP",
  "dimensions": {
    "EI": {
      "letter": "E",
      "score": 8,
      "strength": "moderate_a",
      "pole": "E (Expansive)",
      "balanced": false
    },
    "SN": {
      "letter": "X",
      "score": 0,
      "strength": "balanced",
      "pole": "balanced",
      "balanced": true
    },
    "TF": {
      "letter": "T",
      "score": 1,
      "strength": "slight_a",
      "pole": "T (Analytical)",
      "balanced": false
    },
    "JP": {
      "letter": "P",
      "score": -2,
      "strength": "slight_b",
      "pole": "P (Flexible)",
      "balanced": false
    }
  }
}
```

Text renderer output for the same case must be exactly:

```text
AI Behavioral Style Assessment v3 (v0.3.0)
Type: EXTP
Answered: 70/70

Dimensions:
- EI: E, score 8, strength moderate_a, pole E (Expansive)
- SN: X, score 0, strength balanced, pole balanced
- TF: T, score 1, strength slight_a, pole T (Analytical)
- JP: P, score -2, strength slight_b, pole P (Flexible)
```
