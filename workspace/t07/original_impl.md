# T07 Original Implementation

Status: research complete.

## Current Repository Facts

- `internal/scoring.Score` returns raw per-dimension totals in `Result.DimensionScores`.
- `Score` delegates strict answer validation to `internal/answers.Validate`, initializes the four supported dimensions, and aggregates the selected option's signed score.
- `Score` intentionally does not apply `Question.Reverse` and does not classify thresholds or generate type letters.
- `internal/questionbank` already defines `Dimension`, `Strength`, and `ThresholdRange` types plus the canonical `EI`, `SN`, `TF`, and `JP` dimensions.
- The canonical v3 bank stores thresholds under `bank.Meta.Scoring.Thresholds`:
  - `strong_a`: `[13, 999]`
  - `moderate_a`: `[5, 12]`
  - `slight_a`: `[1, 4]`
  - `slight_b`: `[-4, -1]`
  - `moderate_b`: `[-12, -5]`
  - `strong_b`: `[-999, -13]`
- `questionbank.Validate` already verifies threshold completeness, ordering, and non-overlap for supported buckets.

## T07 Gap

The repository can produce raw dimension totals, but there is no Go API that maps those totals to threshold strengths, poles, letters, or the four-letter type string. Zero-score behavior is not yet represented in code.
