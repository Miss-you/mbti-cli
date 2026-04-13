# T12 BDD Verification

## Decision Verification

No T12-specific Ginkgo file was added.

Reason:

- T12 is a verification and review-repair task, not a product behavior task.
- T12 should close with integrated verification evidence after T07-T10 BDD
  additions.

## Green Verification

Commands:

```bash
go test -count=1 ./...
make fmt
make test
make lint
make build
openspec validate --all
```

Observed results:

- Full Go test: all packages passed.
- `make fmt`: completed.
- `make test`: all packages passed.
- `make lint`: 0 issues.
- `make build`: binary built at `bin/mbti-cli`.
- `openspec validate --all`: 9 specs passed, 0 failed.

## CLI Smoke

Commands:

```bash
./bin/mbti-cli questions --questions questions/mbti-questions-v3.json --count 3 --seed 123 --lang zh --format json
./bin/mbti-cli score --questions questions/mbti-questions-v3.json --answers cmd/mbti-cli/testdata/answers-all-a.json --format json
./bin/mbti-cli questions --questions <malformed-bank> --format json
./bin/mbti-cli score --questions <malformed-bank> --answers cmd/mbti-cli/testdata/answers-all-a.json --format json
```

Observed results:

- Questions success stdout parsed as JSON with `count == 3`,
  `total == 70`, and 3 exported questions; stderr was empty.
- Score success stdout parsed as JSON with type `ESTJ`, answered `70`, and
  total `70`; stderr was empty.
- Both malformed-bank commands exited non-zero, left stdout empty, and wrote
  `parse question bank` to stderr.
