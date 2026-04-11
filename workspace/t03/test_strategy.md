# T03 Test Strategy

## Behavior to prove

- The canonical v3 question bank passes schema validation.
- Invalid bank schemas fail with clear issue text.
- Validation covers metadata totals, supported dimension metadata, dimension counts, question IDs, localized scenarios, option codes, option scores, localized labels, and complete ordered non-overlapping threshold buckets.
- Validation remains static schema validation only; it does not load files, validate answers, score results, or apply `reverse`.

## Focused tests

- `TestValidateAcceptsCanonicalV3QuestionBank`
- `TestValidateRejectsInvalidQuestionBankSchemas`

## Verification commands

- `go test -count=1 ./internal/questionbank`
- `make fmt`
- `make test`
- `make lint`
- `make build`
- `openspec validate core-assessment`
