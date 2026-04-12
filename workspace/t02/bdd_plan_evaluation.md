# T02 BDD Plan Review Evaluation

## Adopted

- Keep the canonical load scenario focused on representative typed-bank assertions plus source metadata.
  Reason: this keeps the BDD proof on loader behavior instead of turning it into a full fixture regression of the JSON shape.
- Keep the missing-file and malformed-JSON scenarios exactly as planned.
  Reason: both map directly to the loader read/parse contract and match the current error wrapping style.
- Use a negative control that can be broken without fixture churn, such as source filename or path-context.
  Reason: this keeps the RED step behavior-focused and easy to verify.

## Rejected

- None.

## Final BDD Shape

T02 will keep the planned loader-focused BDD coverage only:

- canonical v3 bank loads from a filesystem path and returns representative typed-bank fields plus source metadata
- missing file returns a read error with path context and preserves the missing-file cause
- malformed JSON returns a parse error with source path context
- empty path is rejected before reading

No plan changes were needed beyond preserving the current scope.
