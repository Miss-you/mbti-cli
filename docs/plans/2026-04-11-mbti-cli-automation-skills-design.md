# mbti-cli Automation Skills Migration Design

## Goal

将源仓库中的持续开发工作流 skill 迁移到 `mbti-cli`，并改造成适合本项目的自动化开发闭环。

本次迁移后的 skill 集目标是支持这样一条主路径：

1. 从宽泛目标进入设计
2. 从设计生成可认领任务板
3. 认领单个任务并端到端交付
4. 在 PR 上持续处理 AI review

`mbti-cli` 的产品目标是一个面向 Coding CLI / AI Agent 的 Go CLI，用 `questions/*.json` 题库评估行为风格并输出 MBTI 风格结果。

## Non-goals

- 本次不实现新的产品功能，只迁移 skill 与配套文档
- 不为了迁移 skill 额外引入 CI、定时任务或 PR 监控脚本
- 不复制两套同义 workflow；缺失的源 skill 名称用兼容别名解决
- 不为“未来可能的多 provider / 多 runtime”提前抽象

## Current Project Inventory

当前仓库已经具备这些事实：

- Go Cobra CLI 骨架已经存在
- `questions/` 下已有多个版本的题库 JSON
- `docs/plans/` 已用于设计和实施计划
- `.codex/skills/` 当前只有 OpenSpec 相关 skill
- 本仓库当前没有 `.github/workflows/pr-ai-review-monitor.yml`
- 本仓库当前没有 `scripts/pr_ai_review_monitor.py`

这些事实决定了迁移后的 skill 必须优先依赖现有能力，不能引用不存在的本地自动化资产。

## Source Skill Inventory And Mapping

源目录中真实存在的 canonical skill 只有 4 个：

1. `compatibility-first-planning`
2. `deriving-task-board-from-design`
3. `delivering-go-task-end-to-end`
4. `monitoring-pr-ai-reviews`

用户额外要求的两个名称在源仓库中不存在：

1. `breaking-design-into-tasks`
2. `claiming-and-delivering-work`

迁移策略：

- 保留上面 4 个 canonical skill
- 增加 2 个 compatibility alias skill
- alias 只负责触发与转向，不复制第二套复杂流程

对应关系：

| Requested Name | Migration Type | Canonical Behavior |
| --- | --- | --- |
| `compatibility-first-planning` | canonical | 宽泛目标转设计 |
| `deriving-task-board-from-design` | canonical | 从设计生成任务板 |
| `delivering-go-task-end-to-end` | canonical | 单任务端到端交付 |
| `monitoring-pr-ai-reviews` | canonical | PR 上持续处理 AI review |
| `breaking-design-into-tasks` | alias | 指向 `deriving-task-board-from-design` |
| `claiming-and-delivering-work` | alias | 指向 `delivering-go-task-end-to-end` |

## Compatibility Contract For mbti-cli

迁移后的 skill 必须冻结并保护这些 `mbti-cli` 用户面契约：

1. `questions/*.json` 的 schema、版本与加载约定
2. 评分、维度聚合、类型判定与阈值语义
3. CLI 命令、flags、stdout/stderr、exit code 的用户可见行为
4. 题目呈现、作答流程、结果展示等终端交互约束
5. 设计文档、任务板、workspace、OpenSpec 产物之间的一致性

不能把“兼容性”继续理解成 `go-symphony` 的 provider / tracker / dashboard 兼容，而要收敛到 `mbti-cli` 自己的题库、评分和 CLI 契约。

## Terminology Mapping

| Source Concept | mbti-cli Concept |
| --- | --- |
| provider-neutral core | 题库加载、答题会话、评分与结果模型 |
| compatibility shell | Cobra 命令、终端交互、文件输入输出、结果渲染 |
| provider compatibility | 题库 schema、评分语义、CLI 契约稳定性 |
| workflow / tracker parity | 设计文档、任务板、OpenSpec 与代码状态对齐 |
| dashboard / API surface | CLI 命令面、输出格式、可选导出面 |

默认边界：

- core 放业务真相：question model、answer/session model、scoring、result typing
- shell 放交互与适配：Cobra commands、prompt/output、文件读写、JSON 导入导出

只有当当前项目已经证明某个抽象是共性的，才允许上提到 core。

## Unified SOP

每个迁移后的 skill 都必须采用同一套阶段闭环，而不是“写完就算”：

1. 明确本阶段目标、输入与证据
2. 执行最小必要动作
3. 按本 skill 的验收标准审核结果
4. 修复发现的问题
5. 重新审核或复验，直到通过

这个 SOP 必须在 skill 正文里体现为可执行流程，而不是抽象口号。

## Skill Output Standard

每个 skill 的 `SKILL.md` 至少要具备这些属性：

1. `description` 只写触发条件，不摘要 workflow
2. 明确 `When to Use` 和 `Do not use`
3. 明确 required outputs 或阶段性产物
4. 明确流程、检查点、常见错误和下一步
5. 所有路径、命令、依赖都必须在当前仓库可成立
6. 文案围绕 `mbti-cli`，不能残留 `go-symphony` / `symphony` 上下文

## Per-Skill Acceptance Criteria

### `compatibility-first-planning`

必须做到：

- 把宽泛目标收敛成 `mbti-cli` 可实现设计
- 强制输出 goal/non-goals、inventory、compatibility contract、terminology、boundary、phase plan、verification plan、evidence references
- 把 `questions`、scoring 和 CLI contract 明确成兼容面
- 避免为了未来 agent/provider 过度抽象

### `deriving-task-board-from-design`

必须做到：

- 从 `docs/plans/*-design.md` 生成或刷新 `*-design-task.md`
- 保留稳定任务 ID、依赖、状态历史与 `Done When`
- 任务板适合多 agent 并行执行
- 明确 `workspace/<task-id>/` 约定

### `delivering-go-task-end-to-end`

必须做到：

- 认领一个 ready task 并贯穿 research/spec/implementation/verification/review/close
- 保证 task board、workspace、OpenSpec、代码与验证证据一致
- 验证命令只使用当前仓库真实存在的 gate，例如 `make fmt`、`make test`、`make lint`、`make build`
- 对涉及题库和评分的任务，要求有针对性的 contract 验证

### `monitoring-pr-ai-reviews`

必须做到：

- 把 AI review 当作输入而不是命令
- 保护 CLI 契约、题库契约、评分语义和最小 Go 设计
- 不引用不存在的本地 PR 监控脚本或 workflow 作为必需依赖
- 让 GitHub review 处理、验证、回复、重扫形成闭环

### `breaking-design-into-tasks`

必须做到：

- 作为薄别名存在，服务用户词汇习惯
- 清楚说明该 skill 的 canonical path 是 `deriving-task-board-from-design`
- 不复制第二份任务板逻辑

### `claiming-and-delivering-work`

必须做到：

- 作为薄别名存在，服务用户词汇习惯
- 清楚说明该 skill 的 canonical path 是 `delivering-go-task-end-to-end`
- 不复制第二份端到端交付逻辑

## Phase Plan

### Phase 1. Freeze design and mapping

- 写出本设计文档
- 冻结 canonical skill 与 alias skill 的关系

### Phase 2. Migrate canonical skills

- 迁移 `compatibility-first-planning`
- 迁移 `deriving-task-board-from-design`
- 迁移 `delivering-go-task-end-to-end`
- 迁移 `monitoring-pr-ai-reviews`

### Phase 3. Add compatibility aliases

- 添加 `breaking-design-into-tasks`
- 添加 `claiming-and-delivering-work`

### Phase 4. Validate and repair

- 检查 skill 结构、命名、触发描述
- 检查路径与命令是否有效
- 检查是否残留源项目上下文
- 修复问题直到通过

## Verification Plan

迁移完成前，至少执行这些检查：

1. 目标 skill 目录全部存在且都包含 `SKILL.md`
2. 新增 skill 的 frontmatter 合法，且 `description` 只写触发条件
3. skill 正文不残留 `go-symphony`、`symphony` 等源项目语境
4. `monitoring-pr-ai-reviews` 不强依赖本仓库不存在的脚本或 workflow
5. canonical 与 alias 的指向关系清晰，无重复维护逻辑
6. 每个 skill 都包含执行、审核、修复、复验的闭环

## Evidence References

- Source skill: `/Users/lihui/Documents/GitHub/go-symphony/.codex/skills/compatibility-first-planning/SKILL.md`
- Source skill: `/Users/lihui/Documents/GitHub/go-symphony/.codex/skills/deriving-task-board-from-design/SKILL.md`
- Source skill: `/Users/lihui/Documents/GitHub/go-symphony/.codex/skills/delivering-go-task-end-to-end/SKILL.md`
- Source skill: `/Users/lihui/Documents/GitHub/go-symphony/.codex/skills/monitoring-pr-ai-reviews/SKILL.md`
- Local plan: `/Users/lihui/Documents/GitHub/mbti-cli/docs/plans/2026-04-11-mbti-cli-design.md`
- Local implementation plan: `/Users/lihui/Documents/GitHub/mbti-cli/docs/plans/2026-04-11-mbti-cli.md`
- Local project doc: `/Users/lihui/Documents/GitHub/mbti-cli/README.md`
- Local question bank: `/Users/lihui/Documents/GitHub/mbti-cli/questions/mbti-questions.json`
- Local question bank: `/Users/lihui/Documents/GitHub/mbti-cli/questions/mbti-questions-v3.json`
