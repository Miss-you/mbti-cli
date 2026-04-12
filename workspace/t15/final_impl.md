# T15 Final Implementation

## Task

Align `docs/task.md` with the completed and archived core assessment phase 1 work.

## Implementation

1. Rewrite the document purpose so it remains useful after implementation has started.
2. Add a current-status section that names the completed phase 1 surface:
   - canonical v3 question bank schema
   - non-interactive `questions` command
   - canonical answer file and `score` command
   - stable text/JSON result rendering
   - fixtures, golden tests, README wording, and OpenSpec main specs
3. Add `Status` and `Decision Source` columns to the research task table.
4. Mark R03/R04/R06/R07/R08/R09 as resolved for phase 1.
5. Keep R01/R02/R05 as deferred for the future interactive questionnaire / anti-gaming phase.
6. Update recommended order and current bias so the doc points to remaining work rather than already completed phase 1 decisions.

## Scope

Docs-only. No Go code, question JSON, command behavior, output contract, or OpenSpec spec changes.

## Verification

- `rg` checks for `resolved`, `deferred`, phase 1 references, and decision sources.
- `make fmt`
- `make test`
- `make lint`
- `make build`
