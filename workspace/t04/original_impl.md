# T04 Original Implementation

## Current State

- `internal/questionbank` owns the canonical question bank model and file loader.
- There is no `internal/answers` package yet.
- The approved design defines the first-phase answer file as:

```json
{
  "answers": {
    "q01": "A",
    "q02": "C"
  }
}
```

## Existing Boundaries

- T04 only needs to parse the canonical answer file into `questionID -> optionCode`.
- T05 owns validation against a question bank: unknown IDs, invalid option codes, and missing required answers.
- CLI file path handling stays out of T04.

## Constraints

- Preserve `questions` and scoring semantics by not importing or depending on `internal/questionbank`.
- Normalize answer option codes to uppercase for later validation/scoring.
- Keep parser errors clear enough for future CLI callers.
