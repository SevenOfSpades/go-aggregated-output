package output

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerbosity_IsPartOf(t *testing.T) {
	t.Run("it should report normal being part of all", func(t *testing.T) {
		t.Parallel()

		// GIVEN
		normalVerbosity := VerbosityNormal

		// THEN
		isPartOfNormal := normalVerbosity.IsPartOf(VerbosityNormal)
		isPartOfHigh := normalVerbosity.IsPartOf(VerbosityHigh)
		isPartOfAll := normalVerbosity.IsPartOf(VerbosityAll)

		// THEN
		assert.True(t, isPartOfNormal)
		assert.True(t, isPartOfHigh)
		assert.True(t, isPartOfAll)
	})
	t.Run("it should report high being part of high and all", func(t *testing.T) {
		t.Parallel()

		// GIVEN
		highVerbosity := VerbosityHigh

		// THEN
		isPartOfNormal := highVerbosity.IsPartOf(VerbosityNormal)
		isPartOfHigh := highVerbosity.IsPartOf(VerbosityHigh)
		isPartOfAll := highVerbosity.IsPartOf(VerbosityAll)

		// THEN
		assert.False(t, isPartOfNormal)
		assert.True(t, isPartOfHigh)
		assert.True(t, isPartOfAll)
	})
	t.Run("it should report all being part of only all", func(t *testing.T) {
		t.Parallel()

		// GIVEN
		allVerbosity := VerbosityAll

		// THEN
		isPartOfNormal := allVerbosity.IsPartOf(VerbosityNormal)
		isPartOfHigh := allVerbosity.IsPartOf(VerbosityHigh)
		isPartOfAll := allVerbosity.IsPartOf(VerbosityAll)

		// THEN
		assert.False(t, isPartOfNormal)
		assert.False(t, isPartOfHigh)
		assert.True(t, isPartOfAll)
	})
}
