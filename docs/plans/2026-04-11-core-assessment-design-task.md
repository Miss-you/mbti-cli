# Core Assessment 第一阶段设计任务板

## 源设计

- `docs/plans/2026-04-11-core-assessment-design.md`

## 状态说明

- `todo`: 尚未开始，依赖未必满足
- `claimed`: 已认领，正在建立上下文
- `research`: 正在调研或确认实现细节
- `spec`: 正在固化 OpenSpec 或测试策略
- `implementing`: 正在实现
- `verifying`: 正在验证
- `review`: 正在 review 与修复
- `blocked`: 被阻塞
- `done`: 已完成且有验证证据

## 依赖规则

- 只有 `status=todo` 且所有 `Depends On` 都 `done` 的任务可认领
- `ready` 是推导状态，不写入表格
- T01-T04 可在设计批准后并行
- T05-T07 依赖对应 core model 基础
- T08-T09 等 core API 基本稳定后再做
- T10-T11 是收口任务，不应提前开始

## 任务表

| ID | 标题 | 目标 | 依赖 | 并行组 | 状态 | Owner | 认领时间 | Workspace | Change | 完成条件 | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| T01 | Question bank model | 定义题库 JSON 对应 Go structs，覆盖 meta、dimensions、thresholds、questions、options、reverse | - | A | done | codex | 2026-04-11 13:24 CST | workspace/t01/ | core-assessment | v3 题库可无字段丢失地 unmarshal | done: review 97/100; `make fmt`, `make test`, `PATH=/tmp/mbti-cli-tools/bin:$PATH make lint`, `make build`, `go test -count=1 ./...`, `openspec validate core-assessment`; change kept active for later phase tasks |
| T02 | Question bank loader | 从文件路径读取题库，返回 typed bank 和 source metadata | T01 | A | done | codex | 2026-04-11 16:38 CST | workspace/t02/ | core-assessment | valid file loads，missing/malformed file returns clear error | done: review 94/100; `make fmt`, `go test -count=1 ./internal/questionbank`, `make test`, `go test -count=1 ./...`, `make lint`, `make build`, `openspec validate core-assessment`; change kept active for later phase tasks |
| T03 | Schema validator | 校验题库结构、维度计数、option code/score、threshold 完整性 | T01,T02 | A | done | codex | 2026-04-11 17:41 CST | workspace/t03/ | core-assessment | v3 validates，invalid fixtures fail for expected reasons | done: review fixed spec/test_strategy threshold overlap wording; `make fmt`, `go test -count=1 ./internal/questionbank`, `go test -count=1 ./...`, `make test`, `make lint`, `make build`, `openspec validate core-assessment`; no CLI smoke needed; change kept active for later phase tasks |
| T04 | Answer model and parser | 定义 canonical answer file 并解析为 `questionID -> optionCode` | T01 | B | done | codex | 2026-04-11 17:19 CST | workspace/t04/ | core-assessment | map-form answer file 可解析，option code 规范化 | done: `go test -count=1 ./internal/answers`, `openspec validate core-assessment`; plan review 89/100, spec review 96/100, code review accepted; change kept active for later phase tasks |
| T05 | Answer validation | 对照题库验证 unknown ID、invalid option、missing answer | T03,T04 | B | done | codex | 2026-04-12 00:14 CST | workspace/t05/ | core-assessment | strict mode 下非法和缺失答案都有明确错误 | done: TDD red/green; plan review 88/100; spec review 92/100; code review no findings; `make fmt`, `make test`, `make lint`, `make build`, `openspec validate core-assessment`; no CLI smoke needed; assumes bank passed T03 schema validation; change kept active for later phase tasks |
| T06 | Scoring engine | 汇总 dimension score，明确 reverse 不二次应用 | T03,T05 | C | todo | - | - | workspace/t06/ | core-assessment | fixture answers 产生 deterministic dimension totals | first gate: scoring tests |
| T07 | Threshold and type classifier | 按 thresholds 分类 strength，按 EI/SN/TF/JP 生成 type，zero 使用 X/balanced | T06 | C | todo | - | - | workspace/t07/ | core-assessment | 所有 threshold boundary 和 zero case 都有测试 | first gate: classifier table tests |
| T08 | Result renderers | 输出 stable JSON 和 readable text summary | T07 | D | todo | - | - | workspace/t08/ | core-assessment | JSON golden stable，text 不暗示人格诊断 | first gate: renderer golden tests |
| T09 | Questions command | 接入 `mbti-cli questions`，输出 selected question set text/json | T03,T08 | owner | todo | - | - | workspace/t09/ | core-assessment | `questions --questions <file> --format json` 输出 parseable JSON | first gate: Cobra command tests |
| T10 | Score command | 接入 `mbti-cli score`，读取 answer file 并输出 text/json result | T05,T08 | owner | todo | - | - | workspace/t10/ | core-assessment | `score --questions <file> --answers <file> --format text|json` 可跑通 | first gate: CLI smoke tests |
| T11 | Fixtures and golden tests | 建 valid/invalid bank、answer、threshold、golden output fixtures | T03,T05,T08,T10 | owner | todo | - | - | workspace/t11/ | core-assessment | fixtures 覆盖 schema、answers、scoring、rendering、CLI | first gate: `go test ./...` |
| T12 | Verification and review repair | 全量验证、代码 review、修复并复验 | T01,T02,T03,T04,T05,T06,T07,T08,T09,T10,T11 | solo | todo | - | - | workspace/t12/ | core-assessment | fmt/test/lint/build fresh pass，review must-fix 已处理 | gates: `make fmt && make test && make lint && make build` |

## 认领规则

- 先更新本任务板，再创建 `workspace/<task-id>/`
- 每个任务只有一个 owner
- 并行只能发生在写集不重叠时
- T09-T10 由 owner 集成，避免 Cobra command 层和 core API 同时漂移
- 任一任务进入 `done` 前必须写明验证证据

## 执行 SOP

每个任务都按同一闭环执行：

1. 明确本任务输入、输出和写集
2. 写或更新测试，先证明目标行为
3. 实现最小代码
4. 运行本任务 first gate
5. 修复失败
6. 复跑 first gate
7. 更新任务板状态
8. 进入 review 或交接下一任务

## 验收清单

整体完成前必须确认：

- v3 题库通过 loader + validator
- invalid fixtures 覆盖主要 schema 错误
- canonical answer map JSON 可解析和验证
- scoring 未重复应用 `reverse`
- threshold boundary 完整覆盖
- zero score 行为为 `balanced` / `X`
- JSON output 可 golden test
- text output 不暗示人格诊断
- CLI stdout/stderr 分离
- `questions` 和 `score` 命令都有 command tests
- `make fmt`
- `make test`
- `make lint`
- `make build`

## Change Log

- 2026-04-11: 初始化第一阶段 core assessment 任务板，冻结 12 个可执行任务。
