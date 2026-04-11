# Core Assessment 第一阶段设计

## 目标

实现 `mbti-cli` 的第一阶段 core assessment 闭环：

1. 加载题库
2. 校验 schema
3. 接收 answer file 或内存答案
4. 打分
5. 输出 text / JSON 结果
6. 用 fixtures 覆盖 scoring 与输出契约

这一阶段不等待完整 CLI/TUI 调研。目标是先把可验证、可自动化消费的核心路径跑通。

## 非目标

- 不做交互式 `assess` 问卷流程
- 不引入 `survey`、`huh`、`bubbletea` 等 TUI/prompt 库
- 不调用 AI 模型
- 不做 profile/history 持久化
- 不做 Markdown report
- 不做 plugin/provider 抽象
- 不把题库嵌入二进制

## 当前项目盘点

当前仓库事实：

- Go Cobra skeleton 已存在，当前命令包括 root help 和 `version`
- 题库位于 `questions/`
- 当前题库文件包括 `mbti-questions.json`、`mbti-questions-v2.json`、`mbti-questions-v3.json`
- 三个题库实际都是 70 题
- `questions/mbti-questions-v3.json` 的版本是 `0.3.0`
- v3 维度分布为 `EI=18`、`SN=18`、`TF=17`、`JP=17`
- v3 已包含 `meta.dimensions`、`meta.scoring.thresholds`、`questions[].scenario`、`questions[].options[].score`
- 当前 `reverse` 字段存在，但 option score 已经是 signed score
- 当前 Go 代码还没有 question/scoring/answer/result core package

证据：

- `questions/mbti-questions-v3.json`
- `cmd/mbti-cli/root.go`
- `cmd/mbti-cli/version.go`
- `cmd/mbti-cli/root_test.go`
- `docs/project-goals.md`
- `docs/ref/mbti.md`
- `docs/task.md`

## 兼容性契约

第一阶段必须冻结并保护以下契约。

### 题库契约

- 默认 canonical bank 为 `questions/mbti-questions-v3.json`
- 题库结构以 JSON 中的 `meta` 和 `questions` 为准
- `meta.total` 必须等于 `len(questions)`
- `meta.dimensions[dim].count` 必须等于实际题数
- 支持的维度为 `EI`、`SN`、`TF`、`JP`
- 每题必须有唯一 `id`
- 每题必须有有效 `dimension`
- 每题必须有 `scenario.zh` 和 `scenario.en`
- 每题必须有 4 个 option
- option code 必须是 `A/B/C/D`
- option score 必须是 `-2/-1/1/2`
- threshold buckets 必须完整且互不重叠

### 打分契约

- option `score` 是 authoritative signed score
- `reverse` 在第一阶段只作为 metadata 保存和校验，不参与二次反转
- 每道已回答题把所选 option score 加到该题 dimension total
- 维度顺序固定为 `EI`、`SN`、`TF`、`JP`
- score `> 0` 选 pole A
- score `< 0` 选 pole B
- score `== 0` 选 `balanced`，type letter 使用 `X`
- 阈值分类来自题库 `meta.scoring.thresholds`
- 最终 type 由四个维度 letter 拼接，例如 `ENTJ` 或 `EXTP`

### 答案契约

第一阶段 canonical answer file 使用 map 形态：

```json
{
  "answers": {
    "q01": "A",
    "q02": "C"
  }
}
```

规则：

- option code 大小写可标准化为大写
- unknown question ID 报错
- unknown option code 报错
- 默认 strict mode 要求被评分题目全部有答案
- 第一阶段不支持 partial result，除非后续显式设计

### CLI 契约

第一阶段不做交互式 `assess`，而是提供两个 coding CLI / AI Agent 友好的非交互命令：

1. `mbti-cli questions`
2. `mbti-cli score`

`questions` 负责从题库导出题目集：

```bash
mbti-cli questions \
  --questions questions/mbti-questions-v3.json \
  --count 70 \
  --seed 123 \
  --lang zh \
  --format json
```

`score` 负责读取答案并输出结果：

```bash
mbti-cli score \
  --questions questions/mbti-questions-v3.json \
  --answers answers.json \
  --format json
```

stdout / stderr：

- 成功输出只写 stdout
- 错误诊断只写 stderr
- JSON stdout 不混入人类 warning
- 第一阶段 exit code 保持简单：`0` success，`1` data / IO / validation error

### JSON 结果契约

`--format json` 输出应保持稳定，便于自动化消费和 golden test。

建议结构：

```json
{
  "meta": {
    "title": "AI Behavioral Style Assessment v3",
    "version": "0.3.0",
    "answered": 70,
    "total": 70
  },
  "type": "ENTJ",
  "dimensions": {
    "EI": {
      "letter": "E",
      "score": 8,
      "strength": "moderate_a",
      "pole": "E (Expansive)"
    }
  }
}
```

## 术语映射

| 术语 | 含义 |
| --- | --- |
| Question bank | `questions/*.json` 中的完整题库 |
| Selected question set | `questions` 命令输出的题目集合 |
| Answer file | AI Agent 或人类提交的答案 JSON |
| Dimension score | 某个维度下所有回答 option score 的和 |
| Strength | 根据 thresholds 得出的强弱桶 |
| Type | 按 `EI/SN/TF/JP` 顺序拼出的风格标签 |
| Balanced | 某维度 score 为 `0`，type letter 为 `X` |

## Core 与 Shell 边界

### Core packages

推荐包边界：

- `internal/questionbank`
  - JSON model
  - loader
  - schema validator
- `internal/answers`
  - answer file model
  - answer parser
  - answer validation
- `internal/scoring`
  - score aggregation
  - threshold classification
  - type generation
- `internal/result`
  - result DTO
  - text renderer
  - JSON renderer

### CLI shell

`cmd/mbti-cli` 只负责：

- Cobra command wiring
- flags
- path handling
- stdout / stderr
- translating errors to exit behavior

不要把 scoring 或 validation 逻辑塞进 Cobra command。

## 阶段计划

### Phase 1. Data model, loader, validator

目标：

- 能加载 v3 题库
- 能拒绝明显非法 schema
- 不丢失 scoring 所需字段

验收：

- `questions/mbti-questions-v3.json` 能通过 validation
- invalid fixtures 能被明确拒绝

### Phase 2. Answer model and scoring

目标：

- 能读取 canonical answer file
- 能验证 answers
- 能输出 deterministic score result

验收：

- all-A / all-D / mixed fixture 得到稳定结果
- threshold boundary test 通过
- zero score 返回 `balanced` / `X`

### Phase 3. Result rendering

目标：

- 输出 human-readable text
- 输出 stable JSON

验收：

- text output 可读且不说人格诊断
- JSON output 字段稳定，可 golden test

### Phase 4. CLI integration

目标：

- 增加 `questions` 和 `score` 命令
- 支持非交互 AI Agent 流程

验收：

- `mbti-cli questions --questions <file> --format json` 输出题目 JSON
- `mbti-cli score --questions <file> --answers <file> --format text|json` 输出结果
- invalid input 走 stderr + non-zero

### Phase 5. Verification and review repair

目标：

- 用 tests 和 repo gates 证明 core 闭环稳定

验收：

- `make fmt`
- `make test`
- `make lint`
- `make build`
- focused CLI smoke test
- review must-fix 问题已修复并复验

## 粗任务拆解

实现任务会在 paired task board 中细化，粗分为：

1. 建 question bank model
2. 建 loader
3. 建 schema validator
4. 建 answer model/parser
5. 建 scoring engine
6. 建 threshold classifier
7. 建 result model/renderers
8. 接入 `questions` command
9. 接入 `score` command
10. 添加 fixtures/golden tests
11. 跑全量 verification 和 review repair

## Multi-agent 执行模型

执行时使用 owner + 多 agent 方式，但避免写集冲突。

推荐并行方式：

- Agent A: `internal/questionbank` model / loader / validator
- Agent B: `internal/answers` parser / validation
- Agent C: `internal/scoring` score / threshold / type
- Agent D: `internal/result` renderer / golden output
- Owner: `cmd/mbti-cli` command integration and final merge

并行约束：

- 同一 package 同一时间只允许一个 worker 拥有写权限
- CLI integration 等所有 core API 基本稳定后再做
- fixtures 和 golden tests 由 owner 统一整合，避免输出契约漂移

## 验证计划

### Unit tests

- loader loads v3 bank
- validator accepts v3 bank
- validator rejects invalid fixtures
- answer parser accepts canonical map JSON
- answer validator rejects unknown IDs/options/missing answers
- scorer returns expected dimension totals
- threshold classifier covers boundary values
- zero score returns `balanced` and `X`
- type generation uses `EI/SN/TF/JP` order

### Integration tests

- production v3 bank loads and validates
- fixed answer fixture scores deterministically
- CLI `questions --format json` emits parseable JSON
- CLI `score --format json` emits stable JSON
- CLI `score --format text` emits readable summary

### Repo gates

完成前必须运行：

```bash
make fmt
make test
make lint
make build
```

## 验收标准

第一阶段完成时必须满足：

- `questions/mbti-questions-v3.json` 可加载并验证
- canonical answer file 可解析
- scoring 结果可复现
- `reverse` 未被重复应用
- zero score 行为明确
- text / JSON 输出稳定
- CLI command 只做 orchestration
- 所有规定 tests 和 repo gates fresh pass
- review 中的 correctness / regression 问题已修复并复验

## 延后调研

这些保留到第二阶段：

- 交互式问卷 UX
- prompt/TUI library 选型
- 回退修改上一题
- AI 模型直接调用
- 更复杂 anti-gaming 策略
- Markdown report
- profile/history
- 多次结果比较

## 证据引用

- `/Users/lihui/Documents/GitHub/mbti-cli/questions/mbti-questions-v3.json`
- `/Users/lihui/Documents/GitHub/mbti-cli/cmd/mbti-cli/root.go`
- `/Users/lihui/Documents/GitHub/mbti-cli/cmd/mbti-cli/root_test.go`
- `/Users/lihui/Documents/GitHub/mbti-cli/docs/project-goals.md`
- `/Users/lihui/Documents/GitHub/mbti-cli/docs/ref/mbti.md`
- `/Users/lihui/Documents/GitHub/mbti-cli/docs/task.md`
