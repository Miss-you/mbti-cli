# T01 Original Implementation Notes

## Current State

- The repository has Cobra root/version commands only.
- No `internal/questionbank` package exists.
- No Go type currently represents `questions/*.json`.
- In the main workspace, `questions/mbti-questions-v3.json` exists as an untracked file. The isolated worktree does not contain it yet.

## Existing Contracts From Design

- Canonical bank: `questions/mbti-questions-v3.json`.
- Top-level JSON shape: `meta` and `questions`.
- `meta` includes title/version/total, dimension metadata, and scoring thresholds.
- Each question includes `id`, `dimension`, `reverse`, localized `scenario`, and four localized scored options.
- Option `score` is signed and authoritative. `reverse` must be preserved as metadata for later tasks, not interpreted here.

## Gap For This Task

`T01` must provide Go structs that can unmarshal the v3 question bank without dropping fields needed by later loader, validator, answer, and scoring tasks.
