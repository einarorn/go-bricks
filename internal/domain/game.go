package domain

import (
	"go-bricks/internal/models"
	"go-bricks/internal/ports"
)

type Game struct {
	Height int
	Width  int
}

func New(height, width int) Game {
	return Game{
		Height: height,
		Width:  width,
	}
}

func (g Game) Start(gui ports.UserOutput) {
	status := models.GameStatus{
		Height: g.Height,
		Width:  g.Width,
	}
	gui.Draw(status)
}
