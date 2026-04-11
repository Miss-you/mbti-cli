# mbti-cli Human Task List

## Purpose

这份文档记录实现 `mbti-cli` 前需要人工阅读、搜索和判断的调研任务。

它面向人阅读，不是 agent task board。目标是帮助后续设计阶段快速冻结关键产品决策，避免实现时反复返工。

## Project Context

`mbti-cli` 的目标是制作一个面向 Coding CLI / AI Agent 的 MBTI-style CLI 工具。

主要题目来源是仓库内的 `questions/` 目录，产品参考见：

- `docs/project-goals.md`
- `docs/ref/mbti.md`

## Research Tasks

| ID | Topic | Goal | Suggested Search Terms | Output | Done When |
| --- | --- | --- | --- | --- | --- |
| R01 | 终端问卷 UX | 明确 CLI 单选题、进度、回退、非交互输入和结果输出的基础体验 | `CLI questionnaire UX`, `terminal survey single choice`, `command line quiz UX`, `interactive CLI best practices` | 1 页 UX 决策摘要 | 能确定 v1 是否支持回退、进度显示、非交互模式 |
| R02 | Go prompt/TUI 库 | 判断是否引入 `survey`、`huh`、`bubbletea` 等库，还是先用 stdin/stdout | `Go survey library`, `charmbracelet huh`, `bubbletea form`, `Go interactive CLI prompt` | 库选型对比表 | 能给出 v1 推荐方案和不选其他库的理由 |
| R03 | 题库 JSON schema | 明确题库版本、多语言、维度、选项、反向题、阈值和兼容策略 | `questionnaire JSON schema`, `survey JSON format`, `quiz schema versioning`, `JSON schema backward compatibility` | schema 决策草案 | 能定义 `questions/*.json` 的 validation rule |
| R04 | MBTI-like scoring | 明确维度分数如何映射到方向、强弱倾向、类型码和解释文本 | `MBTI scoring dimensions`, `personality type scoring thresholds`, `Likert scoring classification`, `type indicator scoring` | scoring 规则说明 | 能把固定答案集映射到稳定结果 |
| R05 | LLM anti-gaming | 降低模型识别测试意图、迎合题目或按显性维度作答的概率 | `LLM benchmark prompt leakage`, `LLM evaluation anti gaming`, `questionnaire reverse scoring`, `survey response bias mitigation` | 题目呈现和乱序策略 | 能确定是否隐藏维度、打乱顺序、使用反向题 |
| R06 | Agent 作答模式 | 判断 Coding CLI / AI Agent 更适合逐题交互、一次性 prompt、answer file 还是 stdin 管道 | `AI agent benchmark CLI`, `LLM eval answer file format`, `CLI stdin JSON input`, `non interactive CLI questionnaire` | input mode 决策 | 能定义 v1 支持的输入路径和以后扩展点 |
| R07 | 输出契约 | 定义 human-readable、JSON、Markdown report、exit code 等输出 | `CLI JSON output best practices`, `command line report format`, `CLI exit code conventions`, `machine readable CLI output` | output contract 草案 | 能冻结默认输出和 `--format json` 是否进入 v1 |
| R08 | 测试 fixtures | 明确题库和评分引擎的测试样例组织方式 | `golden test CLI`, `Go test fixtures JSON`, `table driven tests Go`, `snapshot testing CLI output` | fixture/test plan | 能列出固定答案集、阈值边界和 schema 错误样例 |
| R09 | 伦理表述 | 避免把 AI 行为风格描述成人格诊断或心理测量结论 | `AI personality assessment disclaimer`, `AI behavior profile wording`, `psychometric test disclaimer`, `responsible AI evaluation wording` | wording guideline | 能写出 README 和 CLI 结果页的 disclaimer |

## Recommended Order

1. 先做 `R03` 和 `R04`，因为它们决定 core 数据模型和 scoring。
2. 再做 `R01`、`R06`、`R07`，因为它们决定 CLI command contract。
3. 然后做 `R05` 和 `R09`，用于校正题目呈现和结果措辞。
4. 最后做 `R02` 和 `R08`，决定实现库和测试落地方式。

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

在没有进一步调研前，v1 的默认倾向是：

- 优先本地 `questions/` 题库
- 优先普通 stdin/stdout 和 Cobra，暂不引入复杂 TUI
- 优先 human-readable 输出，保留 JSON 输出扩展点
- 优先可验证 scoring core，再做交互美化
- 避免把 AI 行为风格包装成人格诊断
