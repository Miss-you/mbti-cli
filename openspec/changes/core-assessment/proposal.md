## Why

The first core-assessment slice needs a typed representation of the canonical v3 question bank before loader, validator, scoring, renderer, and CLI tasks can build on it.

## What Changes

- Add a Go question-bank model that matches the v3 JSON shape.
- Preserve metadata, dimension definitions, thresholds, questions, localized text, option scores, and `reverse`.
- Keep loading, validation, scoring, rendering, and CLI integration out of this first slice.

## Capabilities

### New Capabilities
- `question-bank-model`: Represents the canonical MBTI question bank JSON structure in Go.

### Modified Capabilities

## Impact

- Adds `internal/questionbank`.
- Adds the canonical v3 question bank fixture needed by model tests.
- Establishes types that later answer validation, scoring, rendering, and CLI tasks can consume.
