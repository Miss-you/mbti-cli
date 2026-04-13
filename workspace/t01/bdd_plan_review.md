# T01 BDD Plan Review

## Verdict

Partially aligned. The canonical v3 unmarshal scenario matches the T01 spec and the writing-ginkgo-bdd-tests guidance, but the current plan is slightly over-shaped for a model task.

The main risk is that it drifts from spec-first behavior into implementation proof:

- the round-trip JSON equality check is stronger than T01 needs
- the schema-invalid input scenario is not part of the T01 spec and reads like a validator test

## Recommendations

1. Keep one BDD scenario centered on the canonical v3 bank unmarshaling into `questionbank.Bank`.
   - This is the only clearly business/spec-first behavior in T01.
   - The `Describe/When/It` should read as a contract statement, not as a serialization harness.

2. Remove the `schema-invalid data is unmarshaled into the model` scenario from the T01 BDD plan.
   - T01 explicitly says the model is representation-only, but that does not require a dedicated negative scenario here.
   - That case adds little user value and starts to look like a validator boundary test, which belongs to T03/T05 style work.

3. Trim the assertion list to representative contract checks instead of proving the whole JSON shape by round-trip comparison.
   - Keep checks for metadata, dimensions, thresholds, question count, representative localized fields, reverse flags, and signed scores.
   - Avoid turning the spec into a full structural mirror of `internal/questionbank/model_test.go`.

4. If you keep any negative control, make it prove the field-preservation assertion can fail, not that the suite is wired.
   - The goal is to validate the behavior assertion, not the Ginkgo harness.

5. Do not expand T01 into loader, validation, scoring, rendering, or CLI coverage.
   - Those belong to later tasks and are already excluded by the OpenSpec and task board.

## Non-Goals

- Do not add production code changes.
- Do not add or modify test code in `internal/questionbank`.
- Do not cover loader behavior, invalid-file handling, validation, scoring, or CLI commands in this review.
- Do not convert the existing TDD test into a mechanical BDD rewrite.

## Follow-up Review

The final T01 Ginkgo spec is aligned with the adopted representative-field shape and does not need production or test-code changes.

Additional workspace cleanup recommendations:

1. Update the focused verification command from `-run TestQuestionBank` to `-run TestQuestionbank`, matching the actual suite entrypoint in `internal/questionbank/questionbank_suite_test.go`.
2. Align `workspace/t01/test_strategy.md` with the final representative-field BDD strategy so it no longer describes round-trip JSON equality or exhaustive per-question assertions.
