package output

import "fmt"

// Argument is a handler for setting values in Record during message construction.
type Argument func(Record)

// WithVerbosity will change Verbosity of Record (by default, it is VerbosityNormal).
func WithVerbosity(verbosity Verbosity) Argument {
	return func(r Record) {
		r.setVerbosity(verbosity)
	}
}

// WithMessage will set message in Record.
func WithMessage(msg string) Argument {
	return func(r Record) {
		r.setMessage(msg)
	}
}

// WithParametrizedMessage will set parameterized message in Record.
// It is the same as calling WithMessage(fmt.Sprintf(msg, params...)).
func WithParametrizedMessage(msg string, vals ...any) Argument {
	return WithMessage(fmt.Sprintf(msg, vals...))
}

// WithExtra will set special parameters that can be used for more detailed debugging.
func WithExtra(key string, value any) Argument {
	return func(r Record) {
		r.addExtra(key, value)
	}
}
