# mbti-cli Project Goals

## Project Goal

`mbti-cli` 的目标是制作一个面向 Coding CLI 和 AI Agent 的 MBTI-style CLI 工具。

它不是给人类做传统人格测试，而是通过一组场景题观察 AI Agent 在编码、解释、决策、协作、边界控制等任务中的行为倾向，并输出一个便于理解和比较的风格类型。

## Primary Question Source

本项目的主要题目来源是仓库内的 `questions/` 目录。

当前已存在的题库文件包括：

- `questions/mbti-questions.json`
- `questions/mbti-questions-v2.json`
- `questions/mbti-questions-v3.json`

后续实现应优先把 `questions/` 视为本地 canonical data source，而不是在代码中硬编码题目。

题库相关实现需要保护这些契约：

- question schema
- 题库版本
- 题目总数
- 维度分布
- 选项编码
- 分数方向
- 反向计分语义
- 分类阈值

## Product Positioning

`mbti-cli` 应该是 coding-cli friendly 的工具，而不是网页问卷或泛心理测试工具。

核心使用场景：

- 在终端里运行 AI 行为风格评估
- 让 Coding CLI / AI Agent 直接回答题目
- 快速得到维度分数、类型和解释
- 支持未来把结果用于 agent profile、prompt tuning 或团队内 agent 比较

## User-Facing Principles

- CLI 输出应简洁、可读、可复制
- 默认流程应能一条命令跑通
- 题目展示要避免暴露被测维度
- 结果解释要明确这是“行为风格评估”，不是传统人格诊断
- 错误信息应直接说明题库、输入或评分哪里不合法
- 默认数据源应清晰，必要时允许指定题库文件

## Implementation Principles

- Go 代码应把核心业务逻辑与 Cobra command shell 分开
- 题库加载、schema validation、session、scoring、result typing 属于 core
- 命令参数、终端交互、输出格式、文件路径处理属于 CLI shell
- 不要为了未来 provider、插件或服务端形态提前抽象
- 优先实现一个可验证的本地闭环：load questions -> collect answers -> score -> classify -> render result

## Reference Materials

主要产品和题库设计参考见：

- `docs/ref/mbti.md`

该文件记录了 AI 版 MBTI 的维度重解释、题目设计原则、防识别策略和 CLI 实现启发。

## Research Needed Next

为了更稳地实现这个目标，下一步建议调研这些问题：

1. CLI assessment UX

   搜索和对比终端问卷工具的交互方式，包括单选题、进度提示、可返回修改、非交互模式、stdin/JSON 输入和结果输出格式。

2. Go TUI / prompt libraries

   调研 `survey`、`huh`、`bubbletea`、`cobra` 组合的适用边界。重点判断本项目是否需要交互式库，还是先用普通 stdin/stdout 即可。

3. Question schema design

   调研问卷题库 JSON schema 的常见做法，包括版本、locale、多语言文本、选项、分数、反向题、维度元数据、阈值和向后兼容策略。

4. Scoring and classification

   调研 MBTI-like 工具如何从维度分数映射到类型、强弱倾向和解释文本。重点不是照搬心理测量，而是设计稳定、可解释的工程规则。

5. Anti-gaming and prompt leakage

   调研如何降低 LLM 识别测试意图，包括场景包装、题目乱序、反向题、维度隐藏和避免心理学术语泄露。

6. Agent benchmark input modes

   调研 Coding CLI / AI Agent 更适合怎样作答：逐题交互、一次性 prompt、JSON answer file、命令参数或管道输入。

7. Output contract

   调研适合 CLI 工具的结果输出，包括 human-readable summary、JSON output、Markdown report、exit code 和后续自动化消费方式。

8. Validation and fixtures

   调研如何为题库和评分引擎建立测试 fixture，包括固定答案集、维度边界值、阈值边界、题库 schema 错误样例和 golden output。

9. Ethical framing

   调研如何表述“AI 行为风格评估”才不会暗示 AI 具有人类人格，也不会把 MBTI 当作科学诊断。

## Initial Non-goals

- 不做网页版本
- 不做真实心理测量学诊断
- 不先做云服务或 API server
- 不在 v1 支持复杂插件系统
- 不把题目硬编码进 Go 源码
