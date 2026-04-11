# T09 Original Implementation

Status: researched.

## Current Behavior

- `cmd/mbti-cli` has a Cobra root command and a `version` subcommand.
- There is no `questions` subcommand yet.
- Root command help is covered by `cmd/mbti-cli/root_test.go`.
- `internal/questionbank` already owns:
  - typed question bank JSON model,
  - `LoadFile(path)` filesystem loading,
  - `Validate(bank)` schema validation.
- `internal/result` owns score result rendering only and does not help with question set export.

## Existing Contracts Available to T09

- `questions/mbti-questions-v3.json` is the canonical first-phase bank.
- `questionbank.LoadFile` reports path-specific read/parse errors.
- `questionbank.Validate` aggregates schema errors and does not apply scoring.
- Cobra command code should remain orchestration only.

## Gap

T09 must add `mbti-cli questions` so an agent can fetch a selected question set from a bank file. The command needs command-level tests proving JSON is parseable, stdout/stderr are separated, and invalid bank inputs return errors instead of partial output.
