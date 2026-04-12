## ADDED Requirements

### Requirement: Answer validator accepts complete strict answer sets
The system SHALL provide bank-aware answer validation that accepts a complete answer set when every bank question has one answer and each selected option exists on that question.
Answer validation SHALL consume parser-normalized or already-canonical in-memory option codes and SHALL NOT normalize option codes itself.

#### Scenario: Complete answers validate successfully
- **WHEN** a parsed answer set contains one valid option code for every question in the bank
- **THEN** answer validation succeeds
- **AND** validation does not aggregate scores or apply the `reverse` flag

### Requirement: Answer validator rejects unknown question IDs
The system SHALL reject answer sets that contain IDs not present in the question bank.

#### Scenario: Unknown answer ID is rejected
- **WHEN** an answer set contains an answer for `unknown`
- **AND** the question bank does not contain `unknown`
- **THEN** answer validation returns a clear validation error for the unknown question ID

### Requirement: Answer validator rejects invalid option codes
The system SHALL reject answer sets whose selected option code is not defined by the referenced bank question.

#### Scenario: Invalid option code is rejected
- **WHEN** question `q01` defines options `A`, `B`, `C`, and `D`
- **AND** the answer set selects option `Z` for `q01`
- **THEN** answer validation returns a clear validation error for the invalid option code

### Requirement: Answer validator requires all bank questions in strict mode
The system SHALL enforce strict mode by requiring one answer for every question in the bank.

#### Scenario: Missing answer is rejected
- **WHEN** the bank contains question `q02`
- **AND** the answer set does not contain an answer for `q02`
- **THEN** answer validation returns a clear validation error for the missing answer

### Requirement: Answer validator reports deterministic aggregated errors
The system SHALL return all detected strict answer validation issues in deterministic order.

#### Scenario: Multiple answer issues are reported together
- **WHEN** an answer set contains unknown IDs, invalid option codes, and missing answers
- **THEN** answer validation returns one aggregated validation error
- **AND** the error details are stable across repeated validation runs
