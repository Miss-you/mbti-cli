# T03 BDD Code Review

## Verdict
Needs tightening.

The representative invalid-schema table is broadly aligned with the plan, but the canonical success case still reaches beyond validator behavior and into model/scoring preservation.

## Findings

1. `internal/questionbank/validator_ginkgo_test.go:20-25` asserts `Reverse` and `Score` values after `Validate(bank)` succeeds.
   These checks do not validate the validator contract itself. They are effectively model/unmarshal assertions about fixture contents, and they couple the BDD to fixture internals that can change without affecting schema validation. The plan only requires proving that the canonical v3 bank validates successfully and that scoring behavior is not invoked.

## Suggestions

- Keep the positive-case BDD focused on `Validate(bank)` succeeding.
- If reverse/score preservation is worth protecting, move those assertions to a separate model/unmarshal or loader-level test, not the validator BDD.
- Leave the invalid-schema table in place; its representative cases match the scenario-group approach from the plan.
