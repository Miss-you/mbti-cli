# T11 Proposed Implementation

Status: proposed.

## Plan

Keep `questions/mbti-questions-v3.json` as the production canonical bank and add a small explicit test fixture layer around the already-implemented behavior.

## Fixture Assets

- `internal/questionbank/testdata/valid-bank.json`
- `internal/questionbank/testdata/invalid-meta-total.json`
- `internal/questionbank/testdata/invalid-duplicate-question-id.json`
- `internal/questionbank/testdata/invalid-threshold-overlap.json`
- `internal/answers/testdata/valid-answers.json`
- `internal/answers/testdata/invalid-answers.json`
- `internal/scoring/testdata/bank.json`
- `internal/scoring/testdata/answers.json`
- `internal/result/testdata/summary.json.golden`
- `internal/result/testdata/summary.txt.golden`
- `cmd/mbti-cli/testdata/answers-all-a.json`
- `cmd/mbti-cli/testdata/answers-all-d.json`
- `cmd/mbti-cli/testdata/questions-one-en.txt.golden`
- `cmd/mbti-cli/testdata/score-all-a.json.golden`
- `cmd/mbti-cli/testdata/score-all-d.txt.golden`

## Test Updates

- `internal/questionbank/validator_test.go`: add fixture-backed valid and invalid schema tests while keeping representative in-memory mutation coverage.
- `internal/answers/parser_test.go` and `validator_test.go`: add file-backed answer parse and strict validation tests.
- `internal/scoring/scoring_test.go`: add a fixture-backed deterministic scoring test that includes `reverse=true` and proves signed scores are used once.
- `internal/result/result_test.go`: read JSON and text golden files instead of embedding long expected strings inline.
- `cmd/mbti-cli/questions_test.go`: compare the one-question English text output to a golden file.
- `cmd/mbti-cli/score_test.go`: use committed answer fixtures and compare JSON/text success output to golden files.

## Boundaries

- No production code changes.
- No new runtime dependencies.
- No OpenSpec behavior delta; only add the T11 execution items to the active `core-assessment` change tasks.

## First Gate

Run `go test -count=1 ./internal/questionbank` first because bad fixture shape and path errors should fail there earliest.
