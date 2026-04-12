# T03 BDD Verification

## Negative Control

Temporary change:

- Changed the expected missing-threshold issue in
  `internal/questionbank/validator_ginkgo_test.go` from
  `missing threshold slight_a` to `missing threshold slight_b`.

Command:

```bash
go test -v ./internal/questionbank -run TestQuestionbank -count=1 -ginkgo.v
```

Observed result:

- Failed on the threshold validation issue assertion in
  `internal/questionbank/validator_ginkgo_test.go`.
- Restored the expected issue text before green verification.

## Green Verification

Commands:

```bash
make fmt
go test -v ./internal/questionbank -run TestQuestionbank -count=1 -ginkgo.v
go test -count=1 ./internal/questionbank
go test -count=1 ./...
```

Observed results:

- Focused Ginkgo: 13 specs passed, 0 failed.
- Package test: `ok github.com/Miss-you/mbti-cli/internal/questionbank`.
- Full Go test: all packages passed.
