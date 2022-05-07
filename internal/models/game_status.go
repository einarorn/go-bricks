package models

type (
	GameStatus struct {
		Title      string
		Height     int
		Width      int
		BlockWidth int
		Blocks     []Block
	}

	Block struct {
		PositionA   Coordinate
		PositionB   Coordinate
		IsDestroyed bool
	}

	Coordinate struct {
		X, Y int
	}
)
