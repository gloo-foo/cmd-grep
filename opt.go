package command

// grepIgnoreCaseFlag enables case-insensitive matching (-i).
type grepIgnoreCaseFlag bool

const (
	GrepIgnoreCase   grepIgnoreCaseFlag = true
	GrepNoIgnoreCase grepIgnoreCaseFlag = false
)

func (f grepIgnoreCaseFlag) Configure(flags *flags) { flags.ignoreCase = f }

// grepInvertFlag prints non-matching lines instead of matching ones (-v).
type grepInvertFlag bool

const (
	GrepInvert   grepInvertFlag = true
	GrepNoInvert grepInvertFlag = false
)

func (f grepInvertFlag) Configure(flags *flags) { flags.invert = f }

// grepWholeLineFlag matches only when the entire line equals the pattern (-x).
type grepWholeLineFlag bool

const (
	GrepWholeLine   grepWholeLineFlag = true
	GrepNoWholeLine grepWholeLineFlag = false
)

func (f grepWholeLineFlag) Configure(flags *flags) { flags.wholeLine = f }

// grepExtendedFlag interprets the pattern as an extended regular expression (-E).
type grepExtendedFlag bool

const (
	GrepExtended   grepExtendedFlag = true
	GrepNoExtended grepExtendedFlag = false
)

func (f grepExtendedFlag) Configure(flags *flags) { flags.extended = f }

// grepWordFlag matches the pattern only at word boundaries (-w).
type grepWordFlag bool

const (
	GrepWord   grepWordFlag = true
	GrepNoWord grepWordFlag = false
)

func (f grepWordFlag) Configure(flags *flags) { flags.word = f }

// grepLineNumbersFlag prefixes each emitted line with its 1-based line number (-n).
type grepLineNumbersFlag bool

const (
	GrepLineNumbers   grepLineNumbersFlag = true
	GrepNoLineNumbers grepLineNumbersFlag = false
)

func (f grepLineNumbersFlag) Configure(flags *flags) { flags.lineNumbers = f }

// grepCountFlag emits the count of matching lines instead of the lines (-c).
type grepCountFlag bool

const (
	GrepCount   grepCountFlag = true
	GrepNoCount grepCountFlag = false
)

func (f grepCountFlag) Configure(flags *flags) { flags.count = f }

// flags holds the parsed grep flag set.
type flags struct {
	ignoreCase  grepIgnoreCaseFlag
	invert      grepInvertFlag
	wholeLine   grepWholeLineFlag
	extended    grepExtendedFlag
	word        grepWordFlag
	lineNumbers grepLineNumbersFlag
	count       grepCountFlag
}
