# mbti-cli Human Task List

## Purpose

这份文档记录 `mbti-cli` 的人工调研任务及当前状态。

它面向人阅读，不是 agent task board。最初目标是帮助实现前冻结关键产品决策；现在第一阶段 core assessment 已完成，本文件同时记录哪些决策已经冻结，哪些仍留给后续阶段。

## Project Context

`mbti-cli` 的目标是制作一个面向 Coding CLI / AI Agent 的 MBTI-style CLI 工具。

主要题目来源是仓库内的 `questions/` 目录，产品参考见：

- `docs/project-goals.md`
- `docs/ref/mbti.md`

## Current Status

第一阶段 core assessment 已完成并归档，当前稳定面包括：

- canonical 题库：`questions/mbti-questions-v3.json`
- 非交互 `questions` 命令，用于导出不含 scoring internals 的题目集
- canonical answer file：`{"answers": {"q01": "A"}}`
- 非交互 `score` 命令，用于读取 answer file 并输出 text / JSON 结果
- 题库 schema、answer validation、scoring、threshold classifier、result renderer 的 OpenSpec 主规格
- fixtures、golden tests、README usage 和非人格诊断措辞

决策来源：

- `docs/plans/2026-04-11-core-assessment-design.md`
- `docs/plans/2026-04-11-core-assessment-design-task.md`
- `openspec/specs/`
- `README.md`
- `cmd/mbti-cli/testdata/` 和 `internal/*/testdata/`

## Research Tasks

| ID | Topic | Status | Goal | Suggested Search Terms | Decision Source | Output / Done When |
| --- | --- | --- | --- | --- | --- | --- |
| R01 | 终端问卷 UX | deferred | 明确 CLI 单选题、进度、回退和交互式结果展示体验 | `CLI questionnaire UX`, `terminal survey single choice`, `command line quiz UX`, `interactive CLI best practices` | `docs/plans/2026-04-11-core-assessment-design.md` 非目标明确不做交互式 `assess` | 后续交互式阶段需要 1 页 UX 决策摘要，确定是否支持回退、进度显示和逐题交互 |
| R02 | Go prompt/TUI 库 | deferred | 判断交互式阶段是否引入 `survey`、`huh`、`bubbletea` 等库，还是继续使用 stdin/stdout | `Go survey library`, `charmbracelet huh`, `bubbletea form`, `Go interactive CLI prompt` | `docs/plans/2026-04-11-core-assessment-design.md` 未引入 prompt/TUI 库，Cobra + file/stdout 路径已经满足非交互流程 | 后续交互式阶段需要库选型对比表和不选其他库的理由 |
| R03 | 题库 JSON schema | resolved for phase 1 | 明确题库版本、多语言、维度、选项、反向题、阈值和兼容策略 | `questionnaire JSON schema`, `survey JSON format`, `quiz schema versioning`, `JSON schema backward compatibility` | `openspec/specs/question-bank-model/spec.md`, `openspec/specs/question-bank-validator/spec.md`, `internal/questionbank/` | phase 1 validation rules 已冻结；后续只有新增题库版本或 schema breaking change 时才重开 |
| R04 | MBTI-like scoring | resolved for phase 1 | 明确维度分数如何映射到方向、强弱倾向、类型码和解释文本 | `MBTI scoring dimensions`, `personality type scoring thresholds`, `Likert scoring classification`, `type indicator scoring` | `openspec/specs/scoring-engine/spec.md`, `openspec/specs/scoring-classifier/spec.md`, `internal/scoring/` | phase 1 scoring、threshold boundary、zero=`balanced`/`X` 已冻结并有测试 |
| R05 | LLM anti-gaming | deferred | 降低模型识别测试意图、迎合题目或按显性维度作答的概率 | `LLM benchmark prompt leakage`, `LLM evaluation anti gaming`, `questionnaire reverse scoring`, `survey response bias mitigation` | phase 1 已通过 `questions` 输出隐藏 scores、reverse、thresholds 和维度 metadata；更深策略仍未设计 | 后续需要题目呈现和乱序策略，确认是否引入更强 anti-gaming 约束 |
| R06 | Agent 作答模式 | resolved for phase 1 | 判断 Coding CLI / AI Agent 更适合逐题交互、一次性 prompt、answer file 还是 stdin 管道 | `AI agent benchmark CLI`, `LLM eval answer file format`, `CLI stdin JSON input`, `non interactive CLI questionnaire` | `openspec/specs/answer-parser/spec.md`, `openspec/specs/score-command/spec.md`, `README.md` | phase 1 选择 canonical answer file + `score --answers <file>`；stdin 或逐题交互留给后续设计 |
| R07 | 输出契约 | resolved for phase 1 | 定义 human-readable、JSON、Markdown report、exit code 等输出 | `CLI JSON output best practices`, `command line report format`, `CLI exit code conventions`, `machine readable CLI output` | `openspec/specs/result-renderers/spec.md`, `openspec/specs/questions-command/spec.md`, `openspec/specs/score-command/spec.md`, `README.md` | phase 1 text / JSON、stdout/stderr 分离和 error behavior 已冻结；Markdown report 未进入 phase 1 |
| R08 | 测试 fixtures | resolved for phase 1 | 明确题库和评分引擎的测试样例组织方式 | `golden test CLI`, `Go test fixtures JSON`, `table driven tests Go`, `snapshot testing CLI output` | `cmd/mbti-cli/testdata/`, `internal/answers/testdata/`, `internal/questionbank/testdata/`, `internal/result/testdata/`, `internal/scoring/testdata/` | fixed answer sets、schema invalid fixtures、threshold and golden output coverage 已落地 |
| R09 | 伦理表述 | resolved for phase 1 | 避免把 AI 行为风格描述成人格诊断或心理测量结论 | `AI personality assessment disclaimer`, `AI behavior profile wording`, `psychometric test disclaimer`, `responsible AI evaluation wording` | `README.md`, `internal/result/testdata/summary.txt.golden`, result renderer tests | README 和结果文案已使用 behavior style / factual summary；未来报告型输出需要复用该约束 |

## Recommended Order

如果继续推进第二阶段，推荐顺序是：

1. 先做 `R01` 和 `R02`，因为交互式 `assess` 是否存在以及使用什么 prompt/TUI 库必须一起判断。
2. 再做 `R05`，把已实现的隐藏 scoring internals 和 deterministic shuffle 扩展成更完整的 anti-gaming 策略。
3. 只有当第二阶段改变 answer input、result output、report wording 或题库 schema 时，才重开 `R03`、`R04`、`R06`、`R07`、`R08` 或 `R09`。

## Decision Template

每个调研任务结束后，建议按这个格式记录结论：

```markdown
## <Task ID> <Topic>

### Decision

最终选择是什么。

### Why

为什么这样选，排除了哪些方案。

### Impact

会影响哪些命令、数据结构、测试或文档。

### Follow-up

实现时还需要注意什么。
```

## Current Bias

已冻结的第一阶段倾向：

- 优先本地 `questions/` 题库
- 优先非交互 `questions` / `score` 命令
- 优先 answer file，而不是逐题 prompt 或 stdin 管道
- 优先普通 stdout/stderr 和 Cobra，暂不引入复杂 TUI
- 同时支持 human-readable text 和稳定 JSON 输出
- 优先可验证 scoring core，再做交互美化或报告扩展
- 避免把 AI 行为风格包装成人格诊断

后续阶段如果要加入交互式问卷、Markdown report、profile/history 或更强 anti-gaming，需要先产出新的设计和任务板。
