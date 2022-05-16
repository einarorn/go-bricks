package paddle_test

import (
	"go-bricks/internal/domain/paddle"
	"go-bricks/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPaddle_GetPaddle(t *testing.T) {
	t.Run("Initial position should be in the center and at bottom", func(t *testing.T) {
		p := paddle.NewPaddle(16, 14, 10)
		expectedPositionA := models.Coordinate{
			X: 2,
			Y: 15,
		}
		expectedPositionB := models.Coordinate{
			X: 11,
			Y: 15,
		}

		actualPosition := p.GetPaddle()

		assert.Equal(t, expectedPositionA, actualPosition.PositionA)
		assert.Equal(t, expectedPositionB, actualPosition.PositionB)
	})
}

func TestPaddle_MoveLeft(t *testing.T) {
	t.Run("Move paddle to left should decrease X position by 1", func(t *testing.T) {
		p := paddle.NewPaddle(16, 14, 8)
		expectedPositionA := models.Coordinate{
			X: 2,
			Y: 15,
		}
		expectedPositionB := models.Coordinate{
			X: 9,
			Y: 15,
		}

		p.MoveLeft()
		actualPosition := p.GetPaddle()

		assert.Equal(t, expectedPositionA, actualPosition.PositionA)
		assert.Equal(t, expectedPositionB, actualPosition.PositionB)
	})

	t.Run("Move paddle to far to left should never go out of bounds", func(t *testing.T) {
		p := paddle.NewPaddle(16, 20, 4)
		expectedPositionA := models.Coordinate{
			X: 0,
			Y: 15,
		}
		expectedPositionB := models.Coordinate{
			X: 3,
			Y: 15,
		}

		for i := 0; i < 20; i++ {
			p.MoveLeft()
		}
		actualPosition := p.GetPaddle()

		assert.Equal(t, expectedPositionA, actualPosition.PositionA)
		assert.Equal(t, expectedPositionB, actualPosition.PositionB)
	})
}

func TestPaddle_MoveRight(t *testing.T) {
	t.Run("Move paddle to right should increase X position by 1", func(t *testing.T) {
		p := paddle.NewPaddle(16, 20, 8)
		expectedPositionA := models.Coordinate{
			X: 7,
			Y: 15,
		}
		expectedPositionB := models.Coordinate{
			X: 14,
			Y: 15,
		}

		p.MoveRight()
		actualPosition := p.GetPaddle()

		assert.Equal(t, expectedPositionA, actualPosition.PositionA)
		assert.Equal(t, expectedPositionB, actualPosition.PositionB)
	})

	t.Run("Move paddle to far to right should never go out of bounds", func(t *testing.T) {
		p := paddle.NewPaddle(16, 20, 4)
		expectedPositionA := models.Coordinate{
			X: 16,
			Y: 15,
		}
		expectedPositionB := models.Coordinate{
			X: 19,
			Y: 15,
		}

		for i := 0; i < 20; i++ {
			p.MoveRight()
		}
		actualPosition := p.GetPaddle()

		assert.Equal(t, expectedPositionA, actualPosition.PositionA)
		assert.Equal(t, expectedPositionB, actualPosition.PositionB)
	})
}
