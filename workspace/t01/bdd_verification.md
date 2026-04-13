# T01 BDD Verification

## Negative Control

Temporary change:

- Changed the expected canonical title in `internal/questionbank/model_ginkgo_test.go`
  from `AI Behavioral Style Assessment v3` to `AI Behavioral Style Assessment v4`.

Command:

```bash
go test -v ./internal/questionbank -run TestQuestionbank -count=1 -ginkgo.v
```

Observed result:

- Failed on the title behavior assertion in
  `internal/questionbank/model_ginkgo_test.go`.
- Restored the expected title before green verification.

## Green Verification

Commands:

```bash
make fmt
go test -v ./internal/questionbank -run TestQuestionbank -count=1 -ginkgo.v
go test -count=1 ./internal/questionbank
go test -count=1 ./...
```

Observed results:

- Focused Ginkgo: 1 spec passed, 0 failed.
- Package test: `ok github.com/Miss-you/mbti-cli/internal/questionbank`.
- Full Go test: all packages passed.

## Fresh Verification - 2026-04-13

Commands:

```bash
go test -v ./internal/questionbank -run TestQuestionbank -count=1 -ginkgo.v
go test -count=1 ./internal/questionbank
go test -count=1 ./...
```

Observed results:

- Focused Ginkgo: 13 specs passed, 0 failed.
- Package test: `ok github.com/Miss-you/mbti-cli/internal/questionbank`.
- Full Go test: all packages passed.
