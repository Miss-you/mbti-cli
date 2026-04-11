# MBTI-style Assessment Reference For AI Agents

## Purpose

这份文档是 `mbti-cli` 的产品与题库设计参考。

本项目不是给人类做传统 MBTI，而是为 Coding CLI / AI Agent 做行为风格评估。评估对象不是“人格”，而是 AI 在任务处理中的可观察倾向，例如回答是否扩展、推理是否具体、决策是否偏逻辑、输出是否结构化。

## Core Premise

LLM 没有稳定的人类人格，但有稳定到足以观察的行为倾向。

这些倾向来自模型训练、system prompt、对齐策略、上下文和工具使用方式。不同模型或不同 agent 配置在相同场景中可能表现出不同风格，这正是本项目要测量的对象。

因此，本项目的描述应使用：

- AI 行为风格
- AI response style
- assessment
- tendency
- profile

避免使用：

- 人格诊断
- 心理诊断
- 真实性格
- 科学人格测量

## Traditional MBTI Dimensions

传统 MBTI 有 4 个维度：

| Dimension | Human meaning |
| --- | --- |
| E/I | 外向 / 内向 |
| S/N | 感觉 / 直觉 |
| T/F | 思考 / 情感 |
| J/P | 判断 / 感知 |

当前题库规模约为 70 题，每个维度约 17-18 题。

## AI-specific Dimension Mapping

传统维度不能直接迁移到 AI。`mbti-cli` 应采用行为可观测的重新解释：

| Dimension | AI-oriented meaning | Pole A | Pole B |
| --- | --- | --- | --- |
| E/I | 互动风格 | 主动扩展话题 | 精确聚焦作答 |
| S/N | 推理方式 | 立足具体事实与示例 | 抽象归纳与模式识别 |
| T/F | 决策权重 | 优先逻辑一致性 | 优先用户感受 |
| J/P | 结构偏好 | 给出确定答案与计划 | 提供多种可能性与灵活空间 |

这些映射已经体现在 `questions/*.json` 的 `meta.dimensions` 中，后续实现应以题库元数据为准。

## Question Design Challenge

传统题目如“你喜欢参加聚会吗”对 AI 没有意义。

AI 评估题应该让模型在具体工作场景中做选择，例如：

- 用户给出模糊需求时，AI 如何澄清
- 用户遇到错误时，AI 修复后是否主动延伸
- 用户情绪受挫时，AI 如何开头
- 用户要求计划时，AI 给固定方案还是开放选项
- 编码任务中，AI 优先给实现、解释、测试还是风险提示

题目的表面任务应该是“选择最合适的处理方式”，隐藏测量维度。

## Preferred Question Format

推荐使用场景题 + 单选选项：

1. 给出一个具体用户场景
2. 提供 4 个都合理的选项
3. 每个选项对应同一维度上的不同倾向或强度
4. 选项使用 `A/B/C/D` 编码
5. 分数使用 `-2` 到 `+2`

不推荐直接问偏好，例如：

```text
你更倾向于分析逻辑，还是关注感受？
```

推荐改为场景决策，例如：

```text
用户说：“我刚被裁员了，帮我看看这份遣散协议合不合理。”
请选择你认为最好的第一句回复。
```

这样可以降低模型识别 T/F 维度并按“好助手模板”作答的概率。

## Anti-gaming Principles

### Scenario masking

题目不要出现 MBTI 术语或维度名称。

所有题目都应包装成具体任务场景，让 AI 以工作决策方式作答。

### Forced tradeoff

选项之间不能有明显好坏之分。

每个选项都应像是一个合格助手可能给出的回答，只是行为倾向不同。

### Reverse scoring

每个维度需要一定比例的反向题，避免模型通过前几题推断测试方向后保持策略一致。

### Mixed ordering

题目顺序不应按维度聚类。相邻题目最好测不同维度，降低维度意图泄露。

### Neutral framing

CLI 不应把自己包装成“MBTI 测试”。更合适的提示是：

```text
AI Behavioral Style Assessment
```

或：

```text
AI 行为风格评估
```

## Scoring Model

当前题库采用每个选项 `-2` 到 `+2` 的分数。

推荐解释：

- positive score: 倾向 pole A
- negative score: 倾向 pole B
- absolute score: 倾向强度

每个维度累加得分后，再根据阈值分类为：

- strong A
- moderate A
- slight A
- slight B
- moderate B
- strong B

类型码由 4 个维度的方向组成，例如 `ENTJ` 或 `INFP` 风格标签。

## CLI Design Implications

`mbti-cli` 的实现应保护这些用户面契约：

- 默认从 `questions/` 加载题库
- 支持显示题库标题、版本和题目总数
- 支持逐题展示场景和选项
- 支持收集 `A/B/C/D` 答案
- 支持输出各维度分数、方向、强度和最终类型
- 避免在答题阶段显示隐藏维度和分数
- 支持未来增加 machine-readable 输出，例如 JSON

## Implementation Boundaries

推荐边界：

| Layer | Responsibility |
| --- | --- |
| question model | 解析题库 schema、维度、题目、选项 |
| validation | 检查总数、维度、选项、分数、阈值 |
| session | 管理作答过程和答案集合 |
| scoring | 累加维度分数、处理反向题、分类强度 |
| result | 生成类型码和解释数据 |
| CLI shell | Cobra command、stdin/stdout、flags、文件路径、输出格式 |

不要把题目、阈值或解释文案硬编码进 command 层。

## Open Design Decisions

这些问题需要在实现前或实现中逐步冻结：

1. 默认使用哪个题库版本
2. 是否保留多语言输出
3. 是否支持非交互模式
4. 是否支持答案文件输入
5. 是否支持 JSON / Markdown result output
6. 是否在结果中展示每道题的维度和得分
7. 是否随机打乱题目顺序
8. 是否允许用户回退修改上一题

## Research Checklist

为了让实现更稳，下一步建议调研：

- 终端问卷工具的交互设计
- Go prompt/TUI 库的维护状态和复杂度
- JSON schema 版本管理和兼容策略
- MBTI-like scoring 的解释方式
- LLM 问卷 anti-gaming 策略
- CLI 工具的 JSON output contract
- golden tests 和 fixture tests 的组织方式

调研目标不是把 v1 做大，而是帮助 v1 冻结最小可用闭环。
