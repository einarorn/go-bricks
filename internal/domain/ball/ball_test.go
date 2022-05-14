package ball_test

import (
	"github.com/stretchr/testify/assert"
	"go-bricks/internal/domain/ball"
	"go-bricks/internal/models"
	"testing"
)

func TestBall_GetBall(t *testing.T) {
	t.Run("Initial position should be in the center and one step above bottom", func(t *testing.T) {
		b := ball.NewBall(8, 3)
		expected := models.Coordinate{
			X: 1,
			Y: 6,
		}

		actual := b.GetBall()

		assert.Equal(t, expected, actual.Position)
	})
}
