package paddle_test

import (
	"go-bricks/internal/domain/paddle"
	"go-bricks/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPaddle_GetPosition(t *testing.T) {
	t.Run("Initial position should be in the center and at bottom", func(t *testing.T) {
		p := paddle.NewPaddle(16, 20, 3)
		expectedPositionA := models.Coordinate{
			X: 9,
			Y: 15,
		}
		expectedPositionB := models.Coordinate{
			X: 11,
			Y: 15,
		}

		actualPosition := p.GetPosition()

		assert.Equal(t, expectedPositionA, actualPosition.PositionA)
		assert.Equal(t, expectedPositionB, actualPosition.PositionB)
	})
}

func TestPaddle_MoveLeft(t *testing.T) {
	t.Run("Move paddle to left should decrease X position by 1", func(t *testing.T) {
		p := paddle.NewPaddle(16, 20, 3)
		expectedPositionA := models.Coordinate{
			X: 8,
			Y: 15,
		}
		expectedPositionB := models.Coordinate{
			X: 10,
			Y: 15,
		}

		p.MoveLeft()
		actualPosition := p.GetPosition()

		assert.Equal(t, expectedPositionA, actualPosition.PositionA)
		assert.Equal(t, expectedPositionB, actualPosition.PositionB)
	})

	t.Run("Move paddle to far to left should never go out of bounds", func(t *testing.T) {
		p := paddle.NewPaddle(16, 20, 3)
		expectedPositionA := models.Coordinate{
			X: 0,
			Y: 15,
		}
		expectedPositionB := models.Coordinate{
			X: 2,
			Y: 15,
		}

		for i := 0; i < 20; i++ {
			p.MoveLeft()
		}
		actualPosition := p.GetPosition()

		assert.Equal(t, expectedPositionA, actualPosition.PositionA)
		assert.Equal(t, expectedPositionB, actualPosition.PositionB)
	})
}

func TestPaddle_MoveRight(t *testing.T) {
	t.Run("Move paddle to right should increase X position by 1", func(t *testing.T) {
		p := paddle.NewPaddle(16, 20, 3)
		expectedPositionA := models.Coordinate{
			X: 10,
			Y: 15,
		}
		expectedPositionB := models.Coordinate{
			X: 12,
			Y: 15,
		}

		p.MoveRight()
		actualPosition := p.GetPosition()

		assert.Equal(t, expectedPositionA, actualPosition.PositionA)
		assert.Equal(t, expectedPositionB, actualPosition.PositionB)
	})

	t.Run("Move paddle to far to right should never go out of bounds", func(t *testing.T) {
		p := paddle.NewPaddle(16, 20, 3)
		expectedPositionA := models.Coordinate{
			X: 17,
			Y: 15,
		}
		expectedPositionB := models.Coordinate{
			X: 19,
			Y: 15,
		}

		for i := 0; i < 20; i++ {
			p.MoveRight()
		}
		actualPosition := p.GetPosition()

		assert.Equal(t, expectedPositionA, actualPosition.PositionA)
		assert.Equal(t, expectedPositionB, actualPosition.PositionB)
	})
}
