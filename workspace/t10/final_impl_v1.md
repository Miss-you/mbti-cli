# T10 Final Implementation v1

## Plan

Implement `mbti-cli score` as a thin Cobra command over existing core packages.

## Behavior

- `score --questions <file> --answers <file> --format json` writes stable parseable result JSON to stdout.
- `score --questions <file> --answers <file> --format text` writes the existing factual text summary to stdout.
- `--format` accepts only `text` and `json`; default is `text`.
- Missing question bank path, missing answer path, missing answer file, malformed answer JSON, and strict answer validation errors return errors and write no command stdout.
- JSON and text rendering semantics remain owned by `internal/result`.

## Write Set

- `cmd/mbti-cli/root.go`
- `cmd/mbti-cli/score.go`
- `cmd/mbti-cli/score_test.go`
- `openspec/changes/core-assessment/tasks.md`
- `openspec/changes/core-assessment/specs/score-command/spec.md`
- `docs/plans/2026-04-11-core-assessment-design-task.md`
- `workspace/t10/*.md`

## Review Pass

Score: 93/100.

- CLI contract fidelity: 24/25. The command matches the approved first-phase non-interactive score contract.
- question/scoring semantics: 25/25. The command delegates to existing strict answer validation, score aggregation, classifier, and renderer.
- Go-native maintainability: 18/20. A direct Cobra command is simple; avoid premature shared helper extraction.
- Scope control: 14/15. No interactive assess flow, persistence, provider abstraction, or report generation.
- Verification clarity: 12/15. Needs command tests, repo gates, OpenSpec validation, and a focused CLI smoke that parses JSON.

No high-severity issues found.
