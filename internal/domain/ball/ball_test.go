package ball_test

import (
	"github.com/stretchr/testify/assert"
	"go-bricks/internal/domain/ball"
	"go-bricks/internal/models"
	"testing"
)

func TestBall_GetBall(t *testing.T) {
	t.Run("Initial position should be in the center and one step above bottom", func(t *testing.T) {
		b := ball.NewBall(8, 6)
		expectedA := models.Coordinate{
			X: 2,
			Y: 6,
		}
		expectedB := models.Coordinate{
			X: 3,
			Y: 6,
		}

		actual := b.GetBall()

		assert.Equal(t, expectedA, actual.PositionA)
		assert.Equal(t, expectedB, actual.PositionB)
	})

	t.Run("Initial movement should be one step up", func(t *testing.T) {
		b := ball.NewBall(8, 6)
		expectedA := models.Coordinate{
			X: 2,
			Y: 5,
		}
		expectedB := models.Coordinate{
			X: 3,
			Y: 5,
		}

		b.Move()
		actual := b.GetBall()

		assert.Equal(t, expectedA, actual.PositionA)
		assert.Equal(t, expectedB, actual.PositionB)
	})
}
