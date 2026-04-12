# T03 BDD Plan

## Decision

BDD/Ginkgo coverage is needed.

T03 is a schema validator task with explicit OpenSpec scenarios for valid and
invalid question-bank contracts. The existing standard tests are comprehensive;
the BDD layer should restate the validator behavior by scenario group rather
than mechanically duplicating every unit-test mutation.

## Source Artifacts

- Task board: `docs/plans/2026-04-11-core-assessment-design-task.md`
- Spec: `openspec/specs/question-bank-validator/spec.md`
- Workspace strategy: `workspace/t03/test_strategy.md`
- Production code: `internal/questionbank/validator.go`
- Existing tests: `internal/questionbank/validator_test.go`
- Existing suite from T01: `internal/questionbank/questionbank_suite_test.go`

## Planned Ginkgo Coverage

Add T03 specs:

- `internal/questionbank/validator_ginkgo_test.go`

### Behavior: Canonical v3 bank validates

`Describe("Question bank validator")`

`When("the canonical v3 bank is validated")`

`It("accepts the schema without applying scoring behavior")`

Assertions:

- Unmarshal `questions/mbti-questions-v3.json` directly into `Bank`.
- `Validate(bank)` succeeds.
- Do not call scoring, classifier, rendering, answer validation, or CLI helpers.

### Behavior: Invalid schema invariants return clear validation errors

Use `DescribeTable` entries grouped by OpenSpec scenario, with one or two
representative mutations per group at most, so the BDD stays at the
spec-behavior level instead of mirroring every `validator_test.go` mutation:

- metadata total mismatch
- dimension metadata missing and dimension count mismatch
- question identity/localized text invalid
- option shape invalid
- threshold bucket incomplete/malformed/overlapping

Each entry:

- starts from the canonical v3 bank
- applies a small mutation
- calls `Validate`
- asserts the error is a `*ValidationError`
- asserts the issue text contains the spec-relevant reason
- keeps the table to representative failures for the scenario group

## Out Of Scope

- Loader read/parse errors.
- Answer parser or answer validation behavior.
- Score aggregation, reverse-score application, threshold classification, result
  rendering, or CLI behavior.
- Exhaustively duplicating every `validator_test.go` mutation.

## Negative Control

Because validator behavior already exists, verify one behavior assertion can fail
before restoring it. Prefer temporarily changing an expected issue substring for
one invalid schema case, then run:

```bash
go test -v ./internal/questionbank -run TestQuestionbank -count=1 -ginkgo.v
```

## Verification

Focused:

```bash
go test -v ./internal/questionbank -run TestQuestionbank -count=1 -ginkgo.v
go test -count=1 ./internal/questionbank
```

Full after T03:

```bash
go test -count=1 ./...
```
