## ADDED Requirements

### Requirement: Question bank validator accepts canonical v3 schema
The system SHALL provide a schema validator for the typed question bank model that accepts the canonical v3 bank without applying scoring behavior.

#### Scenario: Canonical v3 bank validates
- **WHEN** `questions/mbti-questions-v3.json` is loaded into the typed question bank model
- **THEN** schema validation succeeds
- **AND** validation does not apply the `reverse` flag or aggregate scores

### Requirement: Question bank validator rejects invalid schema invariants
The system SHALL reject question banks whose metadata, dimensions, thresholds, questions, or options violate the first-phase question-bank contract.

#### Scenario: Metadata total does not match question count
- **WHEN** `meta.total` differs from the number of questions
- **THEN** schema validation returns a clear validation error

#### Scenario: Dimension metadata is incomplete or inconsistent
- **WHEN** supported dimensions are missing, unknown dimensions are present, or dimension counts do not match actual questions
- **THEN** schema validation returns a clear validation error

#### Scenario: Question identity or localized text is invalid
- **WHEN** a question has an empty ID, duplicate ID, unsupported dimension, or missing localized scenario text
- **THEN** schema validation returns a clear validation error

#### Scenario: Option shape is invalid
- **WHEN** a question does not define exactly one `A`, `B`, `C`, and `D` option with localized labels and an allowed score value
- **THEN** schema validation returns a clear validation error

#### Scenario: Threshold buckets are incomplete or malformed
- **WHEN** scoring thresholds omit a supported strength, include an unknown strength, define an unordered range, or overlap another supported threshold range
- **THEN** schema validation returns a clear validation error
