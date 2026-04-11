# T04 Alternative Implementation Notes

## Proposed Package

Add `internal/answers` with:

- `type File struct { Answers map[string]string }`
- `type Set map[string]string`
- `func Parse(data []byte) (Set, error)`

## Behavior

- Decode the top-level `answers` object from JSON.
- Require `answers` to be present and non-null.
- Return a fresh `Set` so callers cannot mutate parser internals.
- Normalize each option code with `strings.TrimSpace` and `strings.ToUpper`.
- Preserve question IDs exactly as supplied.

## Explicit Non-Behavior

- Do not validate whether question IDs exist in a bank.
- Do not validate whether option codes are one of `A/B/C/D`.
- Do not require all bank questions to be answered.
- Do not load from filesystem; future CLI integration can decide how file IO should be wired.
