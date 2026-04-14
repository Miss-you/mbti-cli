# T03-T18 BDD Scope Check

## Finding

`docs/plans/2026-04-11-core-assessment-design-task.md` currently defines T01
through T15 only. There are no T16, T17, or T18 rows in the task board, and no
`workspace/t16`, `workspace/t17`, or `workspace/t18` task directories.

The user request says "15 tasks" and "T03-T18"; those two scopes do not match.
T03-T18 would be 16 task IDs, while the current task board has 15 tasks total.

## Handling

- T03-T15 were evaluated against `writing-ginkgo-bdd-tests`.
- T16-T18 were not created or inferred because there is no source task/spec
  artifact for them in the referenced task board.
- If T16-T18 are added later, each should get its own workspace directory and
  BDD decision before any Ginkgo code is written.
