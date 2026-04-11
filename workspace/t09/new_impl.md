# T09 Candidate Implementation

Status: proposed.

## Shape

Add `cmd/mbti-cli/questions.go` and focused command tests.

Command:

```bash
mbti-cli questions --questions questions/mbti-questions-v3.json --format json
```

Flags:

- `--questions <path>`: required bank file path.
- `--format text|json`: defaults to `text`.
- `--lang zh|en`: defaults to `zh`.
- `--count <n>`: defaults to all questions; rejects negative and values larger than the bank.
- `--seed <n>`: optional deterministic shuffle before count selection.

JSON output:

```json
{
  "meta": {
    "title": "AI Behavioral Style Assessment v3",
    "version": "0.3.0",
    "language": "zh",
    "count": 70,
    "total": 70,
    "source": "questions/mbti-questions-v3.json"
  },
  "questions": [
    {
      "id": "q01",
      "scenario": "...",
      "options": [
        {"code": "A", "label": "..."}
      ]
    }
  ]
}
```

The export intentionally omits `score`, `reverse`, thresholds, and dimension metadata. Those fields are for validation and scoring, not for the answer collection surface.

## Tests

- Red test for `questions --questions <file> --format json` producing parseable JSON with stable metadata and no scoring fields.
- Red test for `--format text --count 1 --lang en` producing readable one-question output.
- Red test for invalid format/lang/count and invalid bank path returning errors and no stdout.

## Trade-offs

- No new public internal package is introduced for T09. Private command-level DTO helpers keep scope small.
- Seeded selection uses Go's deterministic `math/rand` source. It is sufficient for reproducible CLI tests, not intended as a security boundary.
- `--count 0` means all questions, which keeps the required command short and compatible with the task's acceptance example.
