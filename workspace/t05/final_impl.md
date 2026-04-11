# T05 Final Implementation

## Scope

Implement strict, bank-aware answer validation for parsed or already-normalized
in-memory answer sets.

## API

Add:

```go
func Validate(bank questionbank.Bank, answerSet Set) error
```

and:

```go
type ValidationError struct {
    Issues []string
}
```

The error string starts with `answer validation failed` and includes all issues
joined in deterministic order.

## Behavior

- Unknown answer IDs are invalid.
- For known questions, selected option codes must match one option code on that
  question.
- Strict mode requires one answer for every bank question.
- Missing answers are reported in bank question order.
- Unknown and invalid answer entries are reported in sorted answer-ID order.
- The validator does not mutate the answer set.
- The validator does not normalize option codes; `Parse` remains responsible
  for file normalization, and in-memory callers must pass canonical option
  codes.
- The validator does not apply `reverse`, aggregate scores, classify
  thresholds, render output, or wire CLI behavior.

## Review

Independent plan review score: 88/100.

High-severity issues: none.

Reviewer suggestion accepted: the final plan explicitly states that validation
does not canonicalize option codes.

## Tests

Use TDD in `internal/answers`:

1. Add a failing success test with a small in-memory bank and complete answers.
2. Add failing cases for unknown ID, invalid option, and missing answer.
3. Add one aggregation/determinism check so callers get all strict-mode issues
   at once.

First gate:

```bash
go test -count=1 ./internal/answers
```

Full task gates after implementation:

```bash
make fmt
make test
make lint
make build
openspec validate core-assessment
```
