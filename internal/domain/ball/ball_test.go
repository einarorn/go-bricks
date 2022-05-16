package ball_test

import (
	"go-bricks/internal/domain/ball"
	"go-bricks/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestBall_SliceLeft(t *testing.T) {
	t.Run("Slicing to left should should move the ball one step to left with each movement", func(t *testing.T) {
		b := ball.NewBall(8, 6)
		expectedAMove1 := models.Coordinate{
			X: 1,
			Y: 5,
		}
		expectedBMove1 := models.Coordinate{
			X: 2,
			Y: 5,
		}

		b.SliceLeft()
		b.Move()
		actual := b.GetBall()

		assert.Equal(t, expectedAMove1, actual.PositionA)
		assert.Equal(t, expectedBMove1, actual.PositionB)
	})
}

func TestBall_SliceRight(t *testing.T) {
	t.Run("Slicing to right should should move the ball one step to right with each movement", func(t *testing.T) {
		b := ball.NewBall(8, 6)
		expectedAMove1 := models.Coordinate{
			X: 3,
			Y: 5,
		}
		expectedBMove1 := models.Coordinate{
			X: 4,
			Y: 5,
		}

		b.SliceRight()
		b.Move()
		actual := b.GetBall()

		assert.Equal(t, expectedAMove1, actual.PositionA)
		assert.Equal(t, expectedBMove1, actual.PositionB)
	})
}

func TestBall_BounceVertical(t *testing.T) {
	t.Run("Vertical bounce should invert vertical direction", func(t *testing.T) {
		b := ball.NewBall(8, 6)
		expectedAMove1 := models.Coordinate{
			X: 2,
			Y: 6,
		}
		expectedBMove1 := models.Coordinate{
			X: 3,
			Y: 6,
		}

		b.Move()
		b.BounceVertical()
		b.Move()
		actual := b.GetBall()

		assert.Equal(t, expectedAMove1, actual.PositionA)
		assert.Equal(t, expectedBMove1, actual.PositionB)
	})
}

func TestBall_BounceHorizontal(t *testing.T) {
	t.Run("Horizontal bounce should invert horizontal direction", func(t *testing.T) {
		b := ball.NewBall(8, 6)
		expectedAMove1 := models.Coordinate{
			X: 2,
			Y: 4,
		}
		expectedBMove1 := models.Coordinate{
			X: 3,
			Y: 4,
		}

		b.SliceRight()
		b.Move()
		b.BounceHorizontal()
		b.Move()
		actual := b.GetBall()

		assert.Equal(t, expectedAMove1, actual.PositionA)
		assert.Equal(t, expectedBMove1, actual.PositionB)
	})
}
