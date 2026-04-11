## Context

`mbti-cli` currently has Cobra root/version commands and no core package for question banks. The approved first-phase design names `questions/mbti-questions-v3.json` as the canonical bank and defines later loader, validator, answer, scoring, renderer, and CLI tasks that depend on a shared model.

## Goals / Non-Goals

**Goals:**

- Introduce Go data types that match the v3 bank JSON structure.
- Preserve fields required by later validation and scoring tasks.
- Prove the model with a focused unmarshal contract test.

**Non-Goals:**

- No file loader.
- No schema validation.
- No answer parsing or validation.
- No scoring, threshold classification, rendering, or CLI command wiring.

## Decisions

- Create `internal/questionbank` as the model owner because later loader and validator tasks naturally belong there.
- Use JSON-tagged structs and lightweight string aliases for dimensions and strength buckets. This keeps unmarshalling simple while giving later tasks shared constants.
- Represent threshold bucket values as a fixed two-int range so malformed range length can be validated explicitly by a later validator task.
- Preserve `reverse` as raw metadata. T01 does not interpret it or transform scores.
- Add only the v3 canonical question file needed by this task's contract test. Other question bank versions remain out of scope.

## Risks / Trade-offs

- The main workspace currently holds `questions/mbti-questions-v3.json` as an untracked file, so adding it here creates data churn. Mitigation: add only the canonical v3 file required by the approved design and task acceptance.
- Struct types can imply validation guarantees they do not enforce. Mitigation: keep validation out of names and tests; later T03 owns schema checks.
