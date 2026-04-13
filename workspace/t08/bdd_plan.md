# T08 BDD Plan

## Decision

BDD/Ginkgo coverage is needed.

T08 defines stable user and machine result output. Its OpenSpec scenarios are
observable renderer contracts: summary DTO fields, stable JSON shape, explicit
balanced output, factual text output, and missing-classification errors.

## Source Artifacts

- Task board: `docs/plans/2026-04-11-core-assessment-design-task.md`
- Spec: `openspec/specs/result-renderers/spec.md`
- Workspace strategy: `workspace/t08/test_strategy.md`
- Production code: `internal/result/result.go`
- Existing tests: `internal/result/result_test.go`
- Golden files: `internal/result/testdata/summary.json.golden`,
  `internal/result/testdata/summary.txt.golden`

## Planned Ginkgo Coverage

Add a result suite:

- `internal/result/result_suite_test.go`

Add T08 specs:

- `internal/result/result_ginkgo_test.go`

### Behavior: Summary preserves score and classifier metadata

`Describe("Result summary")`

`When("a summary is built from bank metadata, scores, and classifications")`

`It("preserves meta counts, type, and dimensions in fixed order")`

Assertions:

- Expect title, version, answered count, total count, and type.
- Expect EI/SN/TF/JP dimension values are present in the DTO shape.
- Expect balanced SN is normalized to `letter: X`, `strength: balanced`,
  `pole: balanced`, and `balanced: true`.

### Behavior: Missing classifications are rejected

`When("a required dimension classification is missing")`

`It("returns a contextual missing-dimension error")`

Assertions:

- Remove one required dimension classification.
- Expect error containing the missing dimension.

### Behavior: JSON rendering is stable

`When("a summary is rendered as JSON")`

`It("matches the stable golden output and terminates with a newline")`

Assertions:

- Compare against `summary.json.golden`.
- Assert trailing newline.

### Behavior: Text rendering is factual

`When("a summary is rendered as text")`

`It("matches the readable golden output without diagnostic claims")`

Assertions:

- Compare against `summary.txt.golden`.
- Assert no `diagnosis`, `personality`, or `you are` wording.

## Negative Control

Because production behavior already exists, verify at least one renderer
assertion can fail before restoring the final expectation.

Focused command:

```bash
go test -v ./internal/result -run TestResult -count=1 -ginkgo.v
```

## Verification

```bash
go test -v ./internal/result -run TestResult -count=1 -ginkgo.v
go test -count=1 ./internal/result
go test -count=1 ./...
```
