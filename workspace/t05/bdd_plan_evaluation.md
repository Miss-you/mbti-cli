# T05 BDD Plan Review Evaluation

## Adopted

- Keep the core strict-validation path: complete answers pass, unknown IDs are rejected, invalid option codes are rejected, missing answers are rejected, and aggregated errors stay deterministic.
  Reason: this matches `openspec/specs/answer-validation/spec.md` and stays on the `Validate` boundary.
- Keep a minimal lowercase/non-canonical negative control, but frame it as direct validator input and not parser compatibility.
  Reason: the spec says validation SHALL NOT normalize option codes itself, so this is still a valid validator-level check as long as it stays small.
- Keep aggregation assertions representative rather than mirroring every unit-test detail.
  Reason: the BDD should prove observable behavior, not duplicate `validator_test.go`.

## Rejected

- Delete the lowercase/non-normalizing case entirely.
  Reason: that would drop the only explicit negative control for the spec clause that validation does not normalize option codes itself.
- Recast the T05 BDD as parser normalization coverage.
  Reason: normalization belongs to the answer parser, not the validator.

## Final BDD Shape

- One positive scenario: a complete strict answer set validates successfully.
- One compact negative group: unknown question ID, invalid option code, missing answer, and a minimal lowercase/non-canonical input case that is rejected without normalization.
- One aggregation scenario: multiple issues are returned together in stable order across repeated runs.
