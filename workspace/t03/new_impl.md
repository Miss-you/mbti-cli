# T03 Candidate Implementation

## API

Add `func Validate(bank Bank) error` to `internal/questionbank`.

## Behavior

- Return `nil` for the canonical v3 question bank.
- Return a grouped validation error for one or more schema problems.
- Check:
  - `meta.total == len(questions)`
  - supported dimensions are exactly `EI`, `SN`, `TF`, `JP`
  - `meta.dimensions[dim].count` matches actual question counts
  - each question has a unique non-empty ID
  - each question has supported dimension and localized scenario text
  - each question has exactly four options with codes `A/B/C/D`
  - option labels have both languages and scores are `-2`, `-1`, `1`, or `2`
  - threshold buckets contain all six known strengths and each range is ordered

## Non-goals

- No file loading changes.
- No answer validation.
- No scoring or threshold classification.
- No CLI wiring.
