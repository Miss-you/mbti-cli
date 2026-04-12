# T06 BDD Plan Review Evaluation

## Adopted

- Adopt: keep the BDD surface narrow and observable. `Answered`, `Total`, and `DimensionScores` are the right runtime fields for T06, and they match the scoring contract without inventing extra shape checks.
- Adopt: keep the reversed-question scenario focused on stored signed scores. That proves scoring uses the score already attached to the option instead of applying `reverse` a second time.
- Adopt: treat classifier, type, and renderer exclusion as a non-goal. Reflection-based absence assertions would add brittleness without proving a user-visible scoring behavior.
- Adopt: anchor the incomplete-answer case to the existing answer-validation contract and stop at zero-result behavior. That keeps the spec focused on early failure before partial totals leak through.

## Rejected

- None. The review recommendations are compatible with the current plan and do not require expanding scope.

## Final BDD Shape

- `Describe("Scoring engine")`
- `When("a complete strict answer set is scored")`
  - `It("adds selected option scores by dimension and reports answer counts")`
- `When("a selected option belongs to a reversed question")`
  - `It("adds the stored signed score without applying reverse again")`
- `When("answers are incomplete")`
  - `It("returns a validation error without partial dimension totals")`

The current [`workspace/t06/bdd_plan.md`](/Users/apple/Documents/Github/mbti-cli/workspace/t06/bdd_plan.md) already matches this shape, so no plan update was needed.
