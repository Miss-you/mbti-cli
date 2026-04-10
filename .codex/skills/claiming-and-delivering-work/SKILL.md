---
name: claiming-and-delivering-work
description: Use when an mbti-cli design-task board already exists, one task is ready to claim, and full task-delivery coordination is needed across repo artifacts and review gates.
---

# Claiming And Delivering Work

## Overview

这是一个兼容命名 alias，用来承接“认领并交付任务”这一类表述。

它的 canonical workflow 是 `delivering-go-task-end-to-end`。不要在这里再维护第二套端到端交付流程。

## When to Use

在这些场景使用：

- 用户明确说的是 “认领任务并交付 / claim and deliver”
- 已有任务板，且你准备推进一个具体 ready task
- 你想沿用口语化入口词，而不是直接点 canonical skill

这些场景不要使用：

- 设计或任务板仍未就绪
- 当前工作还停留在宽泛规划阶段
- 你已经确定应该直接使用 `delivering-go-task-end-to-end`

## What To Do

1. 确认已存在 `docs/plans/*-design-task.md`
2. 选择一个所有硬依赖都已满足的 ready task
3. 直接转到 `delivering-go-task-end-to-end`
4. 按该 skill 完成认领、spec、实现、验证、review 和关闭

## Shared Alias Loop

1. 先确认目标真的是“认领并交付单个任务”
2. 调用 canonical skill
3. 审核任务板、workspace、代码和验证是否同步
4. 如果发现漂移，就回到 canonical flow 修复
5. 只有所有阶段证据一致时，任务才允许关闭

## Guardrails

- 不要跳过任务板更新直接开工
- 不要把 alias 扩展成独立 workflow
- 保持 workspace、OpenSpec、代码和验证证据一致

## Next Step

使用 `delivering-go-task-end-to-end`。
