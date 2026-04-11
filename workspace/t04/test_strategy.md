# T04 Test Strategy

## Behaviors To Prove

1. Canonical map-form answer JSON parses into `questionID -> optionCode`.
2. Option codes are normalized with trimming and uppercase conversion.
3. Malformed JSON reports answer-file parse context.
4. Missing or null `answers` is rejected before later validation.
5. Structurally invalid answer maps, such as non-string values, return clear parse errors.

## In-Scope Gates

- `go test -count=1 ./internal/answers`

## Out Of Scope

- Unknown question ID validation.
- Invalid option-code validation against a question bank.
- Missing required answer validation.
- Scoring or result rendering.
- Cobra command wiring or filesystem answer loading.
