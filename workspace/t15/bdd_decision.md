# T15 BDD Decision

## Decision

Dedicated BDD/Ginkgo coverage is not needed for T15.

T15 is a human research task-list alignment change in `docs/task.md`. It marks
which phase 1 research items are resolved by design, spec, implementation, and
README evidence, while preserving deferred research for later phases. It does
not alter Go behavior, OpenSpec requirements, CLI contracts, or question-bank
scoring semantics.

## Source Artifacts

- Task board: `docs/plans/2026-04-11-core-assessment-design-task.md`
- Workspace strategy: `workspace/t15/test_strategy.md`
- Documentation target: `docs/task.md`

## Handling

T15 should be verified with focused document checks and repository gates. A new
Ginkgo spec would not have a product behavior to execute.

Required evidence remains:

```bash
rg -n "resolved for phase 1|deferred|Decision Source|OpenSpec|README|docs/plans/2026-04-11-core-assessment-design-task.md" docs/task.md
git diff --check
go test -count=1 ./...
```

## Review Outcome

No BDD test code should be added for T15.
