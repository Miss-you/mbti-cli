# T02 BDD Code Review

## Verdict

Pass, with one tightening suggestion. `internal/questionbank/loader_ginkgo_test.go` matches the T02 plan and the `writing-ginkgo-bdd-tests` rules overall: the scenarios are behavior-first, the error paths stay on loader contract boundaries, and there is no schema or scoring drift.

## Findings

- [medium] The canonical-load spec still asserts the exact question count (`HaveLen(70)`) in [`internal/questionbank/loader_ginkgo_test.go:20-23`](internal/questionbank/loader_ginkgo_test.go#L20). That is a fixture-specific invariant rather than loader behavior, so it pushes the BDD a little closer to regression coverage of the JSON payload than the plan calls for. The plan only asked for representative metadata, the first question ID, and source metadata.

## Suggestions

- Drop the exact question-count assertion and keep the canonical scenario focused on representative typed-bank fields plus source metadata.
- Keep the existing missing-file, malformed-JSON, and empty-path scenarios as they are; they map cleanly to the loader contract and read well as Ginkgo specs.
- If you want a stronger behavior signal without widening scope, keep the title/version checks, first-question ID, and source metadata checks. That is enough to prove the loader contract without binding the spec to fixture size.
