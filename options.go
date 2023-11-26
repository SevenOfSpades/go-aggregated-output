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
	return func(r options.Resolver) {
		options.WriteOrPanic[Level](r, optionLevel, level)
	}
}

// OptionVerbosity will overwrite default verbosity of output (VerbosityAll) with provided value.
func OptionVerbosity(verbosity Verbosity) options.Option {
	return func(r options.Resolver) {
		options.WriteOrPanic[Verbosity](r, optionVerbosity, verbosity)
	}
}

// OptionPrinter sets all handlers using struct compatible with Printer interface.
func OptionPrinter(printer Printer) options.Option {
	return func(r options.Resolver) {
		options.WriteOrPanic[PrintFunc](r, optionDebugPrinterFunc, printer.Debug)
		options.WriteOrPanic[PrintFunc](r, optionInfoPrinterFunc, printer.Info)
		options.WriteOrPanic[PrintFunc](r, optionWarningPrinterFunc, printer.Warning)
		options.WriteOrPanic[PrintFunc](r, optionErrorPrinterFunc, printer.Error)
	}
}

// OptionDebugPrinter sets handler for debug output.
func OptionDebugPrinter(printFunc PrintFunc) options.Option {
	return func(r options.Resolver) {
		options.WriteOrPanic[PrintFunc](r, optionDebugPrinterFunc, printFunc)
	}
}

// OptionInfoPrinter sets handler for info output.
func OptionInfoPrinter(printFunc PrintFunc) options.Option {
	return func(r options.Resolver) {
		options.WriteOrPanic[PrintFunc](r, optionInfoPrinterFunc, printFunc)
	}
}

// OptionWarningPrinter sets handler for warning output.
func OptionWarningPrinter(printFunc PrintFunc) options.Option {
	return func(r options.Resolver) {
		options.WriteOrPanic[PrintFunc](r, optionWarningPrinterFunc, printFunc)
	}
}

// OptionErrorPrinter sets handler for error output.
func OptionErrorPrinter(printFunc PrintFunc) options.Option {
	return func(r options.Resolver) {
		options.WriteOrPanic[PrintFunc](r, optionErrorPrinterFunc, printFunc)
	}
}
