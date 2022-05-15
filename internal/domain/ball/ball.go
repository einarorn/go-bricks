package ball

import "go-bricks/internal/models"

type ball struct {
	ball      models.Ball
	height    int
	width     int
	direction models.Coordinate
}

func NewBall(height, width int) *ball {
	b := ball{
		height: height,
		width:  width,
	}
	b.inti()
	return &b
}

func (b *ball) inti() {
	basePosX := b.width / 2

	b.ball.Position.X = basePosX
	b.ball.Position.Y = b.height - 2
	b.direction = models.Coordinate{
		X: 0,
		Y: -1,
	}
}

func (b *ball) Move() {
	b.ball.Position.X += b.direction.X
	b.ball.Position.Y += b.direction.Y
}

func (b *ball) GetBall() models.Ball {
	return b.ball
}
