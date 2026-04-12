# T06 BDD Verification

## Negative Control

Temporary change:

- Changed the expected reversed-question dimension score in
  `internal/scoring/scoring_ginkgo_test.go` from `-2` to `+2`.

Command:

```bash
go test -v ./internal/scoring -run TestScoring -count=1 -ginkgo.v
```

Observed result:

- Failed on the reversed-question score assertion in
  `internal/scoring/scoring_ginkgo_test.go`.
- Restored the expected score before green verification.

## Green Verification

Commands:

```bash
make fmt
go test -v ./internal/scoring -run TestScoring -count=1 -ginkgo.v
go test -count=1 ./internal/scoring
go test -count=1 ./...
```

Observed results:

- Focused Ginkgo: 3 specs passed, 0 failed.
- Package test: `ok github.com/Miss-you/mbti-cli/internal/scoring`.
- Full Go test: all packages passed.
