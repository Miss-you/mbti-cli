## Purpose

Define the canonical question bank data model and loader contract.

## Requirements

### Requirement: Question bank model preserves canonical JSON fields
The system SHALL provide a Go model for the canonical question bank JSON shape, including top-level metadata, dimension metadata, scoring thresholds, questions, localized scenario text, localized option labels, option scores, and the question `reverse` flag.

#### Scenario: Canonical v3 bank unmarshals into model
- **WHEN** `questions/mbti-questions-v3.json` is unmarshaled into the question bank model
- **THEN** metadata, dimensions, thresholds, question ids, dimensions, reverse flags, scenarios, options, labels, and signed scores are available without field loss

### Requirement: Question bank model types do not validate or score
The system SHALL keep the T01 model types limited to data representation and JSON unmarshalling, without enforcing schema rules or applying scoring behavior.

#### Scenario: T01 model types remain representation-only
- **WHEN** the question bank model types are used by later tasks
- **THEN** validation, file loading, score aggregation, threshold classification, rendering, and CLI behavior are provided by task-owned code outside the T01 model types

### Requirement: Question bank loader reads a JSON file into the typed model
The system SHALL provide a file-path loader for question bank JSON files that returns the typed `Bank` and source metadata.

#### Scenario: Canonical v3 bank loads from path
- **WHEN** `questions/mbti-questions-v3.json` is loaded from its filesystem path
- **THEN** the returned result includes a typed `Bank` with representative metadata and questions
- **AND** the returned source metadata includes the input path, base filename, and positive byte size

#### Scenario: Missing bank file reports a clear read error
- **WHEN** the loader is called with a path that does not exist
- **THEN** it returns an error that includes the failed read operation and path context
- **AND** the error preserves the underlying missing-file error for callers

#### Scenario: Malformed bank file reports a clear parse error
- **WHEN** the loader reads a file that is not valid question bank JSON
- **THEN** it returns an error that includes parse context and the source path

#### Scenario: Empty bank path is rejected before reading
- **WHEN** the loader is called with an empty path
- **THEN** it returns a clear path-required error
