# T02 BDD Verification

## Negative Control

Temporary change:

- Changed the expected source filename in
  `internal/questionbank/loader_ginkgo_test.go` from
  `mbti-questions-v3.json` to `mbti-questions-v4.json`.

Command:

```bash
go test -v ./internal/questionbank -run TestQuestionbank -count=1 -ginkgo.v
```

Observed result:

- Failed on the source filename behavior assertion in
  `internal/questionbank/loader_ginkgo_test.go`.
- Restored the expected filename before green verification.

## Green Verification

Commands:

```bash
make fmt
go test -v ./internal/questionbank -run TestQuestionbank -count=1 -ginkgo.v
go test -count=1 ./internal/questionbank
go test -count=1 ./...
```

Observed results:

- Focused Ginkgo: 5 specs passed, 0 failed.
- Package test: `ok github.com/Miss-you/mbti-cli/internal/questionbank`.
- Full Go test: all packages passed.
