# T05 BDD Code Review

## Verdict

The T05 BDD suite is broadly aligned with the plan and with `writing-ginkgo-bdd-tests`.
It keeps the coverage on `Validate`, preserves the lowercase case as a validator-level negative control, and avoids drifting into parser or scoring behavior.

One part still reads heavier than necessary: the shared bank fixture carries fields that `Validate` does not observe.

## Findings

1. **Medium: the bank fixture is overbuilt for a validator spec.**
   In `internal/answers/validator_ginkgo_test.go:12-53`, the helper builds `Meta`, `Scenario`, `Reverse`, and per-option `Score` values even though `Validate` only inspects question IDs and option codes.
   That makes the spec look closer to a model-preservation or scoring fixture than a behavior-first validator check.
   The lowercase negative control is still framed correctly as direct validator input, but the surrounding bank data is richer than the contract needs.

## Suggestions

- Trim the shared fixture to the smallest bank shape that still supports the validator contract: question IDs plus option codes are enough.
- Keep the lowercase case exactly where it is conceptually, as a direct `Validate` boundary check, but make the title read like a validator rejection rather than a parser compatibility test.
- Leave the aggregated-issues scenario in place; it is the right kind of BDD evidence for deterministic validation order.
