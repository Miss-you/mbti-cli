## Context

`mbti-cli` currently has Cobra root/version commands and no complete core assessment path. The approved first-phase design names `questions/mbti-questions-v3.json` as the canonical bank and defines loader, validator, answer, scoring, renderer, and CLI tasks that depend on a shared model.

## Goals / Non-Goals

**Goals:**

- Introduce Go data types that match the v3 bank JSON structure.
- Preserve fields required by later validation and scoring tasks.
- Prove the model with a focused unmarshal contract test.
- Add a focused file-path loader that returns typed bank data and source metadata.
- Prove missing and malformed files produce clear errors.
- Add a focused canonical answer-file parser that returns normalized answer codes.
- Add a focused question-bank schema validator for metadata, dimensions, question shape, option shape, and thresholds.
- Add strict answer validation against a question bank for unknown IDs, invalid option codes, and missing required answers.

**Non-Goals:**

- No scoring, threshold classification, rendering, or CLI command wiring.

## Decisions

- Create `internal/questionbank` as the question bank owner because the model, loader, and validator naturally belong there.
- Use JSON-tagged structs and lightweight string aliases for dimensions and strength buckets. This keeps unmarshalling simple while giving later tasks shared constants.
- Represent threshold bucket values as a fixed two-int range so malformed range length can be validated explicitly by a later validator task.
- Preserve `reverse` as raw metadata. T01 does not interpret it or transform scores.
- Add only the v3 canonical question file needed by this task's contract test. Other question bank versions remain out of scope.
- Keep loader behavior filesystem-focused for T02: `LoadFile(path)` reads JSON, unmarshals `Bank`, and returns path, filename, and byte-size source metadata.
- Wrap read and parse errors with path-specific context while preserving the underlying error for callers that need `errors.Is`.
- Create `internal/answers` for the answer-file parser so answer parsing, later answer validation, and scoring can remain separate from question bank loading.
- T04 parser behavior is structural only: it decodes the canonical top-level `answers` map and normalizes option codes with trim + uppercase.
- Preserve question IDs as supplied by the answer file. T05 owns bank-aware unknown-ID, invalid-option, and missing-answer validation.
- Add `questionbank.Validate(Bank) error` as the T03 schema gate. It accumulates deterministic issues for metadata, supported dimensions, dimension counts, question IDs, localized text, option codes, score values, and threshold buckets.
- Keep `LoadFile` validation-free so callers can choose when to report parse errors versus schema errors.
- Treat `reverse` as metadata only. T03 validates presence through the typed model but does not transform option scores or scoring behavior.
- Add `answers.Validate(questionbank.Bank, answers.Set) error` as the T05 strict answer gate. It checks answer IDs and option codes against the bank, requires every bank question to be answered, and keeps scoring behavior out of scope.

## Risks / Trade-offs

- The main workspace currently holds `questions/mbti-questions-v3.json` as an untracked file, so adding it here creates data churn. Mitigation: add only the canonical v3 file required by the approved design and task acceptance.
- Struct types can imply validation guarantees they do not enforce. Mitigation: keep validation out of names and tests; later T03 owns schema checks.
- Loader metadata can grow over time. Mitigation: start with only path, base filename, and size because those are enough for diagnostics and later CLI output.
- A validator can drift into scoring behavior. Mitigation: keep T03 limited to static bank schema checks and defer score aggregation and classification to later tasks.
