# T05 Original Implementation

## Current behavior

- `internal/answers.Parse(data)` parses the canonical answer-file shape:
  `{"answers":{"q01":"A"}}`.
- The parser normalizes option values with `strings.TrimSpace` and
  `strings.ToUpper`.
- The parser intentionally does not check whether question IDs exist in a
  bank, whether option codes are valid for a question, or whether all bank
  questions have answers.
- `internal/questionbank.Validate(bank)` validates the bank schema only. It
  already guarantees supported dimensions, unique question IDs, four options,
  valid option codes, supported scores, and threshold shape.

## T05 gap

There is no bank-aware answer validation step. A parsed answer set can contain:

- unknown question IDs
- invalid option codes for a known question
- missing answers for bank questions

Downstream scoring would have to rediscover these errors unless T05 adds a
focused validation API.
