## Purpose

Define threshold, pole, balanced, and type classification for raw scores.

## Requirements

### Requirement: Scoring classifier maps dimension totals to threshold strengths
The system SHALL provide an in-memory scoring classifier that accepts a question bank and a raw scoring result, then returns deterministic per-dimension classifications.

#### Scenario: Non-zero score uses matching threshold range
- **WHEN** a dimension score falls within one configured threshold range
- **THEN** the classifier labels that dimension with the matching strength bucket
- **AND** the classification preserves the raw score

#### Scenario: Threshold boundaries are inclusive
- **WHEN** a dimension score equals the lower or upper bound of a configured threshold range
- **THEN** the classifier treats that score as belonging to that threshold range

### Requirement: Scoring classifier maps score direction to pole metadata and type letters
The system SHALL map positive dimension scores to pole A and negative dimension scores to pole B using the question bank dimension metadata.

#### Scenario: Positive and negative scores produce pole labels and letters
- **WHEN** a dimension has a positive score
- **THEN** the classifier uses that dimension's `pole_a` text and first dimension letter
- **WHEN** a dimension has a negative score
- **THEN** the classifier uses that dimension's `pole_b` text and second dimension letter

### Requirement: Scoring classifier treats zero scores as balanced
The system SHALL classify zero dimension scores as balanced without requiring a threshold bucket.

#### Scenario: Zero score produces X and balanced pole
- **WHEN** a dimension score is `0`
- **THEN** the classifier marks the dimension as balanced
- **AND** the dimension letter is `X`
- **AND** the dimension pole is `balanced`

### Requirement: Scoring classifier generates type in fixed dimension order
The system SHALL generate type strings in fixed `EI`, `SN`, `TF`, `JP` order regardless of map iteration order.

#### Scenario: Mixed dimension scores produce deterministic type
- **WHEN** raw dimension scores include positive, negative, and zero values
- **THEN** the generated type string uses one letter per dimension in `EI`, `SN`, `TF`, `JP` order
- **AND** balanced dimensions contribute `X`

### Requirement: Scoring classifier reports invalid classifier inputs
The system SHALL return contextual errors instead of silently inventing classifier data when required bank metadata is unavailable or a non-zero score cannot be classified.

#### Scenario: Required threshold is missing
- **WHEN** a non-zero dimension score needs a threshold bucket
- **AND** the bank does not define a matching threshold range
- **THEN** classification returns an error for the uncovered score

#### Scenario: Dimension metadata is missing
- **WHEN** a non-zero dimension score needs pole metadata
- **AND** the bank does not define metadata for that dimension
- **THEN** classification returns an error for the missing dimension metadata
