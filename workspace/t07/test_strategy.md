# T07 Test Strategy

Status: accepted.

## What Must Be Proven

1. Every configured threshold bucket classifies inclusive boundary scores correctly.
2. Positive scores use pole A metadata and first dimension letters.
3. Negative scores use pole B metadata and second dimension letters.
4. Zero scores are explicitly balanced, use pole `balanced`, letter `X`, and do not require a threshold bucket.
5. Type strings are generated in fixed `EI`, `SN`, `TF`, `JP` order, including partial `X` cases.
6. Classifier errors are contextual for missing thresholds, uncovered scores, and missing dimension metadata.
7. T07 does not change raw scoring, renderer behavior, or CLI behavior.

## Focused Tests

- `go test -count=1 ./internal/scoring`
- Table-driven classifier tests in `internal/scoring/classifier_test.go`.
- Boundary rows for `strong_a`, `moderate_a`, `slight_a`, `slight_b`, `moderate_b`, and `strong_b`.
- Zero and partial-balanced type rows, including a type such as `EXTP`.
- A canonical v3 bank contract row that loads `questions/mbti-questions-v3.json`, validates it, and classifies representative raw scores with real pole metadata.
- Error rows for missing threshold coverage and missing dimension metadata.

## Repo Gates

- `make fmt`
- `make test`
- `make lint`
- `make build`
- `openspec validate core-assessment`

No CLI smoke test is required for T07 because no Cobra command behavior changes.
