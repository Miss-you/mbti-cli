# T07 BDD Plan Review

## Verdict

No must-fix findings.

The plan stays at the correct behavior layer for T07. It targets the
`Classify` contract from `openspec/specs/scoring-classifier/spec.md` and avoids
CLI, renderer, or raw scoring behavior.

## Recommendations

- Keep the eventual Ginkgo specs narrowly focused on classifier behavior:
  threshold boundaries, positive/negative pole mapping, balanced zero,
  deterministic type order, and contextual classifier errors.
- Do not expand this task into score aggregation, result rendering, or command
  behavior.
