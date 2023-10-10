package output

const (
	VerbosityNormal Verbosity = iota
	VerbosityHigh
	VerbosityAll
)

var verbosityRelations = map[Verbosity]map[Verbosity]bool{
	VerbosityNormal: {
		VerbosityNormal: true,
	},
	VerbosityHigh: {
		VerbosityNormal: true,
		VerbosityHigh:   true,
	},
	VerbosityAll: {
		VerbosityNormal: true,
		VerbosityHigh:   true,
		VerbosityAll:    true,
	},
}

// Verbosity can define details within given Level of output.
type Verbosity uint8

func (v Verbosity) IsPartOf(val Verbosity) bool {
	return verbosityRelations[val][v]
}
