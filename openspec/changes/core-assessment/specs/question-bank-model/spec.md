## ADDED Requirements

### Requirement: Question bank model preserves canonical JSON fields
The system SHALL provide a Go model for the canonical question bank JSON shape, including top-level metadata, dimension metadata, scoring thresholds, questions, localized scenario text, localized option labels, option scores, and the question `reverse` flag.

#### Scenario: Canonical v3 bank unmarshals into model
- **WHEN** `questions/mbti-questions-v3.json` is unmarshaled into the question bank model
- **THEN** metadata, dimensions, thresholds, question ids, dimensions, reverse flags, scenarios, options, labels, and signed scores are available without field loss

### Requirement: Question bank model does not validate or score
The system SHALL keep T01 limited to data representation and JSON unmarshalling, without enforcing schema rules or applying scoring behavior.

#### Scenario: Model package remains representation-only
- **WHEN** the question bank model package is used by later tasks
- **THEN** validation, file loading, score aggregation, threshold classification, rendering, and CLI behavior are provided by later task-owned code, not by the T01 model
