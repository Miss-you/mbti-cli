# T08 BDD Verification

## Negative Control

Temporary change:

- Changed the expected summary title in `internal/result/result_ginkgo_test.go`
  from `AI Behavioral Style Assessment v3` to
  `AI Behavioral Style Assessment v4`.

Command:

```bash
go test -v ./internal/result -run TestResult -count=1 -ginkgo.v
```

Observed result:

- Failed on the summary metadata assertion in
  `internal/result/result_ginkgo_test.go`.
- Restored the expected title before green verification.

## Green Verification

Commands:

```bash
go test -v ./internal/result -run TestResult -count=1 -ginkgo.v
go test -count=1 ./internal/result
go test -count=1 ./...
```

Observed results:

- Focused Ginkgo: 4 specs passed, 0 failed.
- Package test: `ok github.com/Miss-you/mbti-cli/internal/result`.
- Full Go test: all packages passed.
