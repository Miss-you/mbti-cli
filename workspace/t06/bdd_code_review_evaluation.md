# T06 BDD Code Review Evaluation

## Adopted

- Adopt: add a local `scoringSpecQuestion` helper in `internal/scoring/scoring_ginkgo_test.go` so the Ginkgo spec is self-contained and reads as a behavior spec without leaning on `scoring_test.go` fixtures.

## Rejected

- None.

## Verification

- `go test -v ./internal/scoring -run TestScoring -count=1 -ginkgo.v`
- Result: `PASS` (`3 Passed`, `0 Failed`)
