# T07 BDD Verification

## Negative Control

Temporary change:

- Changed the expected mixed-score type in
  `internal/scoring/classifier_ginkgo_test.go` from `EXTP` to `EXTJ`.

Command:

```bash
go test -v ./internal/scoring -run TestScoring -count=1 -ginkgo.v
```

Observed result:

- Failed on the deterministic type assertion in
  `internal/scoring/classifier_ginkgo_test.go`.
- Restored the expected type before green verification.

## Green Verification

Commands:

```bash
go test -v ./internal/scoring -run TestScoring -count=1 -ginkgo.v
go test -count=1 ./internal/scoring
go test -count=1 ./...
```

Observed results:

- Focused Ginkgo: 18 specs passed, 0 failed.
- Package test: `ok github.com/Miss-you/mbti-cli/internal/scoring`.
- Full Go test: all packages passed.
