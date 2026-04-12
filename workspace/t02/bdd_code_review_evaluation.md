# T02 BDD Code Review Evaluation

## Adopted

- Remove the exact `HaveLen(70)` assertion from the canonical load scenario.
  Reason: the T02 plan only calls for representative typed-bank fields and source metadata. The exact count is fixture-specific and pulls the spec toward regression coverage of the JSON payload instead of loader behavior.
- Keep the remaining canonical assertions on title, version, first question ID, and source metadata.
  Reason: these still prove the loader contract without binding the BDD to the full fixture shape.
- Keep the missing-file, malformed-JSON, and empty-path scenarios unchanged.
  Reason: they already match the loader read/parse/path contract and stay within the planned scope.

## Rejected

- None.

## Verification

Planned focused verification:

```bash
go test -v ./internal/questionbank -run TestQuestionbank -count=1 -ginkgo.v
```

Result:

- Passed.
