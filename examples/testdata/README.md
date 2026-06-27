# Test Data Files for Grep Examples

This directory contains test data files used by the grep example tests. Each file is designed to demonstrate specific grep features.

## Files

- **fruits.txt** - Simple list for basic pattern matching
- **log_entries.txt** - Log file with different severity levels (INFO, WARN, ERROR)
- **code_snippets.txt** - JavaScript code for line number examples
- **mixed_case.txt** - Mixed case text for case-insensitive matching
- **config_file.txt** - Configuration file with comments for inverse matching
- **filenames.txt** - Filenames with special characters for fixed string matching
- **words.txt** - Words with common prefixes for whole word matching
- **empty_lines.txt** - Text with empty lines for empty line filtering

## Usage

These files are referenced in the `fromfile_*.go` test files using `yup.File()`:

```go
yup.MustRun(
    Grep("pattern", yup.File("testdata/fruits.txt")),
)
```

## Adding New Test Data

When adding new test data:
1. Create a descriptive filename
2. Keep files small and focused (< 10 lines typically)
3. Make data representative of real-world use cases
4. Document the file's purpose in this README

