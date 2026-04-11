# T02 Final Implementation

## Decision

Implement a focused filesystem loader in `internal/questionbank` using `LoadFile(path) (LoadedBank, error)`.

## Product Writes

- `internal/questionbank/loader.go`
- `internal/questionbank/loader_test.go`

## Coordination Writes

- `workspace/t02/*.md`
- `openspec/changes/core-assessment/*`
- `docs/plans/2026-04-11-core-assessment-design-task.md`

## Plan

1. Add failing loader tests first for valid, missing, malformed, and empty-path cases.
2. Implement `Source`, `LoadedBank`, and `LoadFile`.
3. Keep schema validation out of scope.
4. Update the active `core-assessment` OpenSpec delta and tasks.
5. Run the focused first gate `go test -count=1 ./internal/questionbank`, then full delivery gates.

## Acceptance Evidence To Produce

- Valid canonical v3 bank loads from a file path.
- Loaded result exposes typed `Bank` data.
- Loaded result exposes source path, base filename, and byte size.
- Missing file returns a clear, wrapped read error.
- Malformed JSON returns a clear parse error.
- Empty path returns a clear required-path error.

## Review Result

Independent plan/spec review pass score: 91/100.

No high-severity plan or spec issues remain.

Code review pass score: 94/100.

No must-fix issues found. The loader preserves T02 scope, wraps missing-file errors, separates malformed JSON from later schema validation, and has tests for every task acceptance case.

## Verification Result

Fresh verification run:

- `make fmt`
- `go test ./internal/questionbank` after the initial TDD implementation
- `go test -count=1 ./internal/questionbank`
- `make test`
- `go test -count=1 ./...`
- `make lint`
- `make build`
- `openspec validate core-assessment`

The global `golangci-lint` binary is now v2.11.4, which matches this repository's v2 config.

## OpenSpec Close Decision

`core-assessment` remains active instead of archived because the task board uses that change name for subsequent first-phase tasks. T02's OpenSpec tasks are complete and the change validates.
