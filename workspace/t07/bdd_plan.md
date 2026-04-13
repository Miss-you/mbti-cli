# T07 BDD Plan

## Decision

BDD/Ginkgo coverage is needed.

T07 owns user-visible scoring semantics: threshold buckets, pole direction,
balanced zero handling, deterministic type order, and classifier input errors.
Those are domain behaviors in `openspec/specs/scoring-classifier/spec.md`, not
implementation details.

## Source Artifacts

- Task board: `docs/plans/2026-04-11-core-assessment-design-task.md`
- Spec: `openspec/specs/scoring-classifier/spec.md`
- Workspace strategy: `workspace/t07/test_strategy.md`
- Production code: `internal/scoring/classifier.go`
- Existing tests: `internal/scoring/classifier_test.go`
- Existing Ginkgo suite: `internal/scoring/scoring_suite_test.go`

## Planned Ginkgo Coverage

Add T07 specs:

- `internal/scoring/classifier_ginkgo_test.go`

### Behavior: Threshold scores classify into inclusive buckets

`Describe("Scoring classifier")`

`When("dimension scores fall on configured threshold boundaries")`

`It("labels the matching strength and preserves the raw score")`

Assertions:

- Use a small in-memory bank with the configured threshold buckets.
- Cover representative inclusive boundaries for strong/moderate/slight A and B.
- Assert score, strength, pole, letter, and non-balanced state.

### Behavior: Type is deterministic and zero is balanced

`When("raw dimension scores mix positive, negative, and zero values")`

`It("generates type letters in EI/SN/TF/JP order and marks zero as balanced")`

Assertions:

- Use intentionally unordered score map input.
- Expect type `EXTP`.
- Expect SN zero to render letter `X`, pole `balanced`, and `Balanced=true`.
- Expect positive scores use pole A and negative scores use pole B.

### Behavior: Invalid classifier inputs return contextual errors

`When("required classifier metadata is unavailable")`

`It("returns errors instead of inventing classifications")`

Assertions:

- Missing threshold coverage returns an error containing dimension and score.
- Missing dimension metadata returns an error containing that dimension.

## Negative Control

Because production behavior already exists, verify at least one representative
classifier assertion can fail before restoring the final expectation.

Focused command:

```bash
go test -v ./internal/scoring -run TestScoring -count=1 -ginkgo.v
```

## Verification

```bash
go test -v ./internal/scoring -run TestScoring -count=1 -ginkgo.v
go test -count=1 ./internal/scoring
go test -count=1 ./...
```
