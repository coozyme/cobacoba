package aggregate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestActivity(t *testing.T) {
	newActivity, _ := NewActivity("bagus@bagus.com", "kerja bro")
	t.Run("new activity", func(t *testing.T) {
		assert.NotNil(t, newActivity)
	})

	t.Run("new activity without email", func(t *testing.T) {
		activity, err := NewActivity("", "kerja bro")
		assert.NotNil(t, activity)
		assert.Error(t, err)
	})

	t.Run("new activity without title", func(t *testing.T) {
		activity, err := NewActivity("bagus@bagus.com", "")
		assert.NotNil(t, activity)
		assert.Error(t, err)
	})

	t.Run("new activity", func(t *testing.T) {
		rebuildActivity := RebuildActivity(1, "bagus@bagus.com", "kerja bro")
		assert.NotNil(t, rebuildActivity)
	})
}
