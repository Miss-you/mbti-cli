# T02 BDD Plan Review

## Verdict

Approved. The plan is behavior-first, spec-aligned, and appropriately narrow for a loader-only task. It covers the user-visible `LoadFile` contract without drifting into schema validation, scoring, or CLI behavior that belongs to other tasks.

## Recommendations

- Keep the canonical load scenario focused on representative typed-bank assertions plus source metadata. Do not expand it into a full fixture regression of the entire JSON shape, because that would duplicate T01 and dilute the BDD signal.
- Keep the missing-file and malformed-JSON scenarios exactly as planned. They map cleanly to the loader’s read/parse contract and match the current error wrapping style.
- For the negative control, use one assertion that can be broken deliberately without fixture churn, such as the source filename or path-context check, so the RED step stays behavior-focused.

## Non-Goals

- Schema validation of bank contents
- Dimension, threshold, option, or reverse-flag correctness
- Answer parsing, validation, scoring, rendering, or CLI behavior
- Production code changes
