## Why

The first core-assessment slices need a typed representation of the canonical v3 question bank and a focused filesystem loader before validator, scoring, renderer, and CLI tasks can build on them.

## What Changes

- Add a Go question-bank model that matches the v3 JSON shape.
- Preserve metadata, dimension definitions, thresholds, questions, localized text, option scores, and `reverse`.
- Add a file-path loader that returns the typed bank plus source metadata.
- Keep schema validation, scoring, rendering, and CLI integration out of these early slices.

## Capabilities

### New Capabilities
- `question-bank-model`: Represents the canonical MBTI question bank JSON structure in Go.
- `question-bank-loader`: Loads a question bank JSON file into the typed model with source metadata.

### Modified Capabilities

## Impact

- Adds `internal/questionbank`.
- Adds the canonical v3 question bank fixture needed by model tests.
- Establishes types and loading behavior that later validation, answer, scoring, rendering, and CLI tasks can consume.
