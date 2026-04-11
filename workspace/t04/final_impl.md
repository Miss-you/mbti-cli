# T04 Final Implementation

## Scope

Implement a focused answer-file parser for the first-phase canonical map form.

## Code Changes

- Add package `internal/answers`.
- Add model types:
  - `File` for JSON shape.
  - `Set` for normalized `questionID -> optionCode`.
- Add `Parse(data []byte) (Set, error)`.

## Parser Contract

- Input must be a JSON object containing an `answers` object.
- The parser returns each answer keyed by the input question ID.
- Each option code is normalized with whitespace trimming and uppercase conversion.
- Malformed JSON, missing `answers`, null `answers`, and structurally invalid answer values return errors with answer-file context.

## Boundaries

- T04 does not import `internal/questionbank`.
- T04 does not check whether IDs exist in a bank.
- T04 does not check whether an option code is valid for a question.
- T04 does not enforce complete answers.
- T04 does not load files from disk.

## Review Disposition

- Plan review score: 89/100.
- No high-severity issues.
- Review note accepted in structure: `go test -count=1 ./internal/answers` is the T04 first gate.
- Spec/test-strategy review found verification-scope drift. Resolution: T04 acceptance and test strategy are scoped to `internal/answers`.
