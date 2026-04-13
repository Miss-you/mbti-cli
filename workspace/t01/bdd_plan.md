# T01 BDD Plan

## Decision

BDD/Ginkgo coverage is needed.

T01 has OpenSpec behavior that is more than an implementation detail: the typed
question bank model must preserve the canonical JSON contract without doing
validation, loading, scoring, rendering, or CLI work. Existing standard tests
already cover this, but a Ginkgo spec gives the task a readable behavior
specification.

## Source Artifacts

- Task board: `docs/plans/2026-04-11-core-assessment-design-task.md`
- Spec: `openspec/specs/question-bank-model/spec.md`
- Workspace strategy: `workspace/t01/test_strategy.md`
- Production code: `internal/questionbank/model.go`
- Existing tests: `internal/questionbank/model_test.go`

## Planned Ginkgo Coverage

Add a questionbank suite if one does not already exist:

- `internal/questionbank/questionbank_suite_test.go`

Add T01 specs:

- `internal/questionbank/model_ginkgo_test.go`

### Behavior: Canonical v3 bank unmarshals into the model

`Describe("Question bank model")`

`When("the canonical v3 question bank is unmarshaled")`

`It("preserves representative metadata, dimensions, thresholds, questions, localized text, options, signed scores, and reverse flags")`

Assertions:

- Load `questions/mbti-questions-v3.json`.
- Unmarshal directly into `questionbank.Bank`.
- Check representative fields:
  - title and version
  - total count
  - dimensions and threshold ranges
  - first question and a reversed question
  - a representative option label and signed score
- Do not add schema-invalid fixtures, loader behavior, validation, scoring,
  rendering, or CLI helpers to the T01 BDD plan.

## Negative Control

Because production behavior already exists, verify the canonical field-preservation
assertion can fail before restoring the final expectation. The focused command
should fail on the behavior assertion, not on suite wiring:

```bash
go test -v ./internal/questionbank -run TestQuestionbank -count=1 -ginkgo.v
```

## Verification

Focused:

```bash
go test -v ./internal/questionbank -run TestQuestionbank -count=1 -ginkgo.v
go test -count=1 ./internal/questionbank
```

Full after T01:

```bash
go test -count=1 ./...
```
