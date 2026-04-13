# T10 BDD Plan Review

## Verdict

Mostly aligned, with two scope corrections.

The plan correctly focuses on observable `score` command behavior: stable
JSON/text result output, strict answer validation failures, and empty stdout on
errors.

## Findings

1. The invalid-input plan should include invalid question bank content.
   - The `score` spec requires invalid question bank input to fail without
     partial stdout.
   - The command visibly loads and validates the bank before reading answers.

2. The CLI BDD plan should not prove the internal reverse/signed-score contract.
   - That behavior belongs in `internal/scoring` BDD and standard tests.
   - T10 should protect the final rendered stdout and command error behavior.

## Recommendations

- Add an invalid bank fixture case to the error table.
- Keep signed-score and reverse-specific assertions at the scoring layer.
