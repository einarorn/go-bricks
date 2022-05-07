package domain

import (
	"go-bricks/internal/domain/blocks"
	"go-bricks/internal/models"
	"go-bricks/internal/ports"
)

const title = "Go Bricks!"

type Game struct {
	Height int
	Width  int
}

func NewGame(height, width int) Game {
	return Game{
		Height: height,
		Width:  width,
	}
}

func (g Game) Start(gui ports.UserOutput) {
	b := blocks.NewBlocks(10, 6, 6)

	status := models.GameStatus{
		Title: title,
		Height: g.Height,
		Width:  g.Width,
		BlockWidth: 8,
		Blocks: b.GetBlocks(),
	}
	gui.Draw(status)
}
