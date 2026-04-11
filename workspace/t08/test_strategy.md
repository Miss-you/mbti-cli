# T08 Test Strategy

Status: accepted.

## What Must Be Proven

1. `internal/result` builds a summary from bank metadata, raw score counts, and scoring classification.
2. JSON output is exact, indented, newline-terminated, and ordered as `meta`, `type`, `dimensions`, then `EI`, `SN`, `TF`, `JP`.
3. Balanced dimensions render as `letter: "X"`, `strength: "balanced"`, `pole: "balanced"`, and `balanced: true`.
4. Non-balanced dimensions preserve classifier score, strength, pole, letter, and `balanced: false`.
5. Text output is readable and factual, with one row per dimension in fixed order.
6. Text output does not contain diagnostic or personality-claim wording such as `diagnosis`, `personality`, or `you are`.
7. Missing required dimension classifications return contextual errors instead of silently rendering zero values.
8. T08 does not change scoring, validation, question bank, or CLI behavior.

## Focused Tests

- `go test -count=1 ./internal/result`
- Exact JSON golden string for a mixed `EXTP` case with `SN` balanced.
- Exact text golden string for the same case.
- Error test for a missing dimension classification, for example missing `SN`.

## Golden JSON Case

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

## Golden Text Case

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

## Repo Gates

- `make fmt`
- `make test`
- `make lint`
- `make build`
- `openspec validate core-assessment`

No CLI smoke test is required for T08 because no Cobra command behavior changes.
