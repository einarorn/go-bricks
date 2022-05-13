package domain

import (
	"go-bricks/internal/domain/blocks"
	"go-bricks/internal/domain/paddle"
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
	b := blocks.NewBlocks(8, 6, 8)
	p := paddle.NewPaddle(g.Height, g.Width, 3)

	status := models.GameStatus{
		Title:  title,
		Height: g.Height,
		Width:  g.Width,
		Blocks: b.GetBlocks(),
		Paddle: p.GetPosition(),
	}
	gui.Draw(status)
}
