# T13 Original Implementation

## Current State

- `README.md` still describes the initial Cobra skeleton command surface.
- The implemented root command now registers `version`, `questions`, and `score`.
- Cobra also exposes the standard `completion` command.
- `questions` requires `--questions <file>` and supports `--format text|json`, `--lang zh|en`, `--count`, and `--seed`.
- `score` requires `--questions <file>` and `--answers <file>` and supports `--format text|json`.

## Drift

The README command list omits `questions` and `score`, so a user reading the project docs would not see the core assessment path that is already implemented and covered by tests.

## Scope

This task is documentation-only. It must not change CLI behavior, tests, OpenSpec specs, question bank data, or scoring semantics.
