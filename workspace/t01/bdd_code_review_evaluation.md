# T01 BDD Code Review Evaluation

## Adopted

1. Adopt: remove the full raw JSON round-trip equality check.
   - Reason: the final T01 plan only needs representative field preservation, not a structural mirror of the entire fixture.

2. Adopt: trim the per-question assertions to the minimal representative set.
   - Reason: the adopted plan calls for one canonical first question, one reversed question, and one representative option label plus signed score. The extra last-question assertions were broader than that scope.

3. Adopt: keep the BDD title centered on the observable contract.
   - Reason: a behavior spec should read like a contract for canonical v3 unmarshaling, not like a generic field-preservation dump.

4. Adopt: align the focused verification path with the actual suite entrypoint.
   - Reason: the repository's suite function is `TestQuestionbank`, so the focused `go test -run` command must match that name to exercise the package correctly.

## Rejected

None.

## Verification

- Applied the adopted scope reduction in `internal/questionbank/model_ginkgo_test.go`.
- Focused verification passed with the actual suite entrypoint:

```bash
go test -v ./internal/questionbank -run TestQuestionbank -count=1 -ginkgo.v
```

- Result: PASS
