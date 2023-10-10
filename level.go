package output

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarning
	LevelError
)

// Level defines importance of message and expected output.
type Level uint8

func (l Level) Is(val Level) bool {
	return l == val
}

func (l Level) IsEqualOrHigherThan(val Level) bool {
	return l.Int() >= val.Int()
}

func (l Level) Int() int {
	return int(l)
}
