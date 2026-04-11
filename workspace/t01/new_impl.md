# T01 Proposed Implementation

## Package

Add `internal/questionbank`.

## Types

- `Bank`
- `Meta`
- `Dimension`
- `DimensionMeta`
- `ScoringMeta`
- `Strength`
- `ThresholdRange`
- `Question`
- `LocalizedText`
- `Option`

## Scope Boundaries

- Include JSON tags matching the v3 bank.
- Use string aliases for dimensions and strengths so later validator/scoring tasks can share constants.
- Preserve `reverse` as a bool field.
- Do not add file loading, validation, scoring, or CLI wiring in `T01`.

## Test Shape

Add a package-level test that reads `questions/mbti-questions-v3.json`, unmarshals into `Bank`, and asserts representative fields across meta, dimensions, thresholds, questions, scenarios, options, score, and reverse.
