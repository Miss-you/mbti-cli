# Ginkgo Answers Parse Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Add a small Ginkgo/Gomega practice suite for the answer parser behavior.

**Architecture:** Keep the production parser unchanged and add BDD-style characterization tests beside the existing standard Go tests. The suite focuses on `internal/answers.Parse` because it has clear input/output behavior and minimal setup.

**Tech Stack:** Go 1.24, Ginkgo v2, Gomega, existing `go test ./...` workflow.

---

### Task 1: Add Ginkgo Suite Bootstrap

**Files:**
- Create: `internal/answers/parser_suite_test.go`

**Step 1: Write the failing suite/test scaffold**

Create a Ginkgo suite entrypoint with `RegisterFailHandler(Fail)` and `RunSpecs(t, "Answers Suite")`.

**Step 2: Run test to verify it fails before dependencies are added**

Run: `go test ./internal/answers`

Expected: FAIL with missing `github.com/onsi/ginkgo/v2` or `github.com/onsi/gomega`.

**Step 3: Add test dependencies**

Run: `go get github.com/onsi/ginkgo/v2 github.com/onsi/gomega`

**Step 4: Run test to verify the suite compiles**

Run: `go test ./internal/answers`

Expected: PASS.

### Task 2: Add Parse Behavior Specs

**Files:**
- Create: `internal/answers/parser_ginkgo_test.go`

**Step 1: Write specs for valid answers**

Cover normalizing lowercase and whitespace-padded option codes.

**Step 2: Run package tests**

Run: `go test ./internal/answers`

Expected: PASS.

**Step 3: Write specs for invalid parser inputs**

Cover missing/null `answers`, malformed JSON, structurally invalid values, and deferring bank-aware validation.

**Step 4: Run focused and full tests**

Run: `go test ./internal/answers`

Run: `go test ./...`

Expected: PASS.
