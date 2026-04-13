# T10 BDD Plan Review Evaluation

## Adopted

1. Adopt: add invalid question bank content coverage.
   - Reason: the OpenSpec command contract requires invalid bank input to fail
     without partial stdout.

2. Adopt: remove CLI-level proof of signed option scores being scored once.
   - Reason: the CLI BDD should assert observable stdout/stderr/error behavior;
     reverse and signed-score semantics are owned by `internal/scoring`.

## Rejected

1. Reject keeping the reverse/signed-score internal contract inside the CLI BDD
   plan.
   - Reason: it would make command BDD test implementation details rather than
     the command boundary.

## Plan Adjustment

`workspace/t10/bdd_plan.md` now includes invalid question bank content in the
invalid-input BDD scenario and describes the command as delegating to core
packages rather than proving scoring internals.
