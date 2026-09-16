package command

// grepIgnoreCaseFlag enables case-insensitive matching (-i).
type grepIgnoreCaseFlag bool

const (
	GrepIgnoreCase   grepIgnoreCaseFlag = true
	GrepNoIgnoreCase grepIgnoreCaseFlag = false
)

// grepInvertFlag prints non-matching lines instead of matching ones (-v).
type grepInvertFlag bool

const (
	GrepInvert   grepInvertFlag = true
	GrepNoInvert grepInvertFlag = false
)

// grepWholeLineFlag matches only when the entire line equals the pattern (-x).
type grepWholeLineFlag bool

const (
	GrepWholeLine   grepWholeLineFlag = true
	GrepNoWholeLine grepWholeLineFlag = false
)

// grepExtendedFlag interprets the pattern as an extended regular expression (-E).
type grepExtendedFlag bool

const (
	GrepExtended   grepExtendedFlag = true
	GrepNoExtended grepExtendedFlag = false
)

// grepWordFlag matches the pattern only at word boundaries (-w).
type grepWordFlag bool

const (
	GrepWord   grepWordFlag = true
	GrepNoWord grepWordFlag = false
)

// grepLineNumbersFlag prefixes each emitted line with its 1-based line number (-n).
type grepLineNumbersFlag bool

const (
	GrepLineNumbers   grepLineNumbersFlag = true
	GrepNoLineNumbers grepLineNumbersFlag = false
)

// grepCountFlag emits the count of matching lines instead of the lines (-c).
type grepCountFlag bool

const (
	GrepCount   grepCountFlag = true
	GrepNoCount grepCountFlag = false
)

// flags holds the parsed grep flag set.
type flags struct {
	ignoreCaseEnabled  grepIgnoreCaseFlag
	invertEnabled      grepInvertFlag
	wholeLineEnabled   grepWholeLineFlag
	extendedEnabled    grepExtendedFlag
	wordEnabled        grepWordFlag
	lineNumbersEnabled grepLineNumbersFlag
	countEnabled       grepCountFlag
}

// fold partitions opts: grep's own option values are folded into the flag set,
// and every other argument is passed through unchanged for the framework's
// positional classifier.
func fold(opts []any) (flags, []any) {
	var f flags
	rest := make([]any, 0, len(opts))
	for _, o := range opts {
		switch v := o.(type) {
		case grepIgnoreCaseFlag:
			f.ignoreCaseEnabled = v
		case grepInvertFlag:
			f.invertEnabled = v
		case grepWholeLineFlag:
			f.wholeLineEnabled = v
		case grepExtendedFlag:
			f.extendedEnabled = v
		case grepWordFlag:
			f.wordEnabled = v
		case grepLineNumbersFlag:
			f.lineNumbersEnabled = v
		case grepCountFlag:
			f.countEnabled = v
		default:
			rest = append(rest, o)
		}
	}
	return f, rest
}
