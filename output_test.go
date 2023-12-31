package output

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOutput(t *testing.T) {
	t.Run("it should output messages only for defined level and above", func(t *testing.T) {
		t.Parallel()

		// GIVEN
		debugPrinted := false
		infoPrinted := false
		warningPrinted := false
		errorPrinted := false

		expectedMessageLevel := LevelWarning
		testOutput, err := New(
			OptionLevel(expectedMessageLevel),
			OptionVerbosity(VerbosityNormal),
			OptionDebugPrinter(func(record Record, _ Output) {
				debugPrinted = true
			}),
			OptionInfoPrinter(func(record Record, _ Output) {
				infoPrinted = true
			}),
			OptionWarningPrinter(func(record Record, _ Output) {
				val, ok := record.Extra()["test_key"]
				assert.True(t, ok)
				assert.Equal(t, 9000, val)

				warningPrinted = true
			}),
			OptionErrorPrinter(func(record Record, _ Output) {
				errorPrinted = true

				assert.Equal(t, "Test error message", record.Message())
			}),
		)
		require.NoError(t, err)

		// WHEN
		Debug(testOutput, WithMessage("Test debug message"))
		Info(testOutput, WithMessage("Test info message"))
		Warning(testOutput, WithMessage("Test warning message"), WithExtra("test_key", 9000))
		Error(testOutput, WithParametrizedMessage("Test %s message", "error"))

		// THEN
		assert.False(t, debugPrinted)
		assert.False(t, infoPrinted)
		assert.True(t, warningPrinted)
		assert.True(t, errorPrinted)
	})
	t.Run("it should output messages only for defined verbosity", func(t *testing.T) {
		t.Parallel()

		// GIVEN
		verbNormal := false
		verbHigh := false
		verbAll := false

		testOutput, err := New(
			OptionLevel(LevelDebug),
			OptionVerbosity(VerbosityHigh),
			OptionDebugPrinter(func(record Record, _ Output) {
				switch record.Verbosity() {
				case VerbosityNormal:
					verbNormal = true
				case VerbosityHigh:
					verbHigh = true
				case VerbosityAll:
					verbAll = true
				}
			}),
		)
		require.NoError(t, err)

		// WHEN
		Debug(testOutput, WithMessage("Test normal message"), WithVerbosity(VerbosityNormal))
		Debug(testOutput, WithMessage("Test high message"), WithVerbosity(VerbosityHigh))
		Debug(testOutput, WithMessage("Test all message"), WithVerbosity(VerbosityAll))

		// THEN
		assert.True(t, verbNormal)
		assert.True(t, verbHigh)
		assert.False(t, verbAll)
	})
	t.Run("it should use printer defined using Printer interface", func(t *testing.T) {
		t.Parallel()

		// GIVEN
		testWriter := bytes.NewBuffer(nil)

		testOutput, err := New(
			OptionPrinter(NewWriterPrinter(testWriter)),
		)
		require.NoError(t, err)

		// WHEN
		Debug(testOutput, WithMessage("Test debug message"))
		Info(testOutput, WithMessage("Test info message"))
		Warning(testOutput, WithMessage("Test warning message"), WithExtra("test_key", 9000))
		Error(testOutput, WithParametrizedMessage("Test %s message", "error"))

		// THEN
		result := testWriter.String()
		assert.Contains(t, result, "Test debug message")
		assert.Contains(t, result, "Test info message")
		assert.Contains(t, result, "Test warning message")
		assert.Contains(t, result, "Test error message")
	})
	t.Run("it should allow for output to be nil and ignore all logs", func(t *testing.T) {
		t.Parallel()

		assert.NotPanics(t, func() {
			Debug(nil, WithMessage("Test debug message"))
			Info(nil, WithMessage("Test info message"))
			Warning(nil, WithMessage("Test warning message"), WithExtra("test_key", 9000))
			Error(nil, WithParametrizedMessage("Test %s message", "error"))
		})
	})
}
