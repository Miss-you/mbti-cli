# T09 BDD Plan Review Evaluation

## Adopted

1. Adopt: add invalid question bank content coverage.
   - Reason: the OpenSpec command contract requires invalid bank input to fail
     without partial stdout.

2. Adopt: keep the remaining plan at the CLI boundary.
   - Reason: JSON shape, localized text, deterministic selection, and stderr
     behavior are the observable T09 contract.

## Rejected

1. Reject leaving invalid-input coverage limited to format, language, count,
   path, and missing-file cases.
   - Reason: that would miss the command's validation failure path.

## Plan Adjustment

`workspace/t09/bdd_plan.md` now includes invalid question bank content in the
invalid-input BDD scenario.
