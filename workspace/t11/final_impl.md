# T11 Final Implementation

Status: accepted after independent review.

Review result: 94/100, no high-severity issues after adding explicit threshold-boundary and zero/balanced coverage.

## Accepted Direction

Implement T11 as a fixture and golden-test task only. The code will pin current contracts with committed test assets and will not change scoring, validation, rendering, or CLI behavior.

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

## Boundaries

- No production code changes.
- No new runtime dependencies.
- No new product behavior spec.
- T12 full verification repair remains out of scope.
