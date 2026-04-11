# T10 Final Implementation

Status: accepted for implementation.

## Accepted Plan

Add `mbti-cli score` as a thin Cobra orchestration layer that reads a question bank and answer file, then delegates to existing core packages for validation, scoring, classification, summary construction, and rendering.

## Command Contract

```bash
mbti-cli score \
  --questions questions/mbti-questions-v3.json \
  --answers answers.json \
  --format json
```

Flags:

- `--questions`: required question bank JSON path.
- `--answers`: required canonical answer JSON path.
- `--format`: `text` or `json`, default `text`.

Successful runs write only to stdout. Errors return from Cobra and leave command stdout empty.

## Core Flow

1. Load the bank with `questionbank.LoadFile`.
2. Validate the bank with `questionbank.Validate`.
3. Read the answer file.
4. Parse answers with `answers.Parse`.
5. Score with `scoring.Score`.
6. Classify with `scoring.Classify`.
7. Build a result summary with `result.NewSummary`.
8. Render with `result.RenderJSON` or `result.RenderText`.

## Scope Boundaries

- No interactive `assess` command.
- No partial scoring.
- No new result schema.
- No duplicate scoring or validation logic in Cobra code.
- No shared command abstraction unless later tasks introduce real duplication pressure.

## Review

Local review pass found no must-fix correctness issues in the main flow.

One review improvement was applied:

- Wrapped malformed answer parsing errors with the answer file path.
- Added command-level strict validation tests for missing answers, unknown question IDs, and invalid option codes.

## Final Verification

Fresh verification after the review fix:

```bash
make fmt
go test -count=1 ./cmd/mbti-cli
make test
make lint
make build
openspec validate core-assessment
go run . score --questions questions/mbti-questions-v3.json --answers <tmp-all-A-answers> --format json | jq -e '.meta.answered == 70 and .meta.total == 70 and .type == "ESTJ" and .dimensions.EI.score == 36 and .dimensions.SN.score == 36 and .dimensions.TF.score == 34 and .dimensions.JP.score == 34'
```

OpenSpec change `core-assessment` remains active for T11/T12.
