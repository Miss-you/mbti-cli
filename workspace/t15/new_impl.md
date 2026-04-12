# T15 Proposed Implementation

## Approach

Keep `docs/task.md` as a human-readable research index, but add explicit status columns and a short "Phase 1 decisions now frozen" section.

## Classification

- Mark R03, R04, R06, R07, R08, and R09 as resolved for phase 1 because their decisions are captured by the approved design, current OpenSpec specs, README, fixtures, and command tests.
- Keep R01, R02, and R05 as deferred because interactive questionnaire UX, prompt/TUI library choice, and deeper anti-gaming behavior are still outside the archived phase 1 scope.

## Scope boundary

This is docs-only. It must not change CLI behavior, Go code, question banks, or OpenSpec specs.
