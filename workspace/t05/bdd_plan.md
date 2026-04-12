# T05 BDD Plan

## Decision

BDD/Ginkgo coverage is needed.

T05 owns bank-aware strict answer validation. The OpenSpec scenarios are
behavioral and visible to callers: complete answers pass, unknown IDs fail,
invalid option codes fail, missing answers fail, multiple issues are aggregated
in deterministic order, and validation must not normalize option codes itself.

## Source Artifacts

- Task board: `docs/plans/2026-04-11-core-assessment-design-task.md`
- Spec: `openspec/specs/answer-validation/spec.md`
- Workspace strategy: `workspace/t05/test_strategy.md`
- Production code: `internal/answers/validator.go`
- Existing tests: `internal/answers/validator_test.go`
- Existing suite from T04: `internal/answers/parser_suite_test.go`

## Planned Ginkgo Coverage

Add T05 specs:

- `internal/answers/validator_ginkgo_test.go`

### Behavior: Complete strict answers validate

`Describe("Answer validation")`

`When("every bank question has a defined selected option")`

`It("accepts the answer set without scoring the bank")`

Assertions:

- Build a small bank with three questions and options `A` through `D`.
- `Validate` succeeds for one answer per question.
- Do not call scoring, rendering, CLI, or question-bank schema validation.

### Behavior: Strict validation rejects caller-visible answer issues

Use a small `DescribeTable` or separate `When` blocks for:

- unknown question ID
- invalid option code
- missing answer
- lowercase non-canonical option code rejected without normalization

Each case asserts:

- `Validate` returns `*ValidationError`
- issue text contains the spec-relevant reason

### Behavior: Multiple issues are aggregated deterministically

`When("unknown IDs, invalid option codes, and missing answers appear together")`

`It("returns all validation issues in stable order")`

Assertions:

- Use an answer set whose map iteration order would otherwise be unstable.
- Validate several times.
- Assert the exact `ValidationError.Issues` order.

## Out Of Scope

- Answer file parsing or normalization.
- Question bank schema validation.
- Scoring, reverse score application, rendering, or CLI behavior.
- Exhaustively retesting fixtures already covered by unit tests.

## Negative Control

Because validation behavior already exists, verify one behavior assertion can fail
before restoring it. Prefer temporarily changing the expected deterministic issue
order or an invalid-option issue substring, then run:

```bash
go test -v ./internal/answers -run TestAnswers -count=1 -ginkgo.v
```

## Verification

Focused:

```bash
go test -v ./internal/answers -run TestAnswers -count=1 -ginkgo.v
go test -count=1 ./internal/answers
```

Full after T05:

```bash
go test -count=1 ./...
```
