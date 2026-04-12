# T11 Final Implementation V1

Status: ready for review.

## Accepted Direction

Implement T11 as a fixture and golden-test task only. The code should pin current contracts with committed test assets and should not change scoring, validation, rendering, or CLI behavior.

## Scope

Add package-local fixtures and golden files for:

- schema validation: valid bank plus invalid meta total, duplicate question ID, and threshold overlap banks
- answer parsing and validation: valid complete answers plus an invalid file with unknown ID, invalid option, and missing answers
- scoring: a small four-question bank with a reversed question and deterministic answers
- threshold/classifier regression: fixture-backed table cases for strong/moderate/slight boundaries and one zero-total case that must classify as `balanced` with type letter `X`
- rendering: JSON and text golden outputs for the existing result summary test
- CLI: committed score answer fixtures and stable stdout goldens for representative `questions` and `score` success paths

## Implementation Steps

1. Add the fixture and golden files under `testdata` directories.
2. Add or update tests to load those files.
3. Keep existing table-driven validation tests where they add useful breadth.
4. Add T11 items to `openspec/changes/core-assessment/tasks.md`; do not add product behavior requirements.
5. Run the focused package gate, then repo gates.

## Expected Verification

- `go test -count=1 ./internal/questionbank`
- `go test -count=1 ./internal/answers ./internal/scoring ./internal/result ./cmd/mbti-cli`
- `go test -count=1 ./...`
- `make fmt`
- `make test`
- `make lint`
- `make build`
- `openspec validate core-assessment`

## Review Criteria

- Fixtures cover schema, answers, scoring, rendering, and CLI.
- Threshold fixtures cover bucket boundaries and zero/balanced behavior.
- Golden outputs are stable, readable, and newline-sensitive by design.
- No product behavior or CLI contract changes are introduced.
- The task remains limited to T11 and does not absorb T12 full verification repair.
