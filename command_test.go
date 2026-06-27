package command_test

import (
	"errors"
	"testing"

	command "github.com/gloo-foo/cmd-grep"

	"github.com/gloo-foo/testable"
)

func TestGrep_BasicMatch(t *testing.T) {
	lines, err := testable.TestLines(command.Grep("hello"), "hello world\ngoodbye world\nhello again\n")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(lines) != 2 {
		t.Fatalf("expected 2 lines, got %d: %v", len(lines), lines)
	}
	if lines[0] != "hello world" {
		t.Errorf("line 0: got %q, want %q", lines[0], "hello world")
	}
	if lines[1] != "hello again" {
		t.Errorf("line 1: got %q, want %q", lines[1], "hello again")
	}
}

func TestGrep_IgnoreCase(t *testing.T) {
	lines, err := testable.TestLines(command.Grep("HELLO", command.GrepIgnoreCase), "Hello world\ngoodbye world\nhELLo again\n")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(lines) != 2 {
		t.Fatalf("expected 2 lines, got %d: %v", len(lines), lines)
	}
	if lines[0] != "Hello world" {
		t.Errorf("line 0: got %q, want %q", lines[0], "Hello world")
	}
	if lines[1] != "hELLo again" {
		t.Errorf("line 1: got %q, want %q", lines[1], "hELLo again")
	}
}

func TestGrep_Invert(t *testing.T) {
	lines, err := testable.TestLines(command.Grep("b", command.GrepInvert), "a\nb\nc\n")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(lines) != 2 {
		t.Fatalf("expected 2 lines, got %d: %v", len(lines), lines)
	}
	if lines[0] != "a" {
		t.Errorf("line 0: got %q, want %q", lines[0], "a")
	}
	if lines[1] != "c" {
		t.Errorf("line 1: got %q, want %q", lines[1], "c")
	}
}

func TestGrep_NoMatches(t *testing.T) {
	lines, err := testable.TestLines(command.Grep("xyz"), "a\nb\nc\n")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(lines) != 0 {
		t.Fatalf("expected 0 lines, got %d: %v", len(lines), lines)
	}
}

func TestGrep_EmptyInput(t *testing.T) {
	lines, err := testable.TestLines(command.Grep("anything"), "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(lines) != 0 {
		t.Fatalf("expected 0 lines, got %d: %v", len(lines), lines)
	}
}

func TestGrep_MultipleMatches(t *testing.T) {
	lines, err := testable.TestLines(command.Grep("line"), "line one\nline two\nno match\nline three\n")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(lines) != 3 {
		t.Fatalf("expected 3 lines, got %d: %v", len(lines), lines)
	}
	if lines[0] != "line one" {
		t.Errorf("line 0: got %q, want %q", lines[0], "line one")
	}
	if lines[1] != "line two" {
		t.Errorf("line 1: got %q, want %q", lines[1], "line two")
	}
	if lines[2] != "line three" {
		t.Errorf("line 2: got %q, want %q", lines[2], "line three")
	}
}

func TestGrep_WholeLine(t *testing.T) {
	lines, err := testable.TestLines(command.Grep("hello", command.GrepWholeLine), "hello\nhello world\n")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(lines) != 1 {
		t.Fatalf("expected 1 line, got %d: %v", len(lines), lines)
	}
	if lines[0] != "hello" {
		t.Errorf("line 0: got %q, want %q", lines[0], "hello")
	}
}

func TestGrep_WholeLine_IgnoreCase(t *testing.T) {
	lines, err := testable.TestLines(command.Grep("HELLO", command.GrepWholeLine, command.GrepIgnoreCase), "Hello\nhello world\n")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(lines) != 1 {
		t.Fatalf("expected 1 line, got %d: %v", len(lines), lines)
	}
	if lines[0] != "Hello" {
		t.Errorf("line 0: got %q, want %q", lines[0], "Hello")
	}
}

func TestGrep_Extended(t *testing.T) {
	lines, err := testable.TestLines(command.Grep("he.lo", command.GrepExtended), "hello\nworld\n")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(lines) != 1 {
		t.Fatalf("expected 1 line, got %d: %v", len(lines), lines)
	}
	if lines[0] != "hello" {
		t.Errorf("line 0: got %q, want %q", lines[0], "hello")
	}
}

func TestGrep_Word(t *testing.T) {
	lines, err := testable.TestLines(command.Grep("he", command.GrepWord), "he\nhello\nthe\n")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(lines) != 1 {
		t.Fatalf("expected 1 line, got %d: %v", len(lines), lines)
	}
	if lines[0] != "he" {
		t.Errorf("line 0: got %q, want %q", lines[0], "he")
	}
}

func TestGrep_Word_IgnoreCase(t *testing.T) {
	lines, err := testable.TestLines(command.Grep("HE", command.GrepWord, command.GrepIgnoreCase), "he\nhello\nthe\n")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(lines) != 1 {
		t.Fatalf("expected 1 line, got %d: %v", len(lines), lines)
	}
	if lines[0] != "he" {
		t.Errorf("line 0: got %q, want %q", lines[0], "he")
	}
}

func TestGrep_LineNumbers(t *testing.T) {
	lines, err := testable.TestLines(command.Grep("hello", command.GrepLineNumbers), "foo\nhello\nbar\nhello world\n")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(lines) != 2 {
		t.Fatalf("expected 2 lines, got %d: %v", len(lines), lines)
	}
	if lines[0] != "2:hello" {
		t.Errorf("line 0: got %q, want %q", lines[0], "2:hello")
	}
	if lines[1] != "4:hello world" {
		t.Errorf("line 1: got %q, want %q", lines[1], "4:hello world")
	}
}

func TestGrep_LineNumbers_Invert(t *testing.T) {
	lines, err := testable.TestLines(command.Grep("b", command.GrepLineNumbers, command.GrepInvert), "a\nb\nc\n")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(lines) != 2 {
		t.Fatalf("expected 2 lines, got %d: %v", len(lines), lines)
	}
	if lines[0] != "1:a" {
		t.Errorf("line 0: got %q, want %q", lines[0], "1:a")
	}
	if lines[1] != "3:c" {
		t.Errorf("line 1: got %q, want %q", lines[1], "3:c")
	}
}

func TestGrep_Count(t *testing.T) {
	lines, err := testable.TestLines(command.Grep("hello", command.GrepCount), "foo\nhello\nbar\nhello world\n")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(lines) != 1 {
		t.Fatalf("expected 1 line, got %d: %v", len(lines), lines)
	}
	if lines[0] != "2" {
		t.Errorf("got %q, want %q", lines[0], "2")
	}
}

func TestGrep_Count_NoMatches(t *testing.T) {
	lines, err := testable.TestLines(command.Grep("xyz", command.GrepCount), "a\nb\nc\n")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(lines) != 1 {
		t.Fatalf("expected 1 line, got %d: %v", len(lines), lines)
	}
	if lines[0] != "0" {
		t.Errorf("got %q, want %q", lines[0], "0")
	}
}

func TestGrep_Count_Invert(t *testing.T) {
	lines, err := testable.TestLines(command.Grep("b", command.GrepCount, command.GrepInvert), "a\nb\nc\n")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(lines) != 1 {
		t.Fatalf("expected 1 line, got %d: %v", len(lines), lines)
	}
	if lines[0] != "2" {
		t.Errorf("got %q, want %q", lines[0], "2")
	}
}

func TestGrep_Count_EmptyInput(t *testing.T) {
	lines, err := testable.TestLines(command.Grep("anything", command.GrepCount), "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(lines) != 1 {
		t.Fatalf("expected 1 line, got %d: %v", len(lines), lines)
	}
	if lines[0] != "0" {
		t.Errorf("got %q, want %q", lines[0], "0")
	}
}

func TestGrep_Extended_WholeLine(t *testing.T) {
	// -E -x: the regex is anchored to the whole line, so "ab" matches the line
	// "ab" but not "abc" — the alternation only matches a complete line.
	lines, err := testable.TestLines(command.Grep("ab|cd", command.GrepExtended, command.GrepWholeLine), "ab\nabc\ncd\nxcdx\n")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(lines) != 2 {
		t.Fatalf("expected 2 lines, got %d: %v", len(lines), lines)
	}
	if lines[0] != "ab" {
		t.Errorf("line 0: got %q, want %q", lines[0], "ab")
	}
	if lines[1] != "cd" {
		t.Errorf("line 1: got %q, want %q", lines[1], "cd")
	}
}

func TestGrep_Extended_InvalidPattern(t *testing.T) {
	// An unterminated group is not a valid regular expression. -E must surface a
	// pattern error rather than panic or silently match nothing.
	_, err := testable.TestLines(command.Grep("a(b", command.GrepExtended), "abc\n")
	if err == nil {
		t.Fatal("expected an error for an invalid pattern, got nil")
	}
	if !errors.Is(err, command.ErrInvalidPattern) {
		t.Fatalf("got %v, want errors.Is ErrInvalidPattern", err)
	}
}

func TestGrep_Extended_InvalidPattern_EmptyInput(t *testing.T) {
	// The pattern fails to compile before any input is read, so the error is
	// reported even when the input is empty.
	_, err := testable.TestLines(command.Grep("a(b", command.GrepExtended), "")
	if !errors.Is(err, command.ErrInvalidPattern) {
		t.Fatalf("got %v, want errors.Is ErrInvalidPattern", err)
	}
}

func TestErrInvalidPattern_Message(t *testing.T) {
	if got := command.ErrInvalidPattern.Error(); got != "invalid pattern" {
		t.Errorf("got %q, want %q", got, "invalid pattern")
	}
}
