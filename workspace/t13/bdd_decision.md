# T13 BDD Decision

## Decision

Dedicated BDD/Ginkgo coverage is not needed for T13.

T13 is a README documentation alignment task. It does not add or change Go
package behavior, CLI flags, stdout/stderr contracts, scoring semantics, or
OpenSpec requirements. The underlying product behavior is already covered by
the T09 and T10 command BDD suites.

## Source Artifacts

- Task board: `docs/plans/2026-04-11-core-assessment-design-task.md`
- Workspace strategy: `workspace/t13/test_strategy.md`
- Documentation target: `README.md`
- Existing CLI BDD: `cmd/mbti-cli/questions_command_ginkgo_test.go`
- Existing CLI BDD: `cmd/mbti-cli/score_command_ginkgo_test.go`

## Handling

T13 should be verified with documentation checks and CLI smokes, not with a new
Ginkgo spec that reasserts already-covered command behavior.

Required evidence remains:

```bash
rg -n "mbti-cli (questions|score)|answers|mbti-questions-v3" README.md
rg -n '"answers":|q01|option code|A/B/C/D' README.md
rg -n "AI behavior style|not a human personality diagnosis" README.md
go run . questions --questions questions/mbti-questions-v3.json --count 1 --seed 123 --lang en --format json
go run . score --questions questions/mbti-questions-v3.json --answers cmd/mbti-cli/testdata/answers-all-a.json --format json
go test -count=1 ./...
```

## Review Outcome

No BDD test code should be added for T13.
