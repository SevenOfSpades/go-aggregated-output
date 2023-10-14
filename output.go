package output

import (
	"fmt"
	"os"
	"time"

	"github.com/SevenOfSpades/go-just-options"
)

// New will create new Output instance.
// It can be configured with options:
// * OptionLevel - set level of output.
// * OptionVerbosity - set verbosity of output.
// * Option{Debug|Info|Warning|Error}Printer - set corresponding handler for output level.
func New(opts ...options.Option) (Output, error) {
	opt := options.New().Resolve(opts...)

	optLevel, err := options.ReadOrDefault[Level](opt, optionLevel, LevelDebug)
	if err != nil {
		return nil, fmt.Errorf("output initialization failed: %w", err)
	}
	optVerbosity, err := options.ReadOrDefault[Verbosity](opt, optionVerbosity, VerbosityAll)
	if err != nil {
		return nil, fmt.Errorf("output initialization failed: %w", err)
	}
	optDebugPrinter, err := options.ReadOrDefault[PrintFunc](opt, optionDebugPrinterFunc, func(record Record, _ Output) {
		_, _ = os.Stdout.WriteString(fmt.Sprintf("<|DEBUG|> %s: %s\n", time.Now().Format(time.DateTime), record.Message()))
	})
	if err != nil {
		return nil, fmt.Errorf("output initialization failed: %w", err)
	}
	optInfoPrinter, err := options.ReadOrDefault[PrintFunc](opt, optionInfoPrinterFunc, func(record Record, _ Output) {
		_, _ = os.Stdout.WriteString(fmt.Sprintf("<|INFO|> %s: %s\n", time.Now().Format(time.DateTime), record.Message()))
	})
	if err != nil {
		return nil, fmt.Errorf("output initialization failed: %w", err)
	}
	optWarningPrinter, err := options.ReadOrDefault[PrintFunc](opt, optionWarningPrinterFunc, func(record Record, _ Output) {
		_, _ = os.Stdout.WriteString(fmt.Sprintf("<|WARNING|> %s: %s\n", time.Now().Format(time.DateTime), record.Message()))
	})
	if err != nil {
		return nil, fmt.Errorf("output initialization failed: %w", err)
	}
	optErrorPrinter, err := options.ReadOrDefault[PrintFunc](opt, optionErrorPrinterFunc, func(record Record, _ Output) {
		_, _ = os.Stderr.WriteString(fmt.Sprintf("<|ERROR|> %s: %s\n", time.Now().Format(time.DateTime), record.Message()))
	})
	if err != nil {
		return nil, fmt.Errorf("output initialization failed: %w", err)
	}

	return newOutput(
		optLevel,
		optVerbosity,
		optDebugPrinter,
		optInfoPrinter,
		optWarningPrinter,
		optErrorPrinter,
	), nil
}

func Debug(output Output, args ...Argument) {
	printOutput(output, append(args, func(r Record) {
		r.setLevel(LevelDebug)
	}))
}

func Info(output Output, args ...Argument) {
	printOutput(output, append(args, func(r Record) {
		r.setLevel(LevelInfo)
	}))
}

func Warning(output Output, args ...Argument) {
	printOutput(output, append(args, func(r Record) {
		r.setLevel(LevelWarning)
	}))
}

func Error(output Output, args ...Argument) {
	printOutput(output, append(args, func(r Record) {
		r.setLevel(LevelError)
	}))
}

func printOutput(output Output, args []Argument) {
	r := newRecord()
	for _, x := range args {
		x(r)
	}

	if !r.Level().IsEqualOrHigherThan(output.Level()) {
		return
	}

	if !r.Verbosity().IsPartOf(output.Verbosity()) {
		return
	}

	output.printRecord(r)
}
