# T04 BDD Plan

## Decision

No additional BDD test code is needed for T04.

`internal/answers/parser_ginkgo_test.go` already covers the OpenSpec
parser-level contract: canonical map parsing, normalization, missing/null
`answers`, malformed JSON, structurally invalid values, and deferring
bank-aware validation. The proposed fixture-backed scenario would duplicate
those same spec-level behaviors instead of adding new business coverage.

## Source Artifacts

- Task board: `docs/plans/2026-04-11-core-assessment-design-task.md`
- Spec: `openspec/specs/answer-parser/spec.md`
- Workspace strategy: `workspace/t04/test_strategy.md`
- Production code: `internal/answers/parser.go`
- Existing tests: `internal/answers/parser_test.go`
- Existing Ginkgo suite/spec: `internal/answers/parser_suite_test.go`,
  `internal/answers/parser_ginkgo_test.go`

## Planned Ginkgo Coverage

No new specs are planned.

Keep the existing coverage in `internal/answers/parser_ginkgo_test.go` as the
BDD evidence for the answer parser contract.

## Out Of Scope

- Unknown question ID validation.
- Invalid option-code validation against a question bank.
- Missing required answer validation.
- Scoring, result rendering, or CLI behavior.
- Rewriting existing parser BDD into a different style.

## Negative Control

No new BDD scenario is being added, so no additional negative control is
required for T04.

## Verification

Focused:

```bash
go test -v ./internal/answers -run TestAnswers -count=1 -ginkgo.v
go test -count=1 ./internal/answers
```

Full after T04:

```bash
go test -count=1 ./...
```
