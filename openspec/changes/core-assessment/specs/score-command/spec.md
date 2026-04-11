## ADDED Requirements

### Requirement: Score command evaluates canonical answer files
The system SHALL provide a non-interactive `score` CLI command that loads a question bank and canonical answer file, validates both inputs, scores answers, classifies the result, and writes the rendered result to stdout.

#### Scenario: Score command renders canonical answers as JSON
- **WHEN** `mbti-cli score --questions questions/mbti-questions-v3.json --answers answers.json --format json` is executed with a complete valid answer file
- **THEN** stdout contains parseable JSON
- **AND** the JSON includes bank title, bank version, answered count, total count, type, and all four dimensions
- **AND** the command exits successfully

#### Scenario: Score command renders readable text
- **WHEN** `mbti-cli score --questions questions/mbti-questions-v3.json --answers answers.json --format text` is executed with a complete valid answer file
- **THEN** stdout contains the factual result summary with type, answered count, and one row per dimension
- **AND** the command exits successfully

### Requirement: Score command preserves core scoring contracts
The system SHALL delegate answer parsing, strict answer validation, scoring, threshold classification, and rendering to the existing core packages rather than duplicating those rules in the Cobra command.

#### Scenario: Strict answer validation failures stop scoring output
- **WHEN** the answer file is missing an answer, contains an unknown question ID, or contains an invalid option code
- **THEN** the command returns an error
- **AND** stdout is empty

#### Scenario: Signed option scores are scored once
- **WHEN** a valid answer file is scored
- **THEN** the dimension totals come from the selected option scores
- **AND** reverse metadata is not applied a second time by the command

### Requirement: Score command reports invalid inputs without partial stdout
The system SHALL reject invalid flags, invalid question bank input, and invalid answer file input without writing partial result output.

#### Scenario: Unsupported output format fails
- **WHEN** the command is executed with an unsupported `--format`
- **THEN** it returns an error
- **AND** stdout is empty

#### Scenario: Missing or invalid file paths fail
- **WHEN** the command is executed without a usable question bank path or answer file path
- **THEN** it returns an error with path context
- **AND** stdout is empty

#### Scenario: Malformed answer JSON fails
- **WHEN** the command is executed with an answer file that is not valid canonical answer JSON
- **THEN** it returns an error with answer parsing context
- **AND** stdout is empty
