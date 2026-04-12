# T05 BDD Verification

## Negative Control

Temporary change:

- Changed the expected invalid-option issue in
  `internal/answers/validator_ginkgo_test.go` from
  `question q01 option Z is not valid` to
  `question q01 option Y is not valid`.

Command:

```bash
go test -v ./internal/answers -run TestAnswers -count=1 -ginkgo.v
```

Observed result:

- Failed on the invalid option validation issue assertion in
  `internal/answers/validator_ginkgo_test.go`.
- Restored the expected issue text before green verification.

## Green Verification

Commands:

```bash
make fmt
go test -v ./internal/answers -run TestAnswers -count=1 -ginkgo.v
go test -count=1 ./internal/answers
go test -count=1 ./...
```

Observed results:

- Focused Ginkgo: 12 specs passed, 0 failed.
- Package test: `ok github.com/Miss-you/mbti-cli/internal/answers`.
- Full Go test: all packages passed.
