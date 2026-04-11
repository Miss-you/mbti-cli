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

## 5. Answer Validation

- [x] 5.1 Add failing `internal/answers` validation tests for valid strict answers, unknown IDs, invalid options, missing answers, and aggregated deterministic errors
- [x] 5.2 Implement strict bank-aware `Validate(bank, answers)` behavior with deterministic aggregated validation errors
- [x] 5.3 Run `go test -count=1 ./internal/answers` and keep scoring/CLI behavior out of scope

## 6. Scoring Engine

- [x] 6.1 Add failing `internal/scoring` tests for deterministic dimension totals, reverse metadata not being reapplied, and strict invalid answers failing before aggregation
- [x] 6.2 Implement minimal `Score(bank, answers)` behavior with signed option score aggregation
- [x] 6.3 Run `go test -count=1 ./internal/scoring` and keep threshold classification, result rendering, and CLI behavior out of scope

## 7. Threshold and Type Classifier

- [x] 7.1 Add failing `internal/scoring` classifier tests for threshold boundaries, zero/balanced behavior, fixed type order, and classifier error cases
- [x] 7.2 Implement `Classify(bank, result)` with threshold strength labels, pole metadata, dimension letters, balanced zero handling, and deterministic type generation
- [x] 7.3 Run `go test -count=1 ./internal/scoring` and keep rendering and CLI behavior out of scope
