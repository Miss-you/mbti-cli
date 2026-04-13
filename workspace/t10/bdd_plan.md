# T10 BDD Plan

## Decision

BDD/Ginkgo coverage is needed.

T10 owns a user-facing scoring CLI command. The behavior is observable through
command stdout, stderr, exit/error state, rendered JSON/text summaries, strict
answer validation, malformed input handling, and delegation to the existing core
packages.

## Source Artifacts

- Task board: `docs/plans/2026-04-11-core-assessment-design-task.md`
- Spec: `openspec/specs/score-command/spec.md`
- Workspace strategy: `workspace/t10/test_strategy.md`
- Production code: `cmd/mbti-cli/score.go`, `cmd/mbti-cli/root.go`
- Existing tests: `cmd/mbti-cli/score_test.go`
- CLI fixtures/goldens: `cmd/mbti-cli/testdata/answers-all-a.json`,
  `cmd/mbti-cli/testdata/answers-all-d.json`,
  `cmd/mbti-cli/testdata/score-all-a.json.golden`,
  `cmd/mbti-cli/testdata/score-all-d.txt.golden`

## Planned Ginkgo Coverage

Reuse the CLI suite planned for T09:

- `cmd/mbti-cli/cli_suite_test.go`

Add T10 specs:

- `cmd/mbti-cli/score_command_ginkgo_test.go`

### Behavior: Canonical answers render as JSON

`Describe("Score command")`

`When("complete canonical answers are scored as JSON")`

`It("writes the stable result summary with all four dimensions")`

Assertions:

- Execute `score --questions <canonical> --answers answers-all-a.json --format json`.
- Expect no stderr and no error.
- Compare stdout against the committed JSON golden.
- Parse stdout and assert answered count, total count, type `ESTJ`, and
  dimension scores.

### Behavior: Canonical answers render as factual text

`When("complete canonical answers are scored as text")`

`It("writes the factual renderer output without diagnostic claims")`

Assertions:

- Execute `score --questions <canonical> --answers answers-all-d.json --format text`.
- Compare stdout against the committed text golden.
- Assert no `diagnosis`, `personality`, or `you are` wording.

### Behavior: Invalid inputs do not write partial stdout

`When("score command inputs are invalid")`

`It("returns an error and leaves stdout empty")`

Assertions:

- Cover unsupported format, missing question bank path, missing answer path,
  missing answer file, invalid question bank content, malformed answer JSON,
  missing answer, unknown question ID, and invalid option code.
- Expect returned error, stderr diagnostic, and empty stdout.

## Negative Control

Because production behavior already exists, verify at least one command-output
assertion can fail before restoring the final expectation.

Focused command:

```bash
go test -v ./cmd/mbti-cli -run TestCLI -count=1 -ginkgo.v
```

## Verification

```bash
go test -v ./cmd/mbti-cli -run TestCLI -count=1 -ginkgo.v
go test -count=1 ./cmd/mbti-cli
go test -count=1 ./...
```
