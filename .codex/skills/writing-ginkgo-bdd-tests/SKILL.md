---
name: writing-ginkgo-bdd-tests
description: Use when mbti-cli has task and spec artifacts and needs Go BDD tests with Ginkgo in addition to existing TDD or standard Go tests.
---

# Writing Ginkgo BDD Tests

## Overview

把 task 和 spec 变成可执行的 Ginkgo 行为规格，而不是把现有 `testing`/`testify` 用例机械翻译一遍。

**核心原则：** BDD 用例必须追踪用户可观察行为或规格场景；`Describe/When/It` 要能读成一句有意义的需求句子。

**覆盖原则：** BDD 可以和 TDD 覆盖同一行为。长期希望迁移到 BDD 风格时，应从 task/spec 重新表达核心行为，而不是只补 TDD 没测到的边角。

**BDD 证据原则：** 如果没有看过规格失败，就不知道它是否真的能保护行为。

**REQUIRED SUB-SKILL:** 如果写 BDD 时需要改生产代码，使用 `test-driven-development`，先让 Ginkgo spec 失败，再写最小实现。

## When to Use

使用在这些场景：

- 用户给出或指向 `docs/plans/*-design-task.md`、OpenSpec `spec.md`、`tasks.md`、proposal/design 等 artifact
- 已有 TDD/标准 Go 测试，但还需要一份 Ginkgo/Gomega BDD 用例
- 需要把 task 验收条件、OpenSpec `Scenario`、CLI 契约或领域规则写成可读规格
- Go 包需要新增或扩展 `*_ginkgo_test.go`、`*_suite_test.go`

不要使用在这些场景：

- 只是解释 Ginkgo/testify 区别，没有要求落测试
- 纯前端、非 Go 测试
- 缺少 task/spec 且无法从仓库推断验收行为；如果用户只给任务编号，先从 `docs/plans`、`openspec`、`workspace` 中定位 artifact

## Inputs and Outputs

输入应明确到文件或任务；如果不明确，先在仓库 artifact 中定位：

| 输入 | 用途 |
| --- | --- |
| task 列表 | 确定要覆盖的验收项和完成边界 |
| spec/proposal/design | 确定真实行为、错误分支、兼容约束 |
| 现有 TDD 测试 | 理解已有行为证据和测试 helpers；允许 BDD 重复覆盖核心行为 |
| 生产代码入口 | 确定应从包 API、CLI helper、fixture 还是 renderer 测 |

输出通常包括：

- 新增或更新的 `*_ginkgo_test.go`
- 必要时新增 package-level `*_suite_test.go`
- 必要时加入 `github.com/onsi/ginkgo/v2` 和 `github.com/onsi/gomega`
- 更新 task checklist 或验证记录，说明哪些 spec 场景已有 BDD 覆盖

## Workflow

1. 读取 task 和 spec 文件；列出必须被 BDD 覆盖的行为句子。
2. 读取相关生产代码和现有 `*_test.go`；理解已有测试意图和 helper，但用 task/spec 重新选择 BDD 行为覆盖面。
3. 为每个行为选择测试层级：纯包 API、CLI command helper、fixture/golden、或端到端命令。
4. 设计 Ginkgo 结构：`Describe` 放被测能力，`When`/`Context` 放 spec 条件，`It` 放用户可观察结果。
5. 先写一个最小 Ginkgo spec，运行目标包测试看 RED。
6. 验证 RED 是有意义的：
   - 新行为缺失：失败应指向缺失行为。
   - 已实现行为补 BDD：对每个行为组做 negative control；高风险核心场景逐条验证；恢复后再继续。
   - suite/dependency failure 只能证明测试 harness，不能替代行为断言失败。
7. 跑到 GREEN；如果生产行为缺失，遵循 `test-driven-development` 修生产代码。
8. GREEN 后才重构测试结构、抽 helper、整理 fixtures；重构期间保持测试全绿。
9. 继续补齐 task/spec 场景；优先覆盖正常路径、错误路径、边界、兼容契约。
10. 跑 focused Ginkgo verbose 和全仓测试：
   - `go test -v ./path/to/pkg -run Test<Name> -count=1 -ginkgo.v`
   - `go test ./...`

## BDD Red-Green-Refactor

### RED - Write a Failing Spec

每次只写一个最小行为规格。规格标题要清楚，断言真实代码，mock 只在不可避免时使用。

```go
It("rejects incomplete answer sets before scoring", func() {
    result, err := Score(bank, answers.Set{"q01": "A"})

    Expect(err).To(MatchError(ContainSubstring("missing answer for question q04")))
    Expect(result).To(BeZero())
})
```

运行目标包测试，确认失败原因正确。测试直接通过时，不要继续堆场景；先证明断言能抓住错误。

### GREEN - Make the Spec Pass

如果生产行为已经存在，恢复 negative control，让 spec 回到真实期望并跑绿。  
如果行为缺失，只写让当前 spec 通过的最小生产代码，不顺手实现额外 task。

### REFACTOR - Clean Up While Green

只在 GREEN 后清理：

- 抽 `BeforeEach` 中真正共享的 setup
- 抽 fixture helper
- 合并同一行为的多个输入为 `DescribeTable`
- 改善 `Describe/When/It` 命名

不要在 refactor 阶段新增行为。

## Scenario Shape

每个 `It` 的主体应能看出 Given / When / Then：

```go
It("rejects incomplete answer sets before scoring", func() {
    bank := scoringBank()                         // Given
    answerSet := answers.Set{"q01": "A"}          // Given

    result, err := Score(bank, answerSet)         // When

    Expect(err).To(MatchError(ContainSubstring("missing answer for question q04"))) // Then
    Expect(result).To(BeZero())                   // Then
})
```

不要求固定写注释；要求 setup、动作、断言三段在阅读上清楚。

## Scenario Priority

优先选择这些 BDD 场景：

1. 用户最关心的主流程
2. spec 明确写出的 `Scenario`
3. task 验收条件和兼容契约
4. 错误路径、边界值、历史 bug
5. 低价值实现细节不写 BDD

## BDD Mapping

| Artifact 内容 | Ginkgo 表达 |
| --- | --- |
| capability / command / package behavior | `Describe("Score command", ...)` |
| OpenSpec `Scenario` 或 task 验收条件 | `When("answers are incomplete", ...)` |
| 可观察结果 | `It("prints validation errors without stdout", ...)` |
| 同一行为的多个输入例子 | `DescribeTable` + `Entry` |
| 共享 fixture/setup | `BeforeEach`，保持短小，避免隐藏主行为 |
| 清理临时文件、mock controller | `DeferCleanup` |

好的 `It` 标题描述行为，不描述实现：

```go
It("rejects incomplete answer sets before scoring", func() {
    result, err := Score(bank, answerSet)

    Expect(err).To(MatchError(ContainSubstring("missing answer for question q04")))
    Expect(result).To(BeZero())
})
```

## Ginkgo Suite Pattern

如果目标 package 还没有 suite 文件，新增一个：

```go
package answers

import (
    "testing"

    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
)

func TestAnswers(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "Answers Suite")
}
```

命名规则：

- suite 文件：`<package>_suite_test.go` 或已有项目约定名称
- specs 文件：`<feature>_ginkgo_test.go`
- suite 名称：`Answers Suite`、`Score Command Suite`，能定位 package 或 capability

## Test Design Rules

- 不要逐行翻译 testify 测试；从 task/spec 重建 BDD 场景。
- 可以重复覆盖 TDD 已测行为；重复的理由应是“这是 task/spec 的核心行为”，不是“把旧测试换个语法”。
- 不要让 `Describe/When` 只是文件夹；它们必须说明行为上下文。
- 每个 `It` 只验证一个用户可观察结果或一个紧密结果组。
- 一个 `It` 标题里出现多个无关的 “and” 时，优先拆成多个 specs。
- BDD spec 通过得太快时，对行为组做 negative control，证明它不是空保护。
- 嵌套超过三层时，优先拆 `Describe`、拆文件、或把条件下沉到 helper/fixture 名称。
- 复用现有 fixtures/golden，但不要为了方便修改 canonical fixture。
- mock 不是 Ginkgo 内置能力；优先手写 fake，只有需要验证调用次数、参数或顺序时才用 gomock/testify mock。
- mock 不能成为被测主体；断言用户可观察行为，少量断言调用参数只用于证明协作边界。
- 不要为写 BDD 测试改生产 API；如果现有代码不可测，先说明耦合点，再用 TDD 做最小重构。
- 如果 BDD 暴露 spec 和实现不一致，停止扩大测试面，先报告不一致并修正 artifact 或实现。

## Quick Checklist

- [ ] task/spec 中每个关键 `Scenario` 或验收项都有对应 BDD 场景
- [ ] Ginkgo 输出能读成业务句子
- [ ] 已避免机械复制现有 TDD 测试，即使覆盖行为有意重复
- [ ] 新增行为 spec 看过 meaningful RED；已实现行为组看过 negative control failure
- [ ] RED 失败原因符合预期，不是 typo、import、fixture 路径错误
- [ ] 如果改了生产代码，已遵循 `test-driven-development`
- [ ] package 测试通过
- [ ] `go test ./...` 通过
- [ ] task checklist 或验证记录说明 BDD 覆盖范围

## Rationalizations to Reject

| Excuse | Reality |
| --- | --- |
| “现有 TDD 已经过了，BDD 直接补上就行” | 直接通过的补测不证明能抓回归；至少按行为组做 negative control。 |
| “TDD 已经覆盖，所以 BDD 不需要写” | 如果这是 task/spec 核心行为，BDD 可以重复覆盖，作为长期规格入口。 |
| “只是换成 Ginkgo 写法” | BDD 不是换皮；必须从 task/spec 写行为句子。 |
| “先把所有 specs 写完再跑” | 一次只写一个行为，先看失败，再继续。 |
| “suite/dependency 失败已经算 RED” | 那只证明 harness；行为 spec 仍要证明断言有效。 |
| “mock 调用都对，所以行为对” | mock 行为不是产品行为；优先断言真实输出、错误、状态或文件内容。 |
| “为了测试方便改个 API” | 测试困难说明设计耦合；先说明问题，用 TDD 做最小重构。 |

## Red Flags

- 先改生产代码，再补 Ginkgo
- `It("works")`、`It("returns error")`
- 大量复制现有 testify case，但没有 task/spec 映射
- 因为 TDD 已覆盖就跳过 task/spec 核心 BDD 场景
- 新 spec 第一次运行就全绿，且没有 negative control
- 不能解释某条 spec 对应哪个 task/spec 场景
- mock 断言多于行为断言
- 只跑全仓测试，没有看 verbose Ginkgo 规格输出

遇到这些信号，停止扩大测试面，回到 task/spec 重新写一个最小可失败的行为规格。

## Common Mistakes

| Mistake | Fix |
| --- | --- |
| 把 Ginkgo 当 testify 换皮 | 从 task/spec 写行为句子，再落代码 |
| `It("works")`、`It("returns error")` | 写清条件下的业务结果 |
| 一个 `BeforeEach` 做完所有世界构造 | 只放共享、稳定、必要 setup |
| 所有 case 都用 `DescribeTable` | 只有同一行为多组数据才用 table |
| 为 mock 而 mock | 能用真实代码或手写 fake 就不用 mock |
| 只跑 `go test ./...` | 还要跑 verbose Ginkgo 看规格可读性 |
