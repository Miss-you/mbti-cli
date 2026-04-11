# T06 Original Implementation

Status: research complete.

## Current Repository Facts

- `internal/questionbank` defines the typed bank model, supported dimensions, strength bucket names, threshold ranges, and signed option scores.
- `internal/questionbank.Validate(bank)` enforces schema correctness, including option scores in `-2/-1/1/2` and complete non-overlapping thresholds.
- `internal/answers.Parse(data)` returns `answers.Set`, normalizing option codes to uppercase.
- `internal/answers.Validate(bank, answers)` enforces strict answer completeness and option validity against the bank.
- There is no `internal/scoring` package yet.
- The active OpenSpec change is `core-assessment`; its task file currently covers T01-T05 and reports complete.

## T06 Gap

The repository has the validated inputs needed for scoring, but no Go API that turns a validated bank and answer set into deterministic per-dimension totals. `reverse` is preserved in the question model but no existing code applies it.
