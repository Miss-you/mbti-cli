# T12 BDD Plan

## Decision

Dedicated T12 BDD/Ginkgo coverage is not needed.

T12 is a verification and review-repair task. It does not define a new package
API, command behavior, data contract, renderer contract, or fixture format.
Adding Ginkgo tests for T12 itself would test the repository process rather
than product behavior.

## Source Artifacts

- Task board: `docs/plans/2026-04-11-core-assessment-design-task.md`
- Workspace strategy: `workspace/t12/test_strategy.md`
- Final implementation record: `workspace/t12/final_impl.md`

## Handling

No T12-specific `*_ginkgo_test.go` file is planned.

T12 should be closed by fresh verification after T07-T10 BDD additions:

- focused Ginkgo/package tests for changed packages
- `go test -count=1 ./...`
- repository gates that are available locally
- CLI smoke checks if command BDD additions touched CLI behavior

## Verification

```bash
go test -count=1 ./...
make fmt
make test
make lint
make build
```
