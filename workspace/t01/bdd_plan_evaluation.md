# T01 BDD Plan Review Evaluation

## Adopted

1. Adopt: keep one BDD scenario centered on the canonical v3 bank unmarshaling into `questionbank.Bank`.
   - Reason: this is the only T01 behavior that is clearly spec-first and user-observable in the task board and OpenSpec contract.

2. Adopt: remove the `schema-invalid data is unmarshaled into the model` scenario.
   - Reason: T01 is a representation-only model task, but this negative case reads like validator coverage and belongs to later validation work.

3. Adopt: trim assertions to representative contract checks instead of round-trip JSON equality.
   - Reason: the task needs proof of field preservation, not a full structural mirror of the canonical JSON fixture.

4. Adopt: make any negative control prove the behavior assertion can fail, not that the Ginkgo suite is wired.
   - Reason: the failure signal should protect the contract itself, not just the harness.

5. Adopt: keep T01 out of loader, validation, scoring, rendering, and CLI coverage.
   - Reason: those behaviors are explicitly owned by later tasks and should not be pulled into the model BDD plan.

## Rejected

None.

## Final BDD Shape

- One Ginkgo spec file for `internal/questionbank`.
- One primary scenario: `questions/mbti-questions-v3.json` unmarshals into `questionbank.Bank`.
- Assertions limited to representative fields: title, version, total count, dimensions, threshold ranges, one first question, one reversed question, and a representative option label plus signed score.
- No schema-invalid fixture scenario.
- No loader, validation, scoring, rendering, or CLI coverage.
- Negative control, if used, should fail on the preserved-field assertion.
