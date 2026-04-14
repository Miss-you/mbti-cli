# T14 BDD Decision

## Decision

Dedicated BDD/Ginkgo coverage is not needed for T14.

T14 is an OpenSpec archive cleanup task. It synchronizes completed
`core-assessment` specs into the main spec tree and removes active-change drift.
It does not change production Go behavior or user-observable CLI output.

## Source Artifacts

- Task board: `docs/plans/2026-04-11-core-assessment-design-task.md`
- Workspace strategy: `workspace/t14/test_strategy.md`
- Archived change: `openspec/changes/archive/2026-04-12-core-assessment/`
- Main specs: `openspec/specs/`

## Handling

T14 should be verified through repository state and OpenSpec validation rather
than a Ginkgo spec. A T14-specific Ginkgo test would test process artifacts, not
product behavior.

Required evidence remains:

```bash
test ! -d openspec/changes/core-assessment
test -d openspec/changes/archive/2026-04-12-core-assessment
openspec validate --all
go test -count=1 ./...
```

## Review Outcome

No BDD test code should be added for T14.
