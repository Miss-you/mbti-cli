# T09 BDD Plan Review

## Verdict

Mostly aligned, with one coverage gap.

The plan correctly focuses on observable `questions` command behavior:
stdout JSON/text, stderr/error handling, deterministic selection, and omission
of scoring internals.

## Findings

1. The invalid-input plan should include invalid question bank content, not only
   missing paths or files.
   - The `questions` spec requires invalid question bank input to fail without
     partial stdout.
   - The command visibly loads and validates the bank before rendering.

## Recommendations

- Add an invalid bank fixture case to the error table and assert returned error,
  stderr diagnostic, and empty stdout.
- Keep all other planned scenarios at the command boundary.
