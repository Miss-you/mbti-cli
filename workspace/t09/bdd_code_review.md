# T09 BDD Code Review

## Verdict

No must-fix findings.

`cmd/mbti-cli/questions_command_ginkgo_test.go` stays at the CLI boundary:

- canonical JSON export is parseable and omits scoring internals
- English text export honors `--count 1` and `--lang en`
- seeded bounded selection is deterministic
- invalid inputs, including invalid bank content, return an error, write a
  stderr diagnostic, and leave stdout empty

## Findings

None.
