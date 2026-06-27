// Package alias provides short names for grep command flags.
package alias

import command "github.com/gloo-foo/cmd-grep"

// Grep is the grep command constructor.
var Grep = command.Grep

// -i flag: case-insensitive matching
const IgnoreCase = command.GrepIgnoreCase

// default: case-sensitive matching
const NoIgnoreCase = command.GrepNoIgnoreCase

// -v flag: invert match (print non-matching lines)
const Invert = command.GrepInvert

// default: normal match
const NoInvert = command.GrepNoInvert

// -x flag: match only if entire line equals pattern
const WholeLine = command.GrepWholeLine

// default: substring match
const NoWholeLine = command.GrepNoWholeLine

// -E flag: extended regex (no-op, Go uses extended regex by default)
const Extended = command.GrepExtended

// default: fixed-string match
const NoExtended = command.GrepNoExtended

// -w flag: match pattern only at word boundaries
const Word = command.GrepWord

// default: no word boundary constraint
const NoWord = command.GrepNoWord

// -n flag: prepend line number to each matching line
const LineNumbers = command.GrepLineNumbers

// default: no line numbers
const NoLineNumbers = command.GrepNoLineNumbers

// -c flag: count matching lines instead of printing them
const Count = command.GrepCount

// default: print matching lines
const NoCount = command.GrepNoCount
