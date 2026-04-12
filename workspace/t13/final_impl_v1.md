# T13 Final Implementation V1

## Decision

Align `README.md` with the implemented first-stage CLI by documenting:

- `mbti-cli questions`
- `mbti-cli score`
- the canonical v3 question bank path
- the canonical answer file JSON shape
- the existing local development commands

## Acceptance

The task is complete when:

- README no longer lists only the skeleton commands.
- README examples use currently implemented flags.
- README says results describe AI behavior style, not a human personality diagnosis.
- `rg -n "mbti-cli (questions|score)|answers|mbti-questions-v3" README.md` finds the documented workflow.
- `make fmt`, `make test`, `make lint`, and `make build` pass.
- Focused `go run` smokes for `questions` and `score` still run.

## Change

`Change=-`; this is docs-only and does not alter the CLI contract, command defaults, file formats, product behavior, or OpenSpec requirements.
