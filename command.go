package command

import (
	"bytes"
	"context"
	"fmt"
	"regexp"

	"github.com/destel/rill"
	gloo "github.com/gloo-foo/framework"
	"github.com/gloo-foo/framework/patterns"
)

// Error is the sole error type the package emits, enabling errors.Is checks.
type Error string

func (e Error) Error() string { return string(e) }

// ErrInvalidPattern wraps a regular expression that fails to compile (-E/-w).
const ErrInvalidPattern Error = "invalid pattern"

// GrepPattern is the pattern grep matches each line against: a literal
// substring by default, or a regular expression under -E/-w.
type GrepPattern string

// expression is regular-expression source text ready for regexp.Compile.
type expression string

// matcher reports whether a single line matches the configured pattern.
type matcher func([]byte) bool

// predicate reports whether a line should be emitted, after applying -v.
type predicate func([]byte) bool

// lineNumber is the 1-based ordinal a -n run prefixes to each emitted line.
type lineNumber uint64

// Grep returns a Command that filters lines matching pattern. pattern is a
// required positional argument; the remaining options select the match mode.
//
// Flags:
//   - GrepIgnoreCase (-i): case-insensitive matching
//   - GrepInvert (-v): emit non-matching lines
//   - GrepWholeLine (-x): match only when the whole line equals the pattern
//   - GrepExtended (-E): interpret the pattern as an extended regular expression
//   - GrepWord (-w): match only at word boundaries
//   - GrepLineNumbers (-n): prefix each emitted line with its line number
//   - GrepCount (-c): emit the count of matching lines instead of the lines
func Grep(pattern GrepPattern, opts ...any) gloo.Command[[]byte, []byte] {
	f, rest := fold(opts)
	gloo.NewParameters[gloo.File, struct{}](rest...)
	match, err := compileMatcher(pattern, f)
	if err != nil {
		return failCommand(err)
	}
	return dispatch(f, predicateFor(match, f))
}

// dispatch selects the command shape implied by the output-changing flags.
func dispatch(f flags, keep predicate) gloo.Command[[]byte, []byte] {
	if f.countEnabled {
		return countCommand(keep)
	}
	if f.lineNumbersEnabled {
		return lineNumberCommand(keep)
	}
	return filterCommand(keep)
}

// predicateFor wraps a matcher with the -v inversion.
func predicateFor(match matcher, f flags) predicate {
	if f.invertEnabled {
		return func(line []byte) bool { return !match(line) }
	}
	return predicate(match)
}

// compileMatcher builds the matcher for pattern under f, returning an error only
// when a regular-expression mode (-E or -w) receives an invalid pattern.
func compileMatcher(pattern GrepPattern, f flags) (matcher, error) {
	if expr, ok := regexExpr(pattern, f); ok {
		return regexMatcher(expr)
	}
	return fixedMatcher(pattern, f), nil
}

// regexExpr renders the regular-expression source for pattern when a regex mode
// (-w or -E) is active, reporting false when fixed-string matching applies.
func regexExpr(pattern GrepPattern, f flags) (expression, bool) {
	switch {
	case bool(f.wordEnabled):
		return caseFold(expression(`\b`+regexp.QuoteMeta(string(pattern))+`\b`), f), true
	case bool(f.extendedEnabled):
		return caseFold(anchored(expression(pattern), f), f), true
	default:
		return "", false
	}
}

// anchored wraps expr in ^(?:...)$ when -x demands a whole-line match. The
// non-capturing group binds the anchors to the whole pattern, so an alternation
// like "ab|cd" must match the entire line rather than just one branch.
func anchored(expr expression, f flags) expression {
	if f.wholeLineEnabled {
		return "^(?:" + expr + ")$"
	}
	return expr
}

// caseFold prefixes the case-insensitive flag when -i is active.
func caseFold(expr expression, f flags) expression {
	if f.ignoreCaseEnabled {
		return "(?i)" + expr
	}
	return expr
}

// regexMatcher compiles expr into a matcher, wrapping compile failures.
func regexMatcher(expr expression) (matcher, error) {
	re, err := regexp.Compile(string(expr))
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrInvalidPattern, err)
	}
	return re.Match, nil
}

// fixedMatcher builds a literal substring (or whole-line) matcher honoring -i.
func fixedMatcher(pattern GrepPattern, f flags) matcher {
	needle := foldCase([]byte(pattern), f)
	test := bytes.Contains
	if f.wholeLineEnabled {
		test = bytes.Equal
	}
	return func(line []byte) bool { return test(foldCase(line, f), needle) }
}

// foldCase lowercases b when -i is active, leaving it untouched otherwise.
func foldCase(b []byte, f flags) []byte {
	if f.ignoreCaseEnabled {
		return bytes.ToLower(b)
	}
	return b
}

// countCommand emits a single line: the number of lines satisfying keep.
func countCommand(keep predicate) gloo.Command[[]byte, []byte] {
	return patterns.Aggregate(func(lines [][]byte) ([]byte, error) {
		n := 0
		for _, line := range lines {
			if keep(line) {
				n++
			}
		}
		return fmt.Appendf(nil, "%d", n), nil
	})
}

// filterCommand emits each line satisfying keep, unchanged.
func filterCommand(keep predicate) gloo.Command[[]byte, []byte] {
	return patterns.Filter(func(line []byte) (bool, error) { return keep(line), nil })
}

// lineNumberCommand emits each kept line prefixed with its 1-based line number.
func lineNumberCommand(keep predicate) gloo.Command[[]byte, []byte] {
	return gloo.FuncCommand[[]byte, []byte](func(_ context.Context, in gloo.Stream[[]byte]) gloo.Stream[[]byte] {
		var n lineNumber
		return gloo.WrapFrom(rill.OrderedFlatMap(in.Chan(), 1, func(line []byte) <-chan rill.Try[[]byte] {
			n++
			return rill.FromSlice(numbered(n, line, keep), nil)
		}), in)
	})
}

// numbered returns the prefixed line as a one-element slice when kept, else empty.
func numbered(n lineNumber, line []byte, keep predicate) [][]byte {
	if !keep(line) {
		return [][]byte{}
	}
	return [][]byte{fmt.Appendf(nil, "%d:%s", n, line)}
}

// failCommand returns a Command that emits err once, regardless of input.
func failCommand(err error) gloo.Command[[]byte, []byte] {
	return gloo.FuncCommand[[]byte, []byte](func(_ context.Context, in gloo.Stream[[]byte]) gloo.Stream[[]byte] {
		return gloo.WrapFrom(rill.FromSlice[[]byte](nil, err), in)
	})
}
