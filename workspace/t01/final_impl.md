# T01 Final Implementation

## Decision

Implement only the question bank data model in `internal/questionbank`, plus a focused unmarshal contract test against the canonical v3 bank.

## Product Writes

- `internal/questionbank/model.go`
- `internal/questionbank/model_test.go`
- `questions/mbti-questions-v3.json`

`questions/mbti-questions-v3.json` is a required contract fixture for T01 because the approved design and done condition explicitly name it. The model implementation remains limited to `internal/questionbank`.

## Coordination Writes

- `workspace/t01/*.md`
- `openspec/changes/core-assessment/*`
- `docs/plans/2026-04-11-core-assessment-design-task.md`

These files keep delivery state aligned and do not expand product behavior.

## Plan

1. Add only the canonical v3 question bank file needed by the unmarshal contract test.
2. Write the failing `go test ./internal/questionbank` test first.
3. Add minimal structs and constants needed for JSON unmarshalling.
4. Keep validation out of scope. No count checks, threshold coverage checks, option validation, loader behavior, scoring, rendering, or CLI wiring.
5. Run `go test ./internal/questionbank` as the first gate, then the broader gates required by the delivery workflow.

## Acceptance Evidence

The test must prove:

- v3 title/version/total are preserved.
- all 70 questions unmarshal.
- dimension metadata for `EI`, `SN`, `TF`, and `JP` is addressable.
- scoring thresholds unmarshal as two-value ranges.
- question `id`, `dimension`, `reverse`, localized scenario, option code/label/score all survive unmarshal.

## Review Result

Independent plan review score: 83/100. No high-severity plan issues remain.

Independent spec review score: 92/100. No high-severity spec or test-strategy issues.

Independent code review initially found one must-fix test coverage issue. It was fixed by adding a full raw JSON round-trip comparison and exact checks for every scoring threshold bucket. Re-review score: 97/100. No must-fix issues remain.

## Verification Result

Fresh verification run after review repair:

- `make fmt`
- `go test -count=1 ./internal/questionbank`
- `go test -count=1 ./...`
- `make test`
- `PATH=/tmp/mbti-cli-tools/bin:$PATH make lint`
- `make build`
- `openspec validate core-assessment`

`make lint` requires golangci-lint v2 for this repository's config. The existing default binary was v1, so verification used a temporary v2 binary installed under `/tmp/mbti-cli-tools/bin`.

## OpenSpec Close Decision

`core-assessment` remains active instead of archived because the task board uses that change name for subsequent first-phase tasks. T01's OpenSpec artifacts and tasks are complete and valid.
