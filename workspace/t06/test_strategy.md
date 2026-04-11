# T06 Test Strategy

Status: accepted.

## What Must Be Proven

1. A strict valid answer set produces deterministic dimension totals.
2. Signed option scores are used directly.
3. `Question.Reverse` is not applied a second time.
4. Invalid or incomplete answers fail before aggregation through the existing answer validation contract.
5. T06 does not classify thresholds or generate type labels.

## Focused Tests

- `go test -count=1 ./internal/scoring`
- A fixture bank with multiple dimensions and a reversed question.
- A missing-answer case that expects the existing answer validation error.

## Repo Gates

- `make fmt`
- `make test`
- `make lint`
- `make build`
- `openspec validate core-assessment`

No CLI smoke test is required for T06 because no Cobra command behavior changes.
