# T15 Final Implementation v1

## Task

Align `docs/task.md` with the completed and archived core assessment phase 1 work.

## Edits

1. Rewrite the purpose text so the document is a historical/current research index rather than a purely pre-implementation checklist.
2. Add a "Current Status" section linking phase 1 decisions to the core assessment design, task board, README, and OpenSpec main specs.
3. Add `Status` and `Decision Source` columns to the research task table.
4. Mark phase 1-resolved items and keep remaining items clearly deferred.
5. Update the recommended order and current bias sections so they no longer contradict the implemented `questions` / `score` workflow.

## No OpenSpec Change

No OpenSpec change is required because this task does not change product behavior, CLI contracts, question/scoring semantics, or user-visible output.

## Verification

- `rg` checks for resolved/deferred markers and references.
- `make fmt`
- `make test`
- `make lint`
- `make build`
