package output

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLevel_IsEqualOrHigherThan(t *testing.T) {
	t.Run("it should properly detect level hierarchy", func(t *testing.T) {
		t.Parallel()

		// GIVEN
		levelDebug := LevelDebug
		levelInfo := LevelInfo

		// WHEN
		isLevelDebugEqualOrHigherThanLevelInfo := levelDebug.IsEqualOrHigherThan(levelInfo)
		isLevelInfoEqualOrHigherThanLevelDebug := levelInfo.IsEqualOrHigherThan(levelInfo)
		isLevelInfoEqualOrHigherThanLevelInfo := levelInfo.IsEqualOrHigherThan(levelInfo)

		// THEN
		assert.False(t, isLevelDebugEqualOrHigherThanLevelInfo)
		assert.True(t, isLevelInfoEqualOrHigherThanLevelDebug)
		assert.True(t, isLevelInfoEqualOrHigherThanLevelInfo)
	})
}
