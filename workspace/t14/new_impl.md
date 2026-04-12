# Proposed Implementation

Treat T14 as a docs/spec cleanup task with no product behavior change.

Steps:

1. Add a T14 task-board entry for OpenSpec archive cleanup.
2. Sync each delta spec from `openspec/changes/core-assessment/specs/<capability>/spec.md` into `openspec/specs/<capability>/spec.md`.
3. Because all delta specs contain only `## ADDED Requirements`, create each main spec with the same requirement content under `## Requirements`.
4. Move `openspec/changes/core-assessment/` to `openspec/changes/archive/2026-04-12-core-assessment/`.
5. Verify OpenSpec no longer lists an active `core-assessment` change and the archived specs remain valid.
