# go-aggregated-output

Easy to use output aggregator in form of basic log functions with level and verbosity.

## Usage
```go
package main

import (
	"log"

	output "github.com/SevenOfSpades/go-aggregated-output"
)

func main() {
	// Create new output instance and set desired level and verbosity of output.
	o, err := output.New(output.OptionLevel(output.LevelDebug), output.OptionVerbosity(output.VerbosityHigh))
	if err != nil {
		log.Fatalln(err)
	}

	// This message won't show up in output because verbosity is set to `high` but message is meant for `all`.
	output.Debug(o, output.WithMessage("Debug message with verbosity"), output.WithVerbosity(output.VerbosityAll))
	// Output will receive "Info message with parameter".
	output.Info(o, output.WithParametrizedMessage("Info message with %s", "parameter"))
	// Just a warning message.
	output.Warning(o, output.WithMessage("Warning message"))
	// Extra parameters are passed to output as map[string]any and can be processed within handler.
	output.Error(o, output.WithMessage("Error message"), output.WithExtra("some_additional_parameter", 42))

	// Handlers for each output log function can be overridden when instantiating output.
	o, err = output.New(
		output.OptionDebugPrinter(func(record output.Record) {
			// Everything here will be executed when `output.Debug` is called with this instance of output.
		}),
		output.OptionInfoPrinter(func(record output.Record) {
			// Everything here will be executed when `output.Info` is called with this instance of output.
		}),
		output.OptionWarningPrinter(func(record output.Record) {
			// Everything here will be executed when `output.Warning` is called with this instance of output.
		}),
		output.OptionErrorPrinter(func(record output.Record) {
			// Everything here will be executed when `output.Error` is called with this instance of output.
		}),
	)
	if err != nil {
		log.Fatalln(err)
	}
}

```

### Log Levels

* Debug (`output.LevelDebug` is defined as `output.Level(0)` with underlying type `uint8`)
* Info (`output.LevelInfo` is defined as `output.Level(1)` with underlying type `uint8`)
* Warning (`output.LevelWarning` is defined as `output.Level(2)` with underlying type `uint8`)
* Error (`output.LevelError` is defined as `output.Level(3)` with underlying type `uint8`)

### Verbosity

Verbosity allows to skip certain logs within selected level. It should be treated as "level of details".
Below table explains how verbosity levels interacts with each other.

|  Verbosity   |     Output "normal"     |      Output "high"      |      Output "all"       |
|:------------:|:-----------------------:|:-----------------------:|:-----------------------:|
| Log "normal" | Record passed to output | Record passed to output | Record passed to output |
|  Log "high"  |     Record ignored      | Record passed to output | Record passed to output |
|  Log "all"   |     Record ignored      |     Record ignored      | Record passed to output |
