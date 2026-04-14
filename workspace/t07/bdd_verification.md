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

## Follow-up Review Verification

Adopted review finding:

- Add explicit BDD coverage that zero-score classification succeeds without a
  configured threshold bucket.

Focused command:

```bash
go test -v ./internal/scoring -run TestScoring -count=1 -ginkgo.v
```

Observed result:

- Focused Ginkgo: 19 specs passed, 0 failed.
- The new scenario
  `Scoring classifier when a dimension score is zero and no threshold bucket is configured classifies the dimension as balanced without a strength lookup`
  passed.
- Package result:
  `ok github.com/Miss-you/mbti-cli/internal/scoring 0.242s`.
