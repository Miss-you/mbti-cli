# T13 Test Strategy

## What To Prove

This task must prove that the README matches the current CLI and that the documented commands still work.

## Checks

1. Documentation grep:

   ```bash
   rg -n "mbti-cli (questions|score)|answers|mbti-questions-v3" README.md
   rg -n '"answers":|q01|option code|A/B/C/D' README.md
   rg -n "AI behavior style|not a human personality diagnosis" README.md
   ```

2. Focused CLI smokes:

   ```bash
   go run . questions --questions questions/mbti-questions-v3.json --count 1 --seed 123 --lang en --format json
   go run . score --questions questions/mbti-questions-v3.json --answers cmd/mbti-cli/testdata/answers-all-a.json --format json
   ```

3. Repo gates:

   ```bash
   make fmt
   make test
   make lint
   make build
   ```

## OpenSpec

No OpenSpec change is required because this task only updates README documentation for behavior already specified and implemented by the active `core-assessment` change.
