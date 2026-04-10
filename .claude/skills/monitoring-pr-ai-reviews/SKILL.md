---
name: monitoring-pr-ai-reviews
description: Use when an mbti-cli task is already implemented, a GitHub PR exists or must be opened, and follow-up work is still needed because Copilot or other AI review comments may arrive after the initial push.
---

# Monitoring PR AI Reviews

## Overview

在实现完成后把闭环补完整：开 PR、持续观察 AI review、用兼容性和正确性判断每条建议，而不是盲目同意。

**Core principle:** AI review comments are inputs to evaluate, not orders to follow.

**REQUIRED SUB-SKILLS:** `github:gh-address-comments`, `receiving-code-review`, `verification-before-completion`

## When to Use

在这些场景使用：

- 任务代码已经实现并验证
- PR 已存在，或现在就要开 PR
- Copilot 或其他 AI review 可能在首次 push 之后继续出现
- 剩余工作是 comment triage、修复、回复、重扫和 thread 收口

这些场景不要使用：

- 初始设计阶段
- 原始任务尚未实现
- 人类 review 已经改变产品方向，需要回到设计重新决策

## Shared Review Loop

对每条 review thread，都按同一循环执行：

1. 先定义这条评论的真实主张是什么
2. 对照代码、契约和上下文核实它
3. 判断它是否成立
4. 修复问题或写出拒绝理由
5. 跑 fresh verification 后再回复或收口

不要凭直觉关闭 thread。

## The Flow

1. 确认任务真的 ready for PR
   - 跑 fresh verification
   - 确认任务板、workspace、验证证据一致
2. 打开或刷新 PR
   - 优先使用当前可用的 GitHub 能力
   - 不要把不存在的本地脚本当必需前置
3. 读取 review thread，而不是只看平铺 comments
   - 保留 thread、文件锚点、outdated 状态和 resolution 状态
4. 按 repo 原则评估每条 AI 建议
   - 保护 CLI 契约
   - 保护 `questions/*.json` 契约
   - 保护 scoring 与结果语义
   - 优先最小修复，不扩 scope
5. 只实现 justified changes
   - 先跑最窄证明，再跑更宽 gate
6. 刷新 PR
   - push
   - 回复修复或拒绝理由
   - 只有 fix/rationale 已在 PR 可见时，才关闭 thread
7. 重扫直到没有 unresolved AI review，或只剩明确由用户持有的决策

## Evaluation Rules

接受建议，当它指出的是：

- 真实 bug、回归、race、错误处理缺失
- CLI 命令、flags、stdout/stderr、exit code 的破坏
- `questions` schema 或 scoring 语义漂移
- 缺少证明关键行为的验证

拒绝建议，当它主要要求：

- 为了测试单独扩大 public API
- 为未来 provider/runtime 提前抽象
- 在任务已接近完成时扩大重构范围
- 引入当前 repo 并不需要的框架、插件层或复杂 extension point

## Verification

在回应 AI review 前后，优先使用当前仓库真实存在的命令：

- `make fmt`
- `make test`
- `make lint`
- `make build`

如果 review 触达 CLI 行为、题库加载或 scoring 逻辑，要补相关 focused test / smoke test。

## Optional Automation

如果未来仓库新增 PR review monitor workflow 或本地脚本，可以把它们接入这个 skill。

在那之前，这个 skill 的默认实现必须能在**没有额外脚本**的情况下完成闭环。

## Common Mistakes

- 把 Copilot 建议当成必须执行
- 在 fix 或 rationale 还没出现在 PR 前就关闭 thread
- 回归了 CLI / scoring 契约却只跑局部测试
- 为满足评论而扩大 public API
- 在本仓库没有相关脚本时仍硬引用死路径

## Next Step

当所有 justified AI review 都已处理、相关验证已重跑、thread 已正确收口后，再结束 PR 跟进周期或进入合并流程。
