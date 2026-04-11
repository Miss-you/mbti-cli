# T10 Test Strategy

## Proof Targets

1. `score` is wired into the root command and can score canonical v3 answers as JSON.
2. JSON stdout is parseable, newline-terminated, and uses stable result fields.
3. Text stdout uses the existing factual summary renderer.
4. Invalid inputs produce errors and no command stdout.
5. The implementation uses strict answer validation through existing core packages.

## TDD First Gate

Run:

```bash
go test -count=1 ./cmd/mbti-cli
```

The first red test should fail because the `score` command is not registered.

## Repo Gates

Run fresh after implementation:

```bash
make fmt
make test
make lint
make build
openspec validate core-assessment
```

## Focused CLI Smoke

Create a temporary complete answer file for q01 through q70, run:

```bash
go run . score --questions questions/mbti-questions-v3.json --answers <tmp> --format json
```

Then parse stdout with `jq` and verify:

- `.meta.answered == 70`
- `.meta.total == 70`
- `.type == "ESTJ"`
- `.dimensions.EI.score == 36`
- `.dimensions.SN.score == 36`
- `.dimensions.TF.score == 34`
- `.dimensions.JP.score == 34`
