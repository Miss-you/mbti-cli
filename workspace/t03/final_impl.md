# T03 Final Implementation

## Scope

T03 adds schema validation for a typed `questionbank.Bank`.

## API

- `Validate(bank Bank) error`
- `ValidationError` with deterministic issue text and accessible `Issues []string`

## Validation rules

- `meta.total` must equal the number of questions.
- Metadata must define exactly the supported dimensions `EI`, `SN`, `TF`, and `JP`.
- Each metadata dimension count must match the actual question count.
- Questions must have non-empty unique IDs, supported dimensions, localized scenario text, and exactly four options.
- Options must use codes `A`, `B`, `C`, and `D` exactly once, include localized labels, and use score values `-2`, `-1`, `1`, or `2`.
- Thresholds must define all six supported strengths; each range must be ordered and must not overlap another supported threshold range.

## Non-goals

- `LoadFile` remains structural loading only.
- `Validate` does not score answers, classify thresholds, or apply `reverse`.
- No CLI behavior changes.

## Verification

- First gate: `go test -count=1 ./internal/questionbank`
- Full relevant gates: `make fmt`, `make test`, `make lint`, `make build`, `openspec validate core-assessment`
