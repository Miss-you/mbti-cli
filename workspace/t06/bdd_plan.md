# T06 BDD Plan

## Decision

BDD/Ginkgo coverage is needed.

T06 owns raw scoring behavior: complete strict answers produce deterministic
dimension totals, stored signed option scores are authoritative, `reverse` is
not applied a second time, and invalid strict answers stop scoring before
partial totals are returned.

## Source Artifacts

- Task board: `docs/plans/2026-04-11-core-assessment-design-task.md`
- Spec: `openspec/specs/scoring-engine/spec.md`
- Workspace strategy: `workspace/t06/test_strategy.md`
- Production code: `internal/scoring/scoring.go`
- Existing tests: `internal/scoring/scoring_test.go`

## Planned Ginkgo Coverage

Add a scoring suite if one does not already exist:

- `internal/scoring/scoring_suite_test.go`

Add T06 specs:

- `internal/scoring/scoring_ginkgo_test.go`

### Behavior: Complete answers produce raw dimension totals

`Describe("Scoring engine")`

`When("a complete strict answer set is scored")`

`It("adds selected option scores by dimension and reports answer counts")`

Assertions:

- Build a small four-question bank across multiple dimensions.
- Call `Score` with one answer per question.
- Assert `Answered`, `Total`, and `DimensionScores`.

### Behavior: Reversed question uses stored signed score

`When("a selected option belongs to a reversed question")`

`It("adds the stored signed score without applying reverse again")`

Assertions:

- Include a `reverse=true` question whose selected option has score `-2`.
- Assert the dimension total includes `-2`, not `+2`.

### Behavior: Invalid strict answers stop scoring

`When("answers are incomplete")`

`It("returns a validation error without partial dimension totals")`

Assertions:

- Omit one bank question.
- `Score` returns an answer validation error containing the missing question ID.
- Returned result is zero.

### Behavior: Result stays raw

The T06 BDD should not add reflection or artificial field-shape assertions.
Instead, keep assertions on the runtime result fields T06 owns: answer counts and
raw `DimensionScores`. Do not call threshold classification, type generation,
renderers, or CLI commands.

## Out Of Scope

- Answer parsing.
- Question bank schema validation.
- Threshold classification, MBTI type generation, JSON/text rendering, or CLI
  behavior.
- Reflection-based assertions about fields that do not exist today.

## Negative Control

Because scoring behavior already exists, verify one behavior assertion can fail
before restoring it. Prefer temporarily changing the expected reversed score from
`-2` to `+2`, then run:

```bash
go test -v ./internal/scoring -run TestScoring -count=1 -ginkgo.v
```

## Verification

Focused:

```bash
go test -v ./internal/scoring -run TestScoring -count=1 -ginkgo.v
go test -count=1 ./internal/scoring
```

Full after T06:

```bash
go test -count=1 ./...
```
