package blocks_test

import (
	"fmt"
	"go-bricks/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"

	"go-bricks/internal/domain/blocks"
)

func TestBlocks_GetBlocks(t *testing.T) {
	t.Run("Initializing 2x2 blocks matrix should have certain coordinates", func(t *testing.T) {
		width := 4
		b := blocks.NewBlocks(2, 2, width)

		expectedIndex0PositionA := models.Coordinate{
			X: 0,
			Y: 0,
		}
		expectedIndex0PositionB := models.Coordinate{
			X: width,
			Y: 1,
		}
		expectedIndex1PositionA := models.Coordinate{
			X: width+2,
			Y: 0,
		}
		expectedIndex1PositionB := models.Coordinate{
			X: 2*width+2,
			Y: 1,
		}
		expectedIndex2PositionA := models.Coordinate{
			X: 0,
			Y: 2,
		}
		expectedIndex2PositionB := models.Coordinate{
			X: width,
			Y: 3,
		}
		expectedIndex3PositionA := models.Coordinate{
			X: width+2,
			Y: 2,
		}
		expectedIndex3PositionB := models.Coordinate{
			X: 2*width+2,
			Y: 3,
		}

		actual := b.GetBlocks()
		fmt.Println(actual)

		assert.Equal(t, expectedIndex0PositionA, actual[0].PositionA)
		assert.Equal(t, expectedIndex0PositionB, actual[0].PositionB)
		assert.Equal(t, expectedIndex1PositionA, actual[1].PositionA)
		assert.Equal(t, expectedIndex1PositionB, actual[1].PositionB)
		assert.Equal(t, expectedIndex2PositionA, actual[2].PositionA)
		assert.Equal(t, expectedIndex2PositionB, actual[2].PositionB)
		assert.Equal(t, expectedIndex3PositionA, actual[3].PositionA)
		assert.Equal(t, expectedIndex3PositionB, actual[3].PositionB)
	})
	
	t.Run("Initializing a blocks matrix should have have length of rows x columns", func(t *testing.T) {
		rows := 6
		columns := 6

		b := blocks.NewBlocks(rows, columns, 4)

		assert.Equal(t, rows*columns, len(b.GetBlocks()))
	})

	t.Run("Initialized blocks should have status IsDestroyed as false", func(t *testing.T) {
		b := blocks.NewBlocks(6, 6, 4)

		for _, block := range b.GetBlocks() {
			assert.False(t, block.IsDestroyed)
		}
	})
}
