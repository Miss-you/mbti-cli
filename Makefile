MODULE := github.com/Miss-you/mbti-cli
BINARY := bin/mbti-cli
VERSION ?= dev
COMMIT ?= none
DATE ?= unknown
ARGS ?=
LDFLAGS := -s -w \
	-X $(MODULE)/internal/version.Version=$(VERSION) \
	-X $(MODULE)/internal/version.Commit=$(COMMIT) \
	-X $(MODULE)/internal/version.Date=$(DATE)

.PHONY: fmt test lint build run

fmt:
	files="$$(find . -name '*.go' -not -path './.worktrees/*' -not -path './bin/*')"; \
	if [ -n "$$files" ]; then gofmt -w $$files; fi

test:
	go test ./...

lint:
	golangci-lint run ./...

build:
	mkdir -p bin
	go build -trimpath -ldflags "$(LDFLAGS)" -o $(BINARY) .

run:
	go run -ldflags "$(LDFLAGS)" . $(ARGS)
