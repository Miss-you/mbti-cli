# T05 BDD Plan Review

## Verdict

T05 的 BDD 计划整体方向是对的：它围绕 `Validate` 的可观察行为来写，覆盖了 complete pass、unknown ID、invalid option、missing answer 和 deterministic aggregation，也没有把范围扩成 scoring、CLI 或 question-bank schema 校验。

唯一需要收紧的是 `lowercase / non-normalizing` 这个场景。根据 `openspec/specs/answer-validation/spec.md` 和 `internal/answers/validator.go`，T05 的契约是验证器消费已经 parser-normalized 或 canonical 的 in-memory option code，并且验证器本身不负责归一化。这个 case 更适合作为答案解析层的行为证明，而不是 T05 validator BDD 的核心场景。

## Recommendations

1. 保留当前计划中的 strict validation 主线场景：complete answers pass、unknown ID rejected、invalid option rejected、missing answer rejected、multiple issues aggregated in stable order。
2. 删除或下放 `lowercase / non-normalizing` 场景，不要把 T05 写成 parser 归一化的重复证明。
3. 如果需要证明“validator 不做归一化”，用一个很小的直接 `Set{"q01":"z"}` 负例即可，不要扩成 parser 兼容性测试或 fixture 回归测试。
4. 聚合错误的断言保持代表性和稳定性即可，避免把 `validator_test.go` 里的实现细节整体搬进 BDD。

## Non-Goals

- 不复写 answer parsing / normalization 的行为。
- 不扩大到 question-bank schema validation、scoring、rendering 或 CLI 行为。
- 不把 BDD 写成 `validator_test.go` 的 Ginkgo 换皮版。
- 不要求覆盖所有 fixture 变体，重点是规格里明确写出的行为。
