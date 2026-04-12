## ADDED Requirements

### Requirement: Result renderer builds stable summary DTOs
The system SHALL provide an in-memory result summary DTO built from question bank metadata, raw scoring counts, and scoring classification output.

#### Scenario: Summary preserves scoring and metadata
- **WHEN** a result summary is built from a question bank, raw scoring result, and scoring classification
- **THEN** the summary includes the bank title and version
- **AND** it includes answered and total counts from the raw scoring result
- **AND** it includes the classifier type string
- **AND** it includes dimension rows in fixed `EI`, `SN`, `TF`, `JP` order

#### Scenario: Missing classification is rejected
- **WHEN** a required dimension classification is missing
- **THEN** summary construction returns a contextual error for that missing dimension

### Requirement: JSON result renderer emits stable machine output
The system SHALL render result summaries as stable JSON suitable for golden tests and automated consumers.

#### Scenario: JSON output uses stable fields and dimension order
- **WHEN** a summary is rendered as JSON
- **THEN** the output is indented
- **AND** it is newline-terminated
- **AND** top-level fields are ordered as `meta`, `type`, and `dimensions`
- **AND** dimensions are ordered as `EI`, `SN`, `TF`, and `JP`

#### Scenario: Balanced dimensions are explicit in JSON
- **WHEN** a dimension classification is balanced
- **THEN** JSON renders its letter as `X`
- **AND** its strength as `balanced`
- **AND** its pole as `balanced`
- **AND** its balanced flag as `true`

### Requirement: Text result renderer emits factual readable output
The system SHALL render result summaries as readable text without implying a personality diagnosis.

#### Scenario: Text output summarizes all dimensions
- **WHEN** a summary is rendered as text
- **THEN** the output includes title, version, type, and answered count
- **AND** it includes one row per dimension in fixed `EI`, `SN`, `TF`, `JP` order
- **AND** each row includes letter, score, strength, and pole

#### Scenario: Text output avoids diagnostic claims
- **WHEN** a summary is rendered as text
- **THEN** it does not claim the result is a diagnosis
- **AND** it does not tell the user what their personality is
