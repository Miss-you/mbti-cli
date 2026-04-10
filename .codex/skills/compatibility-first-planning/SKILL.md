---
name: compatibility-first-planning
description: Use when mbti-cli work starts from a broad goal, rewrite, or multi-surface feature and the question-bank contract, scoring semantics, or CLI boundary are still unclear.
---

# Compatibility-First Planning

## Overview

把一个宽泛的 `mbti-cli` 目标收敛成可实施设计，同时避免“未来可能支持更多 agent/provider/runtime”的想象污染当前 v1。

**Core principle:** 先冻结用户可见契约和项目术语，再讨论抽象。

这个 skill 只负责设计，不负责真正实现。设计批准后，下一步是 `deriving-task-board-from-design`。

## When to Use

在这些场景使用：

- 用户说的是 “做一个新功能 / 重构 / 重写 / 支持更多输出形态”
- 变更会同时影响 `questions/`、scoring、CLI 命令面或输出格式
- 你还不确定哪些行为必须稳定，哪些只是内部实现细节
- 你感觉自己快开始抽象，但还没冻结 v1 边界

这些场景不要使用：

- 小型 bugfix
- 单文件改动且验收标准已经非常明确
- 设计已经批准，当前只差生成任务板或落代码

## Shared Phase Loop

对每个 planning artifact，都按同一循环执行：

1. 先定义这个 section 要解决什么问题
2. 基于仓库证据写出当前版本
3. 审核它是否覆盖契约、边界和验收
4. 修复范围漂移、术语漂移或缺失 gate
5. 重新通读，直到它足以指导实现

不要在关键设计缺口仍然存在时进入实现。

## Required Outputs

设计结果必须显式包含：

1. Goal and non-goals
2. Current project inventory
3. Compatibility contract
4. Terminology mapping
5. Core vs shell boundary
6. Phase plan
7. Rough task breakdown
8. Verification plan
9. Evidence references

缺一项都不算 implementation-ready。

设计必须落盘到 `docs/plans/YYYY-MM-DD-<topic>-design.md`。
下游的任务板 skill 依赖这份文件；只有聊天里的草稿或口头确认不算完成。

## The Flow

### 1. Clarify the target

先把“成功”定义成用户可见行为，而不是包结构。

优先冻结这些面：

- 题库文件与 schema 是否保持兼容
- scoring、维度聚合和阈值是否保持语义稳定
- CLI 命令、flags、stdout/stderr、exit code 是否要保持稳定
- 结果展示、导出格式、交互步骤哪些是 v1 范围

### 2. Inventory the current system

在提架构前，先确认当前仓库事实：

- `questions/*.json` 的真实结构
- `cmd/mbti-cli/` 当前命令面
- `internal/` 里已有的可复用逻辑
- `docs/plans/` 里的已有设计约束
- 当前测试和 Makefile gate 到底证明了什么

结论必须尽量绑定到具体文件，而不是凭聊天记忆。

### 3. Freeze compatibility contracts

把外部契约写死，常见包括：

- question schema / 版本策略
- scoring 语义
- 结果类型和解释规则
- CLI 调用方式、输出和错误语义
- 若存在导出能力，其 payload 结构

不要把兼容性写成 “后面实现时再看”。

### 4. Write terminology mapping

显式写清这些术语：

- 什么是 core truth
- 什么是 CLI shell
- 哪些词是用户词汇，哪些词只是内部词汇
- 哪些词不能为了“听起来通用”而被过度泛化

### 5. Draw the architectural boundary

`mbti-cli` 的默认边界是：

- core: question model、answer/session model、scoring、result typing
- shell: Cobra command、终端交互、文件读写、结果渲染

只抽象当前项目已经证明共通的部分。

### 6. Phase the work around closed loops

第一阶段必须闭合真实用户路径，推荐优先级：

1. 读取一个题库
2. 跑通一轮答题或输入收集
3. 得到可验证的 scoring 结果
4. 在 CLI 输出结果
5. 正常退出

不要把第一阶段写成 “先做通用框架”。

### 7. Define checks before implementation

每个设计至少要回答：

- `Scope`: v1 到底覆盖哪条完整路径
- `Contract`: 哪些用户面必须稳定
- `Boundary`: 哪些抽象是现在必须的
- `Operational`: 是否能跑通一个真实闭环
- `Data`: question/scoring contract 如何被证明
- `Maintenance`: 维护者能否快速理解

### 8. Hand off

只有当设计已经写入 `docs/plans/YYYY-MM-DD-<topic>-design.md`，并且被明确批准后，才进入 `deriving-task-board-from-design` 生成执行任务板。

## Quick Reference

| Artifact | Question it answers |
| --- | --- |
| Goal and non-goals | 这次到底做什么，不做什么？ |
| Inventory | 当前仓库已经有什么？ |
| Compatibility contract | 哪些用户行为不能被偷偷改掉？ |
| Terminology mapping | 哪些词必须稳定，哪些词只是内部实现？ |
| Core vs shell | 业务真相和 CLI 适配边界怎么分？ |
| Phase plan | 先闭合哪条真实路径？ |
| Verification plan | 用什么证据证明设计成立？ |

## Default Checks

- `Scope`: 是否只覆盖一个清晰的用户路径
- `Questions`: 是否冻结了 question schema / 版本策略
- `Scoring`: 是否冻结了分数与类型判定语义
- `CLI`: 是否覆盖命令、flags、输出、错误语义
- `Boundary`: 是否把交互壳和业务真相分开
- `Maintenance`: 是否避免为了未来场景过度抽象

## Red Flags

- “先做通用 agent/provider 抽象，以后方便扩展”
- “scoring 规则后面再补”
- “questions schema 很简单，不用先冻结”
- “先按感觉拆 package，契约边走边看”
- “CLI 输出只是皮肤，不算兼容面”

## Common Mistakes

- 从宽泛目标直接跳到包结构
- 把 question/scoring 语义当成实现细节
- 为未来运行时、插件、provider 提前抽象
- 没有证据引用，只写主观判断
- 写完设计却没有清楚下一步任务拆解入口

## Output Standard

好的规划结果应该让读者立刻知道：

- 当前 v1 目标是什么
- 哪些内容明确不做
- `questions`、scoring 和 CLI 哪些面必须稳定
- core 和 shell 怎么分
- 第一条闭环要怎么交付
- 设计文档最终落在哪个 `docs/plans/*-design.md` 路径
- 下一步如何拆成可认领任务

## Next Step

设计批准后，使用 `deriving-task-board-from-design` 创建或刷新对应的 `docs/plans/*-design-task.md`。
