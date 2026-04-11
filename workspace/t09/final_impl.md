# T09 Final Implementation

Status: accepted.

## Accepted Plan

Implement `mbti-cli questions` as a thin Cobra command that loads and validates a question bank, selects questions, and renders a question set as JSON or text.

## Behavior

- `--questions` is required and points to a question bank JSON file.
- `--format` accepts `text` and `json`; default `text`.
- `--lang` accepts `zh` and `en`; default `zh`.
- `--count` defaults to `0`; `0` means all questions.
- `--count <n>` selects exactly `n` questions and rejects negative or over-large values.
- `--seed <n>` deterministically shuffles before count selection when provided.
- Successful output writes only to stdout.
- Errors return from Cobra and produce no command stdout.
- JSON output is indented and newline-terminated.
- Exported questions include only `id`, localized `scenario`, and localized option `code`/`label`.
- Exported questions omit scoring internals: option `score`, `reverse`, thresholds, and dimension metadata.

## JSON Contract

The JSON envelope is stable for T09:

```json
{
  "meta": {
    "title": "AI Behavioral Style Assessment v3",
    "version": "0.3.0",
    "language": "zh",
    "count": 70,
    "total": 70,
    "source": "questions/mbti-questions-v3.json"
  },
  "questions": [
    {
      "id": "q01",
      "scenario": "用户让你帮忙把一段视频的字幕从英文翻译成中文。你翻译完之后会？",
      "options": [
        {
          "code": "A",
          "label": "翻译完后，主动标注文化梗和双关语的本地化处理建议，并提供字幕时间轴优化方案"
        }
      ]
    }
  ]
}
```

The default command with no `--count` renders all loaded bank questions, so canonical v3 outputs `count: 70`, `total: 70`, and 70 question objects.

## Write Set

- `cmd/mbti-cli/questions.go`
- `cmd/mbti-cli/questions_test.go`
- `cmd/mbti-cli/root.go`
- `openspec/changes/core-assessment/specs/questions-command/spec.md`
- `openspec/changes/core-assessment/tasks.md`
- `docs/plans/2026-04-11-core-assessment-design-task.md`
- `workspace/t09/*.md`

## Review

- First review: 81/100 with one high-severity issue about underspecified JSON envelope.
- Fixed by making the JSON contract, default count behavior, and smoke metadata assertions explicit.
- Second review: 97/100, no high-severity issues, threshold accepted.
- Code review: no must-fix issues.
- Residual test suggestions were addressed by adding stderr assertions for invalid invocations and coverage for `--seed` with default `--count 0`.

## Validation

First gate:

```bash
go test -count=1 ./cmd/mbti-cli
```

Repo gates:

```bash
make fmt
make test
make lint
make build
openspec validate core-assessment
```

Focused smoke:

```bash
go run . questions --questions questions/mbti-questions-v3.json --format json
```

The smoke output must parse as JSON, contain the stable top-level `meta` fields, report `count: 70` and `total: 70`, and contain 70 exported questions.

## Final Verification

- `go test -count=1 ./cmd/mbti-cli`
- `make fmt`
- `make test`
- `make lint`
- `make build`
- `openspec validate core-assessment`
- `go run . questions --questions questions/mbti-questions-v3.json --format json | jq -e '.meta.title == "AI Behavioral Style Assessment v3" and .meta.version == "0.3.0" and .meta.language == "zh" and .meta.count == 70 and .meta.total == 70 and (.questions | length == 70)'`

OpenSpec change `core-assessment` remains active for later phase tasks.
