# T04 BDD Plan Review

## Verdict

Not recommended as a new BDD scenario in its current form. `internal/answers/parser_ginkgo_test.go` already covers the spec-level parser contract: canonical map parsing, normalization, missing/null `answers`, malformed JSON, structurally invalid values, and deferring bank-aware validation. Adding a separate fixture-backed scenario for `internal/answers/testdata/valid-answers.json` mostly restates the same behavior and overlaps with the existing unit test fixture coverage in `internal/answers/parser_test.go`.

This is only worth keeping if the intent is specifically to lock the repository fixture itself as a regression guard. That is a test-data stability concern, not a new business/spec-first parser behavior.

## Recommendations

- Keep the existing inline Ginkgo scenarios as the primary BDD evidence for the answer parser contract.
- Do not add a second canonical success scenario just to read `valid-answers.json`; it does not add meaningful spec coverage beyond the current BDD and TDD tests.
- If a fixture-backed BDD must exist, use it as a single representative smoke check and keep the assertions minimal, focused on normalization and map shape only.
- Leave the negative-path parser coverage unchanged; it already matches the OpenSpec requirements and is the useful part of the BDD surface.

## Non-Goals

- Do not expand T04 into bank-aware validation, scoring, rendering, or CLI wiring.
- Do not turn the parser BDD into a golden fixture regression suite.
- Do not modify production code or test code as part of this review step.
