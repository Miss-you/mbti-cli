---
name: breaking-design-into-tasks
description: Use when an approved mbti-cli design exists but the paired design-task board is missing, stale, or too coarse for safe claiming and parallel delivery.
---

# Breaking Design Into Tasks

## Overview

这是一个兼容命名 alias，用来承接用户更口语化的说法。

它的 canonical workflow 是 `deriving-task-board-from-design`。不要在这里再维护第二套任务板流程。

## When to Use

在这些场景使用：

- 用户明确说的是 “拆任务 / break down the design / 先出 task board”
- 已有批准设计，但还没有可执行任务板
- 你想沿用口语化入口词，而不是直接点 canonical skill

这些场景不要使用：

- 设计还未冻结
- 你已经在交付某个具体任务
- 你已经确定应该直接使用 `deriving-task-board-from-design`

## What To Do

1. 确认存在已批准的 `docs/plans/*-design.md`
2. 直接转到 `deriving-task-board-from-design`
3. 创建或刷新对应的 `*-design-task.md`
4. 之后再进入任务认领和交付

## Shared Alias Loop

1. 先确认你要解决的是“设计拆任务”而不是“开始交付”
2. 调用 canonical skill
3. 审核输出任务板是否具备 stable ID、依赖和 `Done When`
4. 如果缺失，就回到 canonical skill 修正
5. 只有任务板可安全认领时，才继续下一步

## Guardrails

- 不要发明第二套表结构
- 不要复制另一份状态机
- 保持 stable ID、claim-before-work、workspace 约定

## Next Step

使用 `deriving-task-board-from-design`。
