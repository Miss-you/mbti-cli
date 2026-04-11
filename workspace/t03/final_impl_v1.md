# T03 Final Implementation v1

## Final plan

Implement a small schema validator in `internal/questionbank`:

1. Add `ValidationError` with an `Issues []string` field and deterministic `Error()` formatting.
2. Add `Validate(bank Bank) error` that accumulates all detected issues.
3. Keep validation rules limited to the approved question-bank contract.
4. Add table-driven tests for the canonical v3 bank and invalid bank mutations.
5. Update the active `core-assessment` OpenSpec artifacts with a validator requirement and task checklist.

## Review pass

Score: 91/100

- CLI contract fidelity: 25/25. No command behavior is introduced.
- Question/scoring semantic fidelity: 24/25. `reverse` remains metadata and scores are validated without scoring.
- Go-native simplicity: 18/20. A plain error type and table tests are enough.
- Scope control: 14/15. The plan avoids loader and CLI changes.
- Verification clarity: 10/15. Needs concrete invalid cases and `openspec validate core-assessment`.

No high-severity issue found. The main improvement is to ensure table tests cover threshold completeness and dimension count mismatches, not only option-level errors.
