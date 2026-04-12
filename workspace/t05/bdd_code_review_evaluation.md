# T05 BDD Code Review Evaluation

## Adopted

- Trim the shared bank fixture to the smallest shape the validator actually reads.
  Reason: `Validate` only inspects question IDs and option codes, so `Meta`, `Scenario`, `Reverse`, and `Score` add noise without increasing behavioral coverage.
- Keep the lowercase/non-canonical input case as a validator-level negative control.
  Reason: the spec explicitly requires validation not to normalize option codes itself, so this remains a valid boundary check.
- Keep the aggregated-issues scenario in place.
  Reason: deterministic multi-issue ordering is observable validator behavior and is worth proving in BDD.

## Rejected

- None.

## Verification

- `go test -v ./internal/answers -run TestAnswers -count=1 -ginkgo.v`
- Result: passed (`12 Passed`, `0 Failed`)
