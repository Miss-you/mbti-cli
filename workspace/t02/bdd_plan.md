# T02 BDD Plan

## Decision

BDD/Ginkgo coverage is needed.

T02 owns user-observable loader behavior: loading a question bank from a
filesystem path, returning source metadata, and reporting path/read/parse errors
clearly. The scenarios are explicit in the question bank model/loader spec and
map cleanly to behavior-style Ginkgo tests.

## Source Artifacts

- Task board: `docs/plans/2026-04-11-core-assessment-design-task.md`
- Spec: `openspec/specs/question-bank-model/spec.md`
- Workspace strategy: `workspace/t02/test_strategy.md`
- Production code: `internal/questionbank/loader.go`
- Existing tests: `internal/questionbank/loader_test.go`
- Existing suite from T01: `internal/questionbank/questionbank_suite_test.go`

## Planned Ginkgo Coverage

Add T02 specs:

- `internal/questionbank/loader_ginkgo_test.go`

### Behavior: Canonical v3 bank loads from path

`Describe("Question bank loader")`

`When("the canonical v3 question bank path is loaded")`

`It("returns the typed bank with source metadata")`

Assertions:

- `LoadFile` succeeds for `questions/mbti-questions-v3.json`.
- Returned bank has representative metadata and the first question ID.
- Source metadata includes the input path, base filename, and positive byte size.

### Behavior: Missing bank file reports read context

`When("the path does not exist")`

`It("returns a read error with path context and preserves the missing-file cause")`

Assertions:

- Error contains `read question bank`.
- Error contains the missing path.
- `errors.Is(err, os.ErrNotExist)` is true.

### Behavior: Malformed bank file reports parse context

`When("the file is not valid question bank JSON")`

`It("returns a parse error with source path context")`

Assertions:

- Temp file contains malformed JSON.
- Error contains `parse question bank`.
- Error contains the temp file path.

### Behavior: Empty path is rejected before reading

`When("the path is empty")`

`It("returns a path-required error")`

Assertions:

- `LoadFile("")` fails with `question bank path is required`.

## Out Of Scope

- Schema validation of loaded banks.
- Threshold, dimension, question, and option invariant checks.
- Answer parsing, answer validation, scoring, rendering, or CLI behavior.

## Negative Control

Because loader behavior already exists, verify one behavior assertion can fail
before restoring it. Prefer temporarily changing the expected source filename or
the expected empty-path error text, then run:

```bash
go test -v ./internal/questionbank -run TestQuestionbank -count=1 -ginkgo.v
```

## Verification

Focused:

```bash
go test -v ./internal/questionbank -run TestQuestionbank -count=1 -ginkgo.v
go test -count=1 ./internal/questionbank
```

Full after T02:

```bash
go test -count=1 ./...
```
