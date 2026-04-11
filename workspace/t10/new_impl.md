# T10 Alternative Implementation Notes

## Option A: Thin Score Command

Add `cmd/mbti-cli/score.go` with local options:

- `--questions <file>` required
- `--answers <file>` required
- `--format text|json`, default `text`

Flow:

1. Validate flag-only inputs.
2. Load and validate question bank.
3. Read and parse answer file.
4. Run `scoring.Score`.
5. Run `scoring.Classify`.
6. Build `result.Summary`.
7. Render JSON or text to stdout.

This follows the `questions` command shape and keeps core behavior in existing packages.

## Option B: Shared CLI Pipeline Helper

Extract shared bank loading and format helpers used by `questions` and `score`.

Rejected for T10 because there are only two commands, the duplication is small, and shared helpers would obscure the command-specific error context before there is enough repetition.

## Preferred

Use Option A. Add focused command tests first, then implement minimal production code.
