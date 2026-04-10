# mbti-cli Design

## Goal

Initialize `mbti-cli` as a maintainable Go-based open source CLI project with a Cobra command tree, basic unit tests, linting, build automation, and an Apache 2.0 license.

## Scope

- Use Go module path `github.com/Miss-you/mbti-cli`
- Use Cobra for the command framework
- Keep the initial command surface minimal: root command and `version`
- Add only the project structure needed for long-term maintenance
- Ensure local development has a single entry point through `Makefile`

## Architecture

### Command layout

- `main.go` is the process entry point
- `cmd/mbti-cli/` contains Cobra commands
- `internal/cli/` holds reusable execution helpers
- `internal/version/` holds build-time version metadata

### Tooling

- `Makefile` provides `fmt`, `test`, `lint`, `build`, and `run`
- `golangci-lint` provides the linting baseline
- `testify` is used for unit tests

### Testing

- Add focused unit tests for reusable command behavior
- Keep command construction testable without starting subprocesses
- Validate both default version output and build-time injected metadata

## Non-goals

- No release automation in this pass
- No CI workflow in this pass
- No extra subcommands beyond the minimum runnable skeleton
- No heavy configuration layering or framework abstraction

## Maintenance principles

- Prefer the standard library unless Cobra materially simplifies command wiring
- Keep logic out of `main.go`
- Keep build metadata isolated so later release tooling can inject values cleanly
- Avoid premature package splitting or custom framework wrappers

