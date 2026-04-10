# mbti-cli

`mbti-cli` is a small Go CLI for coding-related workflows, built with Cobra and kept intentionally minimal for long-term maintenance.

## 当前命令

- `mbti-cli`
- `mbti-cli help`
- `mbti-cli version`
- `mbti-cli completion`

## 目录结构

- `main.go`: 程序入口
- `cmd/mbti-cli/`: Cobra 命令定义
- `internal/cli/`: 可复用的命令执行封装
- `internal/version/`: 版本信息与构建注入变量
- `docs/plans/`: 设计与实施计划

## 开发

```bash
make fmt
make test
make lint
make build
make run
```

`make run` 会复用与 `make build` 一致的版本注入参数，便于本地做接近发布态的 smoke test。
如果需要传子命令，可使用 `make run ARGS=version`。

## 构建产物

默认构建产物位于 `bin/mbti-cli`。

## License

Apache 2.0. See [LICENSE](LICENSE).
