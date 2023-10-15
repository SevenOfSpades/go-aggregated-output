package output

import "errors"

var ErrAlreadySet = errors.New("output argument has been already set")

type (
	Printer interface {
		Debug(Record, Output)
		Info(Record, Output)
		Warning(Record, Output)
		Error(Record, Output)
	}
	Record interface {
		Level() Level
		Verbosity() Verbosity
		Message() string
		Extra() map[string]any

		setLevel(Level)
		setVerbosity(Verbosity)
		setMessage(string)
		addExtra(key string, value any)
	}
	Output interface {
		Level() Level
		Verbosity() Verbosity

		printRecord(Record)
	}
)
