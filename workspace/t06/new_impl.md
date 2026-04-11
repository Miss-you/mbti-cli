# T06 New Implementation

Status: research complete.

## Proposed Shape

Add a small `internal/scoring` package with:

- `Result`, containing answered count, total question count, and a `map[questionbank.Dimension]int` of dimension totals.
- `Score(bank questionbank.Bank, answerSet answers.Set) (Result, error)`.

`Score` should:

1. Call `answers.Validate(bank, answerSet)` so scoring never returns a partial result for invalid or incomplete strict answers.
2. Initialize supported dimensions to zero.
3. Iterate bank questions in bank order.
4. Find the selected option by canonical option code.
5. Add the option's signed `score` directly to that question's dimension total.
6. Never transform scores based on `question.Reverse`.

## Boundaries

- No threshold strength classification. That belongs to T07.
- No type string generation. That belongs to T07.
- No JSON/text rendering. That belongs to T08.
- No CLI wiring. That belongs to T10.
