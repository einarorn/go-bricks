package models

type (
	GameStatus struct {
		Title  string
		Height int
		Width  int
		Ball   Ball
		Blocks []Block
		Paddle Paddle
	}

	Ball struct {
		Position Coordinate
	}

	Block struct {
		PositionA   Coordinate
		PositionB   Coordinate
		IsDestroyed bool
		Color       int
	}

	Paddle struct {
		PositionA Coordinate
		PositionB Coordinate
	}

	Coordinate struct {
		X, Y int
	}
)
