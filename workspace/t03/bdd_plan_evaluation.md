# T03 BDD Plan Review Evaluation

## Adopted

1. 保留 canonical v3 通过场景。
   理由：这是 validator 的正向规格入口，能证明 `Validate` 对标准题库合同的接受行为。

2. invalid 场景按 OpenSpec 分组写代表性用例。
   理由：T03 的目标是覆盖规格边界，不是把 `validator_test.go` 的 mutation 全量搬进 Ginkgo。

3. 每组只保留能说明行为边界的少量 mutation。
   理由：单个代表例已经足以证明该类 contract 失效，threshold 组最多保留少量不同失败原因即可。

4. 断言只抓验证契约本身。
   理由：应只检查 `Validate` 返回 `*ValidationError` 且 issue 文本包含对应原因，不扩到完整 issue 列表或其他行为。

5. 不以 `validator_test.go` 的覆盖面作为 BDD 目标。
   理由：BDD 应保持规格表达，单测可以更细，二者职责不同。

## Rejected

None.

## Final BDD Shape

T03 的 BDD 计划保留一个 canonical v3 正向场景，验证 `Validate` 直接接受标准题库合同。invalid-schema 部分按 OpenSpec 场景分组，只写每组 1 个代表性 mutation，threshold 组最多保留 2 个能区分失败原因的例子。所有断言都停留在 validator 契约层，只验证 `*ValidationError` 和对应 issue 文本，不复刻 unit test 的全量 mutation 集合。
