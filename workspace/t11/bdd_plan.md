# T11 BDD Plan

## Decision

Dedicated T11 BDD/Ginkgo coverage is not needed.

T11 is a fixture and golden-test stabilization task. It does not introduce a
new product behavior beyond the contracts already owned by T01-T10. Writing a
separate T11 Ginkgo spec would either assert fixture existence or duplicate the
same product behaviors already covered by task-owned package and command specs.

## Source Artifacts

- Task board: `docs/plans/2026-04-11-core-assessment-design-task.md`
- Workspace strategy: `workspace/t11/test_strategy.md`
- Final implementation record: `workspace/t11/final_impl.md`
- Fixture/golden files under package-local `testdata` directories
- Existing standard tests that consume those fixtures/goldens

## Handling

No new T11-specific `*_ginkgo_test.go` file is planned.

The fixture/golden contracts remain protected through:

- question-bank Ginkgo/TDD coverage for canonical and invalid bank behavior
- answers Ginkgo/TDD coverage for parse and validation behavior
- scoring/classifier Ginkgo/TDD coverage for scoring totals, reverse handling,
  threshold boundaries, and balanced zero
- result renderer Ginkgo/TDD coverage for JSON/text goldens
- CLI command Ginkgo/TDD coverage for questions and score command goldens

## Verification

Use the existing T11 gates after T07-T10 BDD is added:

```bash
go test -count=1 ./internal/questionbank
go test -count=1 ./internal/answers ./internal/scoring ./internal/result ./cmd/mbti-cli
go test -count=1 ./...
```
