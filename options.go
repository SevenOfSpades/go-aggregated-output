package output

import "github.com/SevenOfSpades/go-just-options"

const (
	optionLevel              options.OptionKey = `level`
	optionVerbosity          options.OptionKey = `verbosity`
	optionDebugPrinterFunc   options.OptionKey = `debug_printer_func`
	optionInfoPrinterFunc    options.OptionKey = `info_printer_func`
	optionWarningPrinterFunc options.OptionKey = `warning_printer_func`
	optionErrorPrinterFunc   options.OptionKey = `error_printer_func`
)

// OptionLevel will overwrite default expected output (LevelDebug) with provided value.
func OptionLevel(level Level) options.Option {
	return func(o options.Options) {
		options.WriteOrPanic[Level](o, optionLevel, level)
	}
}

// OptionVerbosity will overwrite default verbosity of output (VerbosityAll) with provided value.
func OptionVerbosity(verbosity Verbosity) options.Option {
	return func(o options.Options) {
		options.WriteOrPanic[Verbosity](o, optionVerbosity, verbosity)
	}
}

// OptionDebugPrinter sets handler for debug output.
func OptionDebugPrinter(printFunc PrintFunc) options.Option {
	return func(o options.Options) {
		options.WriteOrPanic[PrintFunc](o, optionDebugPrinterFunc, printFunc)
	}
}

// OptionInfoPrinter sets handler for info output.
func OptionInfoPrinter(printFunc PrintFunc) options.Option {
	return func(o options.Options) {
		options.WriteOrPanic[PrintFunc](o, optionInfoPrinterFunc, printFunc)
	}
}

// OptionWarningPrinter sets handler for warning output.
func OptionWarningPrinter(printFunc PrintFunc) options.Option {
	return func(o options.Options) {
		options.WriteOrPanic[PrintFunc](o, optionWarningPrinterFunc, printFunc)
	}
}

// OptionErrorPrinter sets handler for error output.
func OptionErrorPrinter(printFunc PrintFunc) options.Option {
	return func(o options.Options) {
		options.WriteOrPanic[PrintFunc](o, optionErrorPrinterFunc, printFunc)
	}
}
