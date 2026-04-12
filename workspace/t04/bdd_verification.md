# T04 BDD Verification

## Decision

No additional BDD code was added for T04.

Review and evaluation found that `internal/answers/parser_ginkgo_test.go`
already covers the parser-level OpenSpec behavior:

- map-form answer JSON parses and normalizes option codes
- missing or null `answers` is rejected
- malformed JSON is rejected with parser context
- non-string answer values are rejected
- bank-aware validation is deferred

## Negative Control

No new BDD scenario was added, so no additional negative control was required.

## Green Verification

Commands:

```bash
make fmt
go test -v ./internal/answers -run TestAnswers -count=1 -ginkgo.v
go test -count=1 ./internal/answers
go test -count=1 ./...
```

Observed results:

- Focused Ginkgo: 6 specs passed, 0 failed.
- Package test: `ok github.com/Miss-you/mbti-cli/internal/answers`.
- Full Go test: all packages passed.
