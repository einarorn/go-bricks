package blocks

import (
	"go-bricks/internal/models"
)

type blocks struct {
	blocks  []models.Block
	rows    int
	columns int
	width   int
}

func NewBlocks(rows, columns, width int) *blocks {
	b := blocks{
		blocks:  make([]models.Block, rows*columns),
		rows:    rows,
		columns: columns,
		width:   width,
	}
	b.init()
	return &b
}

func (b *blocks) init() {
	var x, y int
	color := 1
	for i := 0; i < b.rows*b.columns; i++ {
		b.blocks[i].PositionA = models.Coordinate{
			X: x,
			Y: y,
		}
		b.blocks[i].PositionB = models.Coordinate{
			X: x + b.width - 1,
			Y: y + 1,
		}
		b.blocks[i].Color = color

		x = x + b.width
		if b.rows*(b.width) <= x {
			x = 0
			y = y + 2
			color++
		}
	}
}

func (b *blocks) GetBlocks() []models.Block {
	return b.blocks
}
