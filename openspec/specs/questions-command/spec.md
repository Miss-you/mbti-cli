## Purpose

Define the non-interactive command for exporting selected question sets.

## Requirements

### Requirement: Questions command exports a selected question set
The system SHALL provide a non-interactive `questions` CLI command that loads a question bank from a file path, validates it, selects questions, and writes the selected question set to stdout.

#### Scenario: Questions command renders canonical bank as JSON
- **WHEN** `mbti-cli questions --questions questions/mbti-questions-v3.json --format json` is executed
- **THEN** stdout contains parseable JSON
- **AND** the JSON includes bank title, version, selected count, total count, language, source path, and exported questions
- **AND** the command exits successfully

#### Scenario: Questions command renders localized readable text
- **WHEN** `mbti-cli questions --questions questions/mbti-questions-v3.json --format text --count 1 --lang en` is executed
- **THEN** stdout contains one English question with option codes and option labels
- **AND** the command exits successfully

### Requirement: Questions command keeps scoring internals out of exported prompts
The system SHALL export only the information needed to answer selected questions and SHALL NOT expose option scores, reverse flags, thresholds, or dimension metadata in the `questions` command output.

#### Scenario: JSON question export omits scoring internals
- **WHEN** `mbti-cli questions --questions questions/mbti-questions-v3.json --format json` is executed
- **THEN** each exported question includes an ID, localized scenario text, and localized options
- **AND** exported options include code and localized label
- **AND** stdout does not include option scores, reverse flags, thresholds, or dimension metadata

### Requirement: Questions command supports bounded deterministic selection
The system SHALL support selecting all questions by default, selecting a bounded count with `--count`, and deterministic ordering when `--seed` is provided.

#### Scenario: Count limits selected questions
- **WHEN** the command is executed with `--count 3`
- **THEN** stdout contains exactly 3 exported questions

#### Scenario: Seeded selection is deterministic
- **WHEN** the command is executed twice with the same `--count` and `--seed` values
- **THEN** both runs produce the same selected question IDs in the same order

### Requirement: Questions command reports invalid inputs without partial stdout
The system SHALL reject invalid flags or invalid question bank input and avoid writing partial command output on failure.

#### Scenario: Unsupported output format fails
- **WHEN** the command is executed with an unsupported `--format`
- **THEN** it returns an error
- **AND** stdout is empty

#### Scenario: Unsupported language fails
- **WHEN** the command is executed with an unsupported `--lang`
- **THEN** it returns an error
- **AND** stdout is empty

#### Scenario: Invalid count fails
- **WHEN** the command is executed with a negative count or a count larger than the loaded bank
- **THEN** it returns an error
- **AND** stdout is empty

#### Scenario: Missing or invalid bank fails
- **WHEN** the command is executed without a usable question bank path
- **THEN** it returns an error with question bank path context
- **AND** stdout is empty
