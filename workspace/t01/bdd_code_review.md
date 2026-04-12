# T01 BDD Code Review

## Verdict

Needs revision.

The Ginkgo test captures the right T01 capability, but the implementation goes beyond the adopted final plan and turns the spec into a near-full structural mirror of the fixture. That weakens the BDD value and adds brittle assertions that the plan explicitly asked to avoid.

## Findings

1. The spec exceeds the adopted assertion scope by asserting full JSON round-trip equality and then walking every question.

   In [internal/questionbank/model_ginkgo_test.go](/Users/apple/Documents/Github/mbti-cli/internal/questionbank/model_ginkgo_test.go:21), the test marshals the model back to JSON and compares the entire decoded payload with the original raw fixture. Lines 31-70 then add exhaustive checks across title, localized title, full dimension metadata, both threshold extremes, the full question count, the first question, the second question, and the last question.

   That is broader than the final plan in [workspace/t01/bdd_plan_evaluation.md](/Users/apple/Documents/Github/mbti-cli/workspace/t01/bdd_plan_evaluation.md:24), which limited T01 to representative fields: title, version, total count, dimensions, threshold ranges, one first question, one reversed question, and one representative option label plus signed score. The current test is still correct as a contract check, but it no longer matches the adopted BDD shape and is drifting toward implementation-detail coverage.

2. The suite naming and documented focused verification are slightly out of sync.

   [internal/questionbank/questionbank_suite_test.go](/Users/apple/Documents/Github/mbti-cli/internal/questionbank/questionbank_suite_test.go:8) registers `TestQuestionbank`, while [workspace/t01/bdd_plan.md](/Users/apple/Documents/Github/mbti-cli/workspace/t01/bdd_plan.md:58) documents the focused command with `-run TestQuestionBank`. That mismatch makes the plan's negative-control command unreliable as written, so the verification path in the workspace artifacts does not exactly describe the implemented suite.

## Suggestions

- Remove the full `raw` vs `roundTripped` equality check and keep only representative field assertions.
- Trim the per-question assertions to the minimum set the plan adopted: one canonical first question, one reversed question, and one representative option label plus signed score.
- Keep the BDD title and `It` text centered on the observable contract, not on field preservation in the abstract.
- Align the suite name and any documented `go test -run` command so the focused verification path actually selects this package.
