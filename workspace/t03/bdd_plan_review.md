# T03 BDD Plan Review

## Verdict

基本通过。这个 plan 是 behavior/spec-first 的，方向和 `writing-ginkgo-bdd-tests` 一致：它把 T03 限定在题库 schema validator 的可观察行为上，没有把范围滑向 loader、answer validation、scoring 或 CLI。

唯一需要收紧的是 invalid table 的粒度。当前草案已经比单测更克制，但如果后续写成“把 `validator_test.go` 的所有 mutation 都搬进 Ginkgo”，就会变成 fixture 回归表，而不是代表性 BDD。

## Recommendations

1. 保留一个 canonical v3 通过场景即可，作为 validator 的正向规格入口。
2. invalid 场景按 spec 分组写代表性用例，不要追求穷举：
   - `meta.total` 不一致
   - dimension metadata 缺失或不一致
   - question ID / locale 文本无效
   - option 形状无效
   - threshold 缺失、无序或重叠
3. table 里优先保留“能说明行为边界”的 mutation。每组 1 个代表例通常就够，threshold 组最多保留 2 个不同失败原因即可。
4. 断言只抓验证契约本身：`Validate` 返回 `*ValidationError`，并且 issue 文本包含对应原因；不要断言完整 issue 列表，也不要复刻 unit test 的全部细节。
5. 不要把 `validator_test.go` 的覆盖面当成 BDD 目标。那里可以更细，T03 的 BDD 只需要证明规格成立。

## Non-Goals

- 不覆盖 loader 读写、文件解析或路径错误。
- 不覆盖 answer parser、scoring、classification、rendering。
- 不把所有 `validator_test.go` 的 mutation 逐项搬成 Ginkgo table。
- 不修改 production code 或其他测试文件。
