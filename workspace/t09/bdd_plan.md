# T09 BDD Plan

## Decision

BDD/Ginkgo coverage is needed.

T09 owns a user-facing CLI command. The behavior is observable through command
stdout, stderr, exit/error state, JSON shape, localized text, deterministic
selection, and failure handling.

## Source Artifacts

- Task board: `docs/plans/2026-04-11-core-assessment-design-task.md`
- Spec: `openspec/specs/questions-command/spec.md`
- Workspace strategy: `workspace/t09/test_strategy.md`
- Production code: `cmd/mbti-cli/questions.go`, `cmd/mbti-cli/root.go`
- Existing tests: `cmd/mbti-cli/questions_test.go`
- CLI goldens: `cmd/mbti-cli/testdata/questions-one-en.txt.golden`

## Planned Ginkgo Coverage

Add a CLI suite if one does not already exist:

- `cmd/mbti-cli/cli_suite_test.go`

Add T09 specs:

- `cmd/mbti-cli/questions_command_ginkgo_test.go`

### Behavior: Canonical bank exports as JSON

`Describe("Questions command")`

`When("the canonical bank is exported as JSON")`

`It("writes parseable prompt data without scoring internals")`

Assertions:

- Execute `questions --questions <canonical> --format json`.
- Expect no stderr and no error.
- Parse stdout as JSON.
- Assert meta title, version, language, count, total, source, and first
  exported question.
- Assert stdout omits scores, reverse flags, thresholds, and dimension metadata.

### Behavior: Localized text export honors count and language

`When("a bounded English text export is requested")`

`It("writes one localized question with option labels")`

Assertions:

- Execute `questions --format text --count 1 --lang en`.
- Compare against the committed text golden.

### Behavior: Seeded bounded selection is deterministic

`When("the same seed and count are used twice")`

`It("selects the same question IDs in the same order")`

Assertions:

- Run the command twice with `--count 3 --seed 7`.
- Compare exported question IDs.

### Behavior: Invalid inputs do not write partial stdout

`When("questions command inputs are invalid")`

`It("returns an error and leaves stdout empty")`

Assertions:

- Cover unsupported format, unsupported language, invalid count, over-large
  count, missing path, missing bank file, and invalid question bank content.
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
