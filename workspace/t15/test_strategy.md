# T15 Test Strategy

## What to prove

- `docs/task.md` clearly distinguishes phase 1-resolved research from deferred research.
- Resolved claims point to durable repository artifacts instead of chat-only context.
- The task remains docs-only: no product behavior or OpenSpec requirement changes.
- Existing Go gates remain green after the documentation update.

## Focused checks

- Use `rg` to confirm the document includes `resolved for phase 1`, `deferred`, `Decision Source`, `OpenSpec`, `README`, and `docs/plans/2026-04-11-core-assessment-design-task.md`.
- Use `git diff --name-only` to confirm only the task board, `docs/task.md`, and `workspace/t15/` changed.
- Run `make fmt`, `make test`, `make lint`, and `make build`.

## Explicitly skipped

- CLI smoke tests: skipped because this task does not change CLI code or command behavior.
- `openspec validate`: skipped because this task does not edit OpenSpec specs or active changes.
