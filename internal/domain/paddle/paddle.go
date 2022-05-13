package paddle

import "go-bricks/internal/models"

type paddle struct {
	paddle models.Paddle
	height int
	width  int
	size   int
}

func NewPaddle(height, width, size int) *paddle {
	p := paddle{
		height: height,
		width:  width,
		size:   size,
	}
	p.init()
	return &p
}

func (p *paddle) init() {
	basePosX := p.width / 2
	sizeOffset := (p.size - 1) / 2

	p.paddle.PositionA.X = basePosX - sizeOffset
	p.paddle.PositionA.Y = p.height - 1
	p.paddle.PositionB.X = basePosX + sizeOffset
	p.paddle.PositionB.Y = p.height - 1
}

func (p *paddle) GetPosition() models.Paddle {
	return p.paddle
}

func (p *paddle) MoveLeft() {
	if p.paddle.PositionA.X > 0 {
		p.paddle.PositionA.X--
		p.paddle.PositionB.X--
	}
}

func (p *paddle) MoveRight() {
	if p.paddle.PositionB.X < p.width-1 {
		p.paddle.PositionA.X++
		p.paddle.PositionB.X++
	}
}
