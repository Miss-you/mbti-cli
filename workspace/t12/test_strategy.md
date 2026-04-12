# T12 Test Strategy

Goal: prove that the first-stage core assessment loop remains aligned across code, specs, CLI behavior, and repository quality gates.

Required gates:

- `make fmt`
- `make test`
- `make lint`
- `make build`
- `openspec validate core-assessment`

Focused CLI smoke checks:

- Build the binary and run `bin/mbti-cli questions --questions questions/mbti-questions-v3.json --count 3 --seed 123 --lang zh --format json`.
  - Assert stdout is parseable JSON.
  - Assert stderr is empty.
- Run `bin/mbti-cli score --questions questions/mbti-questions-v3.json --answers cmd/mbti-cli/testdata/answers-all-a.json --format json`.
  - Assert stdout is parseable JSON.
  - Assert stderr is empty.
  - Assert the result includes the expected `type` field.
- Run malformed question-bank smoke checks for both `questions` and `score`.
  - Assert each command exits non-zero.
  - Assert stdout is empty.
  - Assert stderr contains `parse question bank`.

Review checklist:

- Question bank validator still protects meta totals, dimension counts, option codes/scores, and threshold completeness/non-overlap.
- Answer validation still rejects unknown IDs, invalid options, and missing strict answers.
- Scoring still uses signed option scores once and does not reapply `reverse`.
- Classifier still maps zero scores to `balanced` and type letter `X`.
- Result renderers keep stable JSON/text output and avoid diagnostic claims.
- CLI commands keep success output on stdout and errors on stderr with no partial stdout.

No extra unit tests are planned unless review or a gate reveals an unprotected defect.
