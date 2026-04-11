# T09 Test Strategy

Status: accepted.

## What Must Be Proven

1. `mbti-cli questions --questions <file> --format json` loads and validates the bank before rendering.
2. JSON stdout is parseable, indented, newline-terminated, and contains stable metadata plus exported questions.
3. JSON output does not expose scoring internals: option scores, reverse flags, thresholds, or dimension metadata.
4. Text output is readable and localizes scenario/option labels according to `--lang`.
5. `--count` limits the selected set and rejects invalid counts.
6. `--seed` makes reduced selections deterministic.
7. Invalid `--format`, invalid `--lang`, missing path, or invalid bank input returns an error and does not write partial stdout.
8. T09 does not change answer parsing, scoring, classification, or result rendering behavior.

## Focused Tests

- `go test -count=1 ./cmd/mbti-cli`
- Command JSON test decodes stdout into a struct and checks:
  - meta title/version/language/count/total/source,
  - first question ID and localized text,
  - no `score`, `reverse`, `threshold`, or `dimensions` substrings.
- Command text test uses `--format text --count 1 --lang en`.
- Error table covers invalid format, invalid lang, negative count, count larger than total, and missing file.
- Deterministic selection test runs the same `--count 3 --seed 7` command twice and compares IDs.
- Seed default-all test covers `--seed` with omitted `--count`.
- Invalid invocation tests assert no stdout and stderr containing the same diagnostic as the returned error.

## Repo Gates

- `make fmt`
- `make test`
- `make lint`
- `make build`
- `openspec validate core-assessment`

## CLI Smoke

Run:

```bash
go run . questions --questions questions/mbti-questions-v3.json --format json
```

Then confirm stdout parses as JSON and `len(questions) == 70`.
