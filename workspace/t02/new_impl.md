# T02 New Implementation Options

## Option A: `LoadFile(path) (LoadedBank, error)`

Add a small loader to `internal/questionbank`:

- `LoadedBank` contains a typed `Bank` and `Source`.
- `Source` records the input path, base filename, and byte size.
- `LoadFile` handles empty path, read errors, and JSON parse errors.
- Errors wrap the underlying `os` or `encoding/json` error while adding operation and path context.

This is the recommended option. It keeps future CLI code simple without adding validation or command behavior early.

## Option B: `Load(path) (*Bank, Source, error)`

Return the model and metadata separately.

This is usable, but the multi-return API is less cohesive once later code passes the loaded bank through validators and scoring.

## Option C: Reader-Based Loader

Expose `Load(r io.Reader, source Source)` and keep filesystem handling outside the package.

This is more flexible, but T02 explicitly asks for file-path loading. It adds an abstraction before another caller needs it.

## Decision

Use Option A. It satisfies the task directly and keeps the package boundary clean:

- `model.go` remains representation-only.
- `loader.go` owns filesystem and JSON decoding.
- Validation remains for T03.
