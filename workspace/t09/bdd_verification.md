# T09 BDD Verification

## Negative Control

Temporary change:

- Changed the expected JSON export count in
  `cmd/mbti-cli/questions_command_ginkgo_test.go` from `70` to `69`.

Command:

```bash
go test -v ./cmd/mbti-cli -run TestCLI -count=1 -ginkgo.v
```

Observed result:

- Failed on the questions command JSON metadata assertion in
  `cmd/mbti-cli/questions_command_ginkgo_test.go`.
- Restored the expected count before green verification.

## Green Verification

Commands:

```bash
go test -v ./cmd/mbti-cli -run TestCLI -count=1 -ginkgo.v
go test -count=1 ./cmd/mbti-cli
go test -count=1 ./...
```

Observed results:

- Focused CLI Ginkgo: 21 specs passed, 0 failed.
- Package test: `ok github.com/Miss-you/mbti-cli/cmd/mbti-cli`.
- Full Go test: all packages passed.
