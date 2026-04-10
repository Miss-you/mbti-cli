# mbti-cli Automation Skills Design Task Board

## Source Design

- `docs/plans/2026-04-11-mbti-cli-automation-skills-design.md`

## Status Legend

- `todo`: 尚未开始，且依赖未必满足
- `claimed`: 已认领，正在建立上下文
- `research`: 正在调研与收集证据
- `spec`: 正在确定方案与 skill 文案
- `implementing`: 正在落 skill 文件
- `verifying`: 正在做结构和引用校验
- `review`: 正在做内容评审与问题分流
- `blocked`: 被阻塞，需要先解决前置问题
- `done`: 已完成且有验证证据

## Dependency Rules

- 只有 `status=todo` 且所有 `Depends On` 都 `done` 的任务可被认领
- `ready` 不是独立状态，是否可做由依赖推导
- `T02` 到 `T05` 可以并行
- `T06` 依赖 `T03`
- `T07` 依赖 `T04`
- `T08` 依赖 `T02` 到 `T07`

## Task Table

| ID | Title | Goal | Depends On | Parallel | Status | Owner | Claimed At | Workspace | Change | Done When | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| T01 | Freeze migration design | 固化迁移目标、SOP、验收标准与 canonical/alias 映射 | - | no | done | codex | 2026-04-11 | - | - | 设计文档已落盘且可指导后续实现 | source inventory 已核对 |
| T02 | Port planning skill | 迁移并改写 `compatibility-first-planning` | T01 | yes | done | codex | 2026-04-11 | - | - | skill 可用于 mbti-cli 宽泛目标转设计 | 已补足固定设计文档路径与批准后交接要求 |
| T03 | Port task-board skill | 迁移并改写 `deriving-task-board-from-design` | T01 | yes | done | codex | 2026-04-11 | - | - | skill 可从 design 生成稳定 task board | 已固化状态机、claim rule 和 stable ID 约束 |
| T04 | Port delivery skill | 迁移并改写 `delivering-go-task-end-to-end` | T01 | yes | done | codex | 2026-04-11 | - | - | skill 可驱动单任务从认领到关闭 | 已适配 mbti-cli 的 questions/scoring/CLI gate |
| T05 | Port PR monitoring skill | 迁移并改写 `monitoring-pr-ai-reviews` | T01 | yes | done | codex | 2026-04-11 | - | - | skill 可处理本 repo PR 上的 AI review 闭环 | 已移除对不存在脚本/workflow 的强依赖 |
| T06 | Add design-to-task alias | 添加 `breaking-design-into-tasks` 兼容别名 | T03 | no | done | codex | 2026-04-11 | - | - | alias 清晰指向 canonical skill | 保持薄别名，不复制第二套逻辑 |
| T07 | Add delivery alias | 添加 `claiming-and-delivering-work` 兼容别名 | T04 | no | done | codex | 2026-04-11 | - | - | alias 清晰指向 canonical skill | 保持薄别名，不复制第二套逻辑 |
| T08 | Validate and repair migrated skills | 做 repo 内校验并修复问题直到 clean | T02,T03,T04,T05,T06,T07 | no | done | codex | 2026-04-11 | - | - | 结构、文案、引用、术语和闭环全部通过 | 已确认 6 个 skill 存在、6 个 trigger-only description，source leakage clean |

## Claiming Rules

- 先更新本任务板，再开始对应任务实现
- 每个任务默认只对应一个 owner 和一个 workspace
- 如果需要并行，只能在写集不重叠时并行
- alias task 不能脱离 canonical task 单独扩张 scope

## Change Log

- 2026-04-11: 初始化任务板，冻结本次迁移的 8 个稳定任务。
- 2026-04-11: 完成 6 个 skill 迁移与 1 轮修复校验；本次为 bootstrap 迁移，未回填 per-task workspace，后续常规任务按 board 规则创建。
