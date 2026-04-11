# T01 Final Implementation v1

## Decision

Implement only the question bank data model in `internal/questionbank`, plus a focused unmarshal contract test against the canonical v3 bank.

## Product Writes

- `internal/questionbank/model.go`
- `internal/questionbank/model_test.go`
- `questions/mbti-questions-v3.json`

`questions/mbti-questions-v3.json` is included only because the approved design and T01 completion condition name it as the canonical contract file, while the isolated worktree does not inherit the main workspace's untracked copy.

## Coordination Writes

- `workspace/t01/*.md`
- `openspec/changes/core-assessment/*`
- `docs/plans/2026-04-11-core-assessment-design-task.md`

These files are required by the delivery workflow. They do not expand product behavior or T01 implementation scope.

## Detailed Plan

1. Add only `questions/mbti-questions-v3.json` to the isolated worktree from the main workspace fixture because the approved design and T01 completion condition depend on that canonical file.
2. Write the failing `go test ./internal/questionbank` test first.
3. Add minimal structs and constants needed for JSON unmarshalling.
4. Keep validation out of scope. This means no checks for counts, threshold coverage, option code validity, or missing fields beyond what JSON unmarshalling and the test observe.
5. Run `go test ./internal/questionbank` as the first gate, then the broader gates required by the delivery workflow.

## Not Yet Done At This Stage

This document is a plan for the implementation phase. `internal/questionbank` is not considered delivered until the code exists and `go test ./internal/questionbank` passes fresh in the worktree.

## Acceptance Evidence

The test must prove:

- v3 title/version/total are preserved.
- all 70 questions unmarshal.
- dimension metadata for `EI`, `SN`, `TF`, and `JP` is addressable.
- scoring thresholds unmarshal as two-value ranges.
- question `id`, `dimension`, `reverse`, localized scenario, option code/label/score all survive unmarshal.

## OpenSpec

Use change `core-assessment`, aligned with the task board. The change scope for this task is the first model slice only; later tasks will extend behavior.
