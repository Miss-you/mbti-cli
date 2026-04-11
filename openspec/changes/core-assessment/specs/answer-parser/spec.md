## ADDED Requirements

### Requirement: Canonical answer file parser returns normalized answer set
The system SHALL provide an answer-file parser for the first-phase canonical map shape, returning a `questionID -> optionCode` answer set with option codes normalized to uppercase.

#### Scenario: Canonical answer map parses into normalized answers
- **WHEN** an answer file JSON object contains an `answers` map such as `{"q01":"a","q02":" C "}`
- **THEN** the parser returns answers keyed by `q01` and `q02`
- **AND** the returned option codes are `A` and `C`

### Requirement: Answer parser rejects structurally invalid answer files
The system SHALL reject answer-file inputs that cannot be parsed as the canonical answer map shape before bank-aware validation runs.

#### Scenario: Missing answers object is rejected
- **WHEN** the input JSON object does not contain an `answers` object
- **THEN** the parser returns a clear answer-file structure error

#### Scenario: Null answers object is rejected
- **WHEN** the input JSON object contains `"answers": null`
- **THEN** the parser returns a clear answer-file structure error

#### Scenario: Malformed answer JSON is rejected
- **WHEN** the input is not valid JSON
- **THEN** the parser returns a clear answer-file parse error

#### Scenario: Non-string answer values are rejected
- **WHEN** the `answers` object contains a non-string answer value
- **THEN** the parser returns a clear answer-file parse error

### Requirement: Answer parser does not perform bank-aware validation
The system SHALL keep T04 answer parsing independent from question-bank validation, option existence checks, completeness checks, scoring, rendering, and CLI wiring.

#### Scenario: Bank-aware validation is deferred
- **WHEN** the parser receives syntactically valid answer map entries
- **THEN** it returns the normalized answer set without checking whether question IDs exist in a bank
- **AND** it does not check whether an option code is valid for a specific question
- **AND** it does not require all bank questions to be answered
