package output

type defaultOutput struct {
	level          Level
	verbosity      Verbosity
	debugPrinter   PrintFunc
	infoPrinter    PrintFunc
	warningPrinter PrintFunc
	errorPrinter   PrintFunc
}

func newOutput(
	level Level,
	verbosity Verbosity,
	debugPrinter PrintFunc,
	infoPrinter PrintFunc,
	warningPrinter PrintFunc,
	errorPrinter PrintFunc,
) Output {
	return &defaultOutput{
		level:          level,
		verbosity:      verbosity,
		debugPrinter:   debugPrinter,
		infoPrinter:    infoPrinter,
		warningPrinter: warningPrinter,
		errorPrinter:   errorPrinter,
	}
}

func (o *defaultOutput) Level() Level {
	return o.level
}

func (o *defaultOutput) Verbosity() Verbosity {
	return o.verbosity
}

func (o *defaultOutput) printRecord(record Record) {
	switch record.Level() {
	case LevelDebug:
		o.debugPrinter(record, o)
	case LevelInfo:
		o.infoPrinter(record, o)
	case LevelWarning:
		o.warningPrinter(record, o)
	case LevelError:
		o.errorPrinter(record, o)
	}
}
