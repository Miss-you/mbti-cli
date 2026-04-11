# T05 Alternative Implementation

## Option

Add `internal/answers.Validate(bank questionbank.Bank, answerSet answers.Set)`.

The validator should:

- build a question index from the bank
- build per-question option-code sets from each question
- report unknown answer IDs in deterministic sorted order
- report invalid option codes for known questions in deterministic sorted order
- report missing answers in bank question order
- return an aggregated `answers.ValidationError`

## Boundaries

- Do not apply `reverse`.
- Do not aggregate scores.
- Do not introduce strict/partial mode plumbing yet. T05 is strict mode only.
- Do not change parser behavior except as needed by tests, which is currently
  not needed.
- Do not wire Cobra commands.

## Why this shape

Keeping validation in `internal/answers` lets scoring consume a known-good
answer set without duplicating answer-file concerns. Importing
`internal/questionbank` from `internal/answers` does not create an import cycle.
