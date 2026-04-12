# T11 Original Implementation

Status: researched from `origin/main@3885b36`.

## Current Behavior

- `internal/questionbank` defines the typed v3 bank model, file loader, and schema validator. Existing tests load the canonical v3 bank and use in-memory mutations for invalid schema cases.
- `internal/answers` parses canonical `{"answers": {...}}` files, normalizes option codes, and validates strict answer coverage against a bank. Existing tests mostly use inline JSON and synthetic banks.
- `internal/scoring` validates answers before aggregation, sums signed option scores by dimension, and intentionally does not apply `reverse` a second time.
- `internal/result` builds result DTOs and renders stable JSON/text. Existing golden assertions are inline string literals.
- `cmd/mbti-cli` wires `questions` and `score` as thin Cobra orchestration layers. Existing CLI tests generate answer files in temp dirs and assert selected fields or substrings.

## Coverage Gaps

- No committed `testdata` fixture corpus exists for valid/invalid bank shapes, answer files, scoring inputs, or CLI score inputs.
- Golden output is not stored as standalone fixtures, so renderer and CLI output contracts are harder to inspect and update deliberately.
- CLI tests prove parseability and key fields, but not exact stable stdout for representative success paths.
- Existing schema and answer validation coverage is good, but it is not file-backed in the way T11 asks for.

## Likely Write Set

- Add package-local `testdata` fixtures under `internal/questionbank`, `internal/answers`, `internal/scoring`, `internal/result`, and `cmd/mbti-cli`.
- Update tests to load fixtures and compare stable output with golden files.
- Update `openspec/changes/core-assessment/tasks.md` with T11 test-asset tasks.
- Do not change non-test production behavior.

## Risks

- Golden files will intentionally be sensitive to JSON field order, indentation, trailing newlines, and text formatting.
- Fixture paths must stay package-local and simple so `go test ./...` works from any package.
- Fixtures should include at least one reversed question so a future accidental double-application of `reverse` is caught.
