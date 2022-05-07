package points_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go-bricks/internal/domain/points"
)

func TestPoints_GetPoints(t *testing.T) {
	t.Run("Initial points value should be zero", func(t *testing.T) {
		p := points.Points{}
		expected := 0

		actual := p.GetPoints()

		assert.Equal(t, expected, actual, "points value is not zero")
	})
}

func TestPoints_AddPoints(t *testing.T) {
	t.Run("Add 50 points single time", func(t *testing.T) {
		p := points.Points{}
		expected := 50
		p.AddPoints(50)

		actual := p.GetPoints()

		assert.Equal(t, expected, actual, "points value is not 50")
	})

	t.Run("Add 50 points two times", func(t *testing.T) {
		p := points.Points{}
		expected := 100
		p.AddPoints(50)
		p.AddPoints(50)

		actual := p.GetPoints()

		assert.Equal(t, expected, actual, "points value is not 100")
	})
}
