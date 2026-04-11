# T10 Original Implementation

Status: researched from `origin/main` 6106fda.

## Current State

- `cmd/mbti-cli/root.go` wires `version` and `questions`, but not `score`.
- `cmd/mbti-cli/questions.go` already establishes the CLI pattern:
  load a question bank, validate it, render to `cmd.OutOrStdout()`, and return errors without partial stdout.
- `internal/answers` parses canonical answer files in the map shape:
  `{"answers":{"q01":"A"}}`.
- `internal/answers.Validate` enforces strict complete answers against a loaded bank.
- `internal/scoring.Score` aggregates signed option scores and calls strict answer validation.
- `internal/scoring.Classify` maps dimension scores to strengths, poles, and type letters.
- `internal/result` builds stable result DTOs and renders newline-terminated JSON or factual text.

## Missing Surface

There is no `mbti-cli score` command. Users cannot yet run:

```bash
mbti-cli score --questions questions/mbti-questions-v3.json --answers answers.json --format json
```

## Constraints

- Keep Cobra code as a thin shell.
- Do not duplicate answer validation, scoring, classification, or rendering logic in `cmd/mbti-cli`.
- Keep successful output on stdout.
- Preserve command tests as the first gate for CLI behavior.
