package domain

import "fmt"

type Game struct {
	height int
	width  int
}

func New(height, width int) Game {
	return Game{
		height: height,
		width:  width,
	}
}

func (g Game) Start() {
	fmt.Printf("Game started... Height: %d, Width: %d", g.height, g.width)
}
