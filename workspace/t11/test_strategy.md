# T11 Test Strategy

## Purpose

Prove that committed fixtures and golden files cover the first-phase core assessment contracts without changing product behavior.

## What Must Be Proven

- Schema fixtures: a valid bank loads and validates; invalid bank files fail for meta total, duplicate question ID, and threshold overlap.
- Answer fixtures: a valid canonical answer file parses and validates; an invalid answer file reports unknown ID, invalid option, and missing answers.
- Scoring fixtures: deterministic dimension totals come from stored signed option scores; a reversed question is not inverted a second time.
- Classifier fixtures: threshold boundary cases still map to the expected strength buckets; a zero score maps to `balanced` and type letter `X`.
- Result goldens: JSON and text renderers preserve exact stable output, including indentation and trailing newline.
- CLI goldens: representative `questions` and `score` success paths write exact expected stdout and no stderr.

## Focused Gates

1. `go test -count=1 ./internal/questionbank`
2. `go test -count=1 ./internal/answers ./internal/scoring ./internal/result ./cmd/mbti-cli`
3. `go test -count=1 ./...`

## Repo Gates

- `make fmt`
- `make test`
- `make lint`
- `make build`
- `openspec validate core-assessment`

## Smoke Checks

- `go run . questions --questions questions/mbti-questions-v3.json --format json --count 1 --lang en`
- `go run . score --questions questions/mbti-questions-v3.json --answers cmd/mbti-cli/testdata/answers-all-a.json --format json`

## Out of Scope

- No new product behavior.
- No new CLI flags.
- No T12 full verification repair beyond proving T11 changes.
