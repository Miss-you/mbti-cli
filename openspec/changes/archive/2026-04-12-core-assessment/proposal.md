## Why

The first core-assessment slices need a typed representation of the canonical v3 question bank and a focused filesystem loader before validator, scoring, renderer, and CLI tasks can build on them.

## What Changes

- Add a Go question-bank model that matches the v3 JSON shape.
- Preserve metadata, dimension definitions, thresholds, questions, localized text, option scores, and `reverse`.
- Add a file-path loader that returns the typed bank plus source metadata.
- Add a focused canonical answer-file parser that returns normalized `questionID -> optionCode` answers.
- Add question-bank schema validation for metadata, dimensions, questions, options, and thresholds.
- Add strict answer validation against a bank for unknown IDs, invalid options, and missing answers.
- Keep scoring, rendering, and CLI integration out of these early slices.

## Capabilities

### New Capabilities
- `question-bank-model`: Represents the canonical MBTI question bank JSON structure in Go.
- `question-bank-loader`: Loads a question bank JSON file into the typed model with source metadata.
- `question-bank-validator`: Validates canonical question bank schema invariants after loading/unmarshalling.
- `answer-parser`: Parses the first-phase canonical answer JSON map into normalized answer codes.
- `answer-validation`: Validates parsed answer sets against a question bank in strict mode.

### Modified Capabilities

## Impact

- Adds `internal/questionbank`.
- Adds `internal/answers`.
- Adds the canonical v3 question bank fixture needed by model tests.
- Establishes types, loading, schema validation, answer parsing, and strict answer validation behavior that later scoring, rendering, and CLI tasks can consume.
