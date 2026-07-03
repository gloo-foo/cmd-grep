package alias_test

import (
	"slices"
	"testing"

	gloo "github.com/gloo-foo/framework"
	"github.com/gloo-foo/testable"

	grep "github.com/gloo-foo/cmd-grep/alias"
)

// The alias package re-exports the constructor and flag constants under
// unprefixed names. A mis-wired re-export (say, Invert bound to the disabled
// constant, or Grep bound to the wrong function) compiles cleanly, so only
// behavior can prove the wiring. Each test exercises one re-export and asserts
// the GNU grep output it must produce.

const matchInput = "alpha\nbeta\nALPHA\nalphabet\n"

func assertLines(t *testing.T, got, want []string) {
	t.Helper()
	if !slices.Equal(got, want) {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func lines(t *testing.T, cmd gloo.Command[[]byte, []byte]) []string {
	t.Helper()
	out, err := testable.TestLines(cmd, matchInput)
	if err != nil {
		t.Fatal(err)
	}
	return out
}

func TestAlias_GrepFiltersMatchingLines(t *testing.T) {
	// Default: literal substring match on the pattern.
	assertLines(t, lines(t, grep.Grep("alpha")), []string{"alpha", "alphabet"})
}

func TestAlias_IgnoreCaseAndNoIgnoreCase(t *testing.T) {
	// -i folds case; the disabled form behaves like the default.
	assertLines(t, lines(t, grep.Grep("ALPHA", grep.IgnoreCase)),
		[]string{"alpha", "ALPHA", "alphabet"})
	assertLines(t, lines(t, grep.Grep("ALPHA", grep.NoIgnoreCase)),
		[]string{"ALPHA"})
}

func TestAlias_InvertAndNoInvert(t *testing.T) {
	// -v prints non-matching lines; the disabled form matches normally.
	assertLines(t, lines(t, grep.Grep("alpha", grep.Invert)),
		[]string{"beta", "ALPHA"})
	assertLines(t, lines(t, grep.Grep("alpha", grep.NoInvert)),
		[]string{"alpha", "alphabet"})
}

func TestAlias_WholeLineAndNoWholeLine(t *testing.T) {
	// -x matches only when the whole line equals the pattern.
	assertLines(t, lines(t, grep.Grep("alpha", grep.WholeLine)),
		[]string{"alpha"})
	assertLines(t, lines(t, grep.Grep("alpha", grep.NoWholeLine)),
		[]string{"alpha", "alphabet"})
}

func TestAlias_ExtendedAndNoExtended(t *testing.T) {
	// -E treats the pattern as a regular expression; disabled is a literal match.
	assertLines(t, lines(t, grep.Grep("al.ha", grep.Extended)),
		[]string{"alpha", "alphabet"})
	assertLines(t, lines(t, grep.Grep("al.ha", grep.NoExtended)),
		[]string{})
}

func TestAlias_WordAndNoWord(t *testing.T) {
	// -w matches only at word boundaries; disabled allows substrings.
	assertLines(t, lines(t, grep.Grep("alpha", grep.Word)),
		[]string{"alpha"})
	assertLines(t, lines(t, grep.Grep("alpha", grep.NoWord)),
		[]string{"alpha", "alphabet"})
}

func TestAlias_LineNumbersAndNoLineNumbers(t *testing.T) {
	// -n prefixes each emitted line with its 1-based line number.
	assertLines(t, lines(t, grep.Grep("alpha", grep.LineNumbers)),
		[]string{"1:alpha", "4:alphabet"})
	assertLines(t, lines(t, grep.Grep("alpha", grep.NoLineNumbers)),
		[]string{"alpha", "alphabet"})
}

func TestAlias_CountAndNoCount(t *testing.T) {
	// -c emits the count of matching lines instead of the lines themselves.
	assertLines(t, lines(t, grep.Grep("alpha", grep.Count)),
		[]string{"2"})
	assertLines(t, lines(t, grep.Grep("alpha", grep.NoCount)),
		[]string{"alpha", "alphabet"})
}
