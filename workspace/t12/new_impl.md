# T12 Proposed Implementation

Use T12 as the final verification and review repair pass for the first-stage core assessment work.

Plan:

1. Keep implementation changes minimal and evidence-driven.
2. Run the required repo gates: `make fmt`, `make test`, `make lint`, and `make build`.
3. Run `openspec validate core-assessment`.
4. Run focused CLI smoke checks for both first-stage commands:
   - `questions --questions questions/mbti-questions-v3.json --count 3 --seed 123 --lang zh --format json`
   - `score --questions questions/mbti-questions-v3.json --answers cmd/mbti-cli/testdata/answers-all-a.json --format json`
5. Review implementation against the design and OpenSpec specs, prioritizing correctness, stdout/stderr separation, reverse scoring, threshold boundaries, balanced zero handling, and stable output contracts.
6. Fix only must-fix findings, then rerun the affected focused checks plus the full required gates.

OpenSpec:

- No new change is needed because T12 does not introduce new behavior.
- The existing `core-assessment` change remains the linked change and is validated as part of closure.
