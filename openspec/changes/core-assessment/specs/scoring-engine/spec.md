## ADDED Requirements

### Requirement: Scoring engine aggregates signed option scores by dimension
The system SHALL provide an in-memory scoring engine that accepts a question bank and an answer set, then returns deterministic raw totals grouped by question dimension.

#### Scenario: Complete answers produce dimension totals
- **WHEN** a complete strict answer set selects one option for each bank question
- **THEN** scoring adds each selected option's signed score to that question's dimension total
- **AND** the result includes answered and total question counts

### Requirement: Scoring engine does not apply reverse metadata
The system SHALL treat option score values as authoritative signed scores and SHALL NOT apply the question `reverse` flag as a second transformation during aggregation.

#### Scenario: Reversed question uses stored signed score
- **WHEN** a question has `reverse=true`
- **AND** the selected option has a signed score of `-2`
- **THEN** scoring adds `-2` to the dimension total
- **AND** scoring does not invert the score to `+2`

### Requirement: Scoring engine rejects invalid strict answer sets before aggregation
The system SHALL rely on bank-aware strict answer validation before aggregating scores.

#### Scenario: Missing answer prevents scoring
- **WHEN** the answer set omits a bank question
- **THEN** scoring returns a validation error
- **AND** scoring does not return partial dimension totals

### Requirement: Scoring engine leaves threshold and type behavior out of scope
The system SHALL NOT classify strength buckets, generate MBTI type strings, render JSON output, or render text output in the raw scoring engine.

#### Scenario: Raw score result has no classifier fields
- **WHEN** scoring succeeds
- **THEN** the result contains raw dimension totals
- **AND** the result does not contain threshold strength labels or type letters
