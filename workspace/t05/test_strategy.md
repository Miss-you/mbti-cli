# T05 Test Strategy

## What must be proven

- A complete answer set passes strict validation when every selected option code
  exists on its question.
- Unknown answer IDs fail.
- Invalid option codes fail.
- Missing answers fail.
- Multiple issues are reported together in stable order.
- Validation remains bank-aware only; it does not score answers or interpret the
  `reverse` flag.

## First gate

```bash
go test -count=1 ./internal/answers
```

This proves the new validator behavior directly.

## Full gates

```bash
make fmt
make test
make lint
make build
openspec validate core-assessment
```

No focused CLI smoke test is required for T05 because no Cobra command or
stdout/stderr behavior changes in this task.
