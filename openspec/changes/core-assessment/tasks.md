## 1. Question Bank Model

- [x] 1.1 Add the canonical v3 question bank file needed by the unmarshal contract test
- [x] 1.2 Add a failing `internal/questionbank` unmarshal contract test for representative v3 fields
- [x] 1.3 Implement minimal JSON model types and constants
- [x] 1.4 Run `go test ./internal/questionbank` and keep validation/scoring behavior out of scope

## 2. Question Bank Loader

- [x] 2.1 Add failing loader unit tests for valid, missing, malformed, and empty-path cases
- [x] 2.2 Implement `LoadFile(path)` with typed `Bank` result and source metadata
- [x] 2.3 Run `go test -count=1 ./internal/questionbank` and keep schema validation out of scope

## 3. Answer Model and Parser

- [x] 3.1 Add failing `internal/answers` parser tests for canonical map parsing, option-code normalization, missing/null answers, malformed JSON, and structurally invalid values
- [x] 3.2 Implement minimal answer file model and `Parse(data)` behavior
- [x] 3.3 Run `go test -count=1 ./internal/answers` and keep bank-aware answer validation out of scope

## 4. Question Bank Schema Validator

- [x] 4.1 Add failing validator tests for canonical v3 success and representative invalid bank schemas
- [x] 4.2 Implement `Validate(bank)` with deterministic aggregated validation errors
- [x] 4.3 Run `go test -count=1 ./internal/questionbank` and keep scoring/answer validation out of scope
