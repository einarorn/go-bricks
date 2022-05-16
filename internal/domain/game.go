package domain

import (
	"go-bricks/internal/domain/ball"
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
	b := blocks.NewBlocks(9, 6, 8)
	p := paddle.NewPaddle(g.Height, g.Width, 10)
	bl := ball.NewBall(g.Height, g.Width)

	status := models.GameStatus{
		Title:  title,
		Height: g.Height,
		Width:  g.Width,
		Ball:   bl.GetBall(),
		Blocks: b.GetBlocks(),
		Paddle: p.GetPaddle(),
	}
	gui.Draw(status)

	/*for i := 0; i<15; i++ {
		time.Sleep(time.Millisecond * 50)

		bl.Move()
		status.Ball = bl.GetBall()
		gui.Draw(status)
	}*/
}
