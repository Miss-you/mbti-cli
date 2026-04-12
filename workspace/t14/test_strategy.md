# Test Strategy

T14 is a docs/spec cleanup task, so verification focuses on repository state rather than TDD against production code.

Required checks:

- `openspec status --change core-assessment --json` before archive shows all artifacts done.
- `grep -R "## ADDED Requirements" openspec/specs` returns no matches after sync.
- `test ! -d openspec/changes/core-assessment`.
- `test -d openspec/changes/archive/2026-04-12-core-assessment`.
- `openspec list --json` shows no active changes.
- `openspec validate --all`.
- `make fmt`.
- `make test`.
- `make lint`.
- `make build`.

No focused CLI smoke is required because T14 does not change CLI behavior.
