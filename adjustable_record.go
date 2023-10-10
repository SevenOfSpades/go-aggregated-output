package output

import "fmt"

type adjustableRecord struct {
	alreadySet map[string]bool
	level      Level
	verbosity  Verbosity
	message    string
	extra      map[string]any
}

func newRecord() Record {
	return &adjustableRecord{
		level:     LevelInfo,
		verbosity: VerbosityNormal,
		extra:     make(map[string]any),
	}
}

func (r *adjustableRecord) Level() Level {
	return r.level
}

func (r *adjustableRecord) Verbosity() Verbosity {
	return r.verbosity
}

func (r *adjustableRecord) Message() string {
	return r.message
}

func (r *adjustableRecord) Extra() map[string]any {
	return r.extra
}

func (r *adjustableRecord) setLevel(level Level) {
	r.assertOrMarkAlreadySet("level")
	r.level = level
}

func (r *adjustableRecord) setVerbosity(verbosity Verbosity) {
	r.assertOrMarkAlreadySet("verbosity")
	r.verbosity = verbosity
}

func (r *adjustableRecord) setMessage(msg string) {
	r.assertOrMarkAlreadySet("message")
	r.message = msg
}

func (r *adjustableRecord) addExtra(key string, value any) {
	r.extra[key] = value
}

func (r *adjustableRecord) assertOrMarkAlreadySet(valName string) {
	if _, ok := r.alreadySet[valName]; ok {
		panic(fmt.Errorf("cannot set value for '%s': %w", valName, ErrAlreadySet))
	}
}
