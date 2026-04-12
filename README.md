# mbti-cli

`mbti-cli` is a small Go CLI for coding-related workflows, built with Cobra and kept intentionally minimal for long-term maintenance.

The current assessment flow describes AI behavior style, not a human personality diagnosis.

## 当前命令

- `mbti-cli`
- `mbti-cli help`
- `mbti-cli version`
- `mbti-cli questions`
- `mbti-cli score`
- `mbti-cli completion`

## 导出题目

`questions` 从题库导出题目，默认输出中文 text，也可以输出稳定 JSON。

```bash
mbti-cli questions \
  --questions questions/mbti-questions-v3.json \
  --count 70 \
  --seed 123 \
  --lang zh \
  --format json
```

常用参数：

- `--questions`: 题库 JSON 路径，当前 canonical 文件是 `questions/mbti-questions-v3.json`
- `--format`: `text` 或 `json`
- `--lang`: `zh` 或 `en`
- `--count`: 导出题目数量，`0` 表示全部题目
- `--seed`: 需要稳定乱序时使用的 seed

## 评分

`score` 读取题库和 answer file，输出 text 或 JSON 结果。

```bash
mbti-cli score \
  --questions questions/mbti-questions-v3.json \
  --answers answers.json \
  --format json
```

answer file 使用 `answers` map。每个 key 是 question ID，每个 value 是 option code `A/B/C/D`。

```json
{
  "answers": {
    "q01": "A",
    "q02": "C"
  }
}
```

第一阶段默认要求被评分题目都有答案；unknown question ID、unknown option code 或缺失答案都会返回非零退出码，并把错误信息写到 stderr。

## 目录结构

- `main.go`: 程序入口
- `cmd/mbti-cli/`: Cobra 命令定义
- `internal/cli/`: 可复用的命令执行封装
- `internal/questionbank/`: 题库模型、加载和 schema validation
- `internal/answers/`: answer file 解析和验证
- `internal/scoring/`: 维度聚合、阈值分类和 type 生成
- `internal/result/`: text / JSON 结果渲染
- `internal/version/`: 版本信息与构建注入变量
- `questions/`: 本地题库数据
- `docs/plans/`: 设计与实施计划

## 开发

```bash
make fmt
make test
make lint
make build
make run
```

`make run` 会复用与 `make build` 一致的版本注入参数，便于本地做接近发布态的 smoke test。
如果需要传子命令，可使用 `make run ARGS=version`。
例如：

```bash
make run ARGS='questions --questions questions/mbti-questions-v3.json --count 1 --format text'
```

## 构建产物

默认构建产物位于 `bin/mbti-cli`。

## License

Apache 2.0. See [LICENSE](LICENSE).
