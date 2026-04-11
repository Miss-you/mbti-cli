# T03 Original Implementation

## Current state

- `internal/questionbank/model.go` contains representation-only structs for the canonical v3 question bank JSON.
- `internal/questionbank/loader.go` reads a file path, unmarshals `Bank`, and returns source metadata.
- Existing tests prove field preservation and loader error context.
- No code currently validates schema invariants such as meta total, supported dimensions, question ID uniqueness, option shape, score set, or threshold completeness.

## Constraints

- T01 intentionally kept model types validation-free.
- T02 intentionally kept `LoadFile` filesystem-focused and did not validate loaded data.
- T03 should add bank schema validation without scoring answers or interpreting `reverse`.
