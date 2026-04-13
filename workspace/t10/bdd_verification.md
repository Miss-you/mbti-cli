# T10 BDD Verification

## Negative Control

T10 shares the `cmd/mbti-cli` Ginkgo suite with T09. The CLI negative control
changed the expected questions JSON count from `70` to `69`, proving the suite
fails on a command-output assertion before restoration.

Command:

```bash
go test -v ./cmd/mbti-cli -run TestCLI -count=1 -ginkgo.v
```

Observed result:

- Failed on the CLI behavior assertion.
- Restored the expected value before green verification.

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
