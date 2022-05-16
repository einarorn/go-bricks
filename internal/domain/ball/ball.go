package ball

import (
	"go-bricks/internal/models"
)

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
	basePosY := b.height - 2

	b.ball.PositionA.X = basePosX - 1
	b.ball.PositionA.Y = basePosY
	b.ball.PositionB.X = basePosX
	b.ball.PositionB.Y = basePosY
	b.direction = models.Coordinate{
		X: 0,
		Y: -1,
	}
}

func (b *ball) Move() {
	b.ball.PositionA.X += b.direction.X
	b.ball.PositionA.Y += b.direction.Y
	b.ball.PositionB.X += b.direction.X
	b.ball.PositionB.Y += b.direction.Y
}

func (b *ball) SliceLeft() {
	b.direction.X += -1
}

func (b *ball) SliceRight() {
	b.direction.X += 1
}

func (b *ball) BounceVertical() {
	b.direction.Y *= -1
}

func (b *ball) BounceHorizontal() {
	b.direction.X *= -1
}

func (b *ball) GetBall() models.Ball {
	return b.ball
}
