package tcell

import (
	"fmt"
	"go-bricks/internal/models"

	"github.com/gdamore/tcell/v2"
)

const (
	borderWidth = 1
	offsetX     = 2
	offsetY     = 2
)

var (
	blockColor1Even = tcell.NewRGBColor(231, 72, 86)
	blockColor1Odd  = tcell.NewRGBColor(201, 64, 78)
	blockColor2Even = tcell.NewRGBColor(255, 212, 0)
	blockColor2Odd  = tcell.NewRGBColor(221, 181, 0)
	blockColor3Even = tcell.NewRGBColor(26, 198, 17)
	blockColor3Odd  = tcell.NewRGBColor(23, 170, 15)
	blockColor4Even = tcell.NewRGBColor(38, 106, 255)
	blockColor4Odd  = tcell.NewRGBColor(32, 91, 219)
	blockColor5Even = tcell.NewRGBColor(255, 106, 0)
	blockColor5Odd  = tcell.NewRGBColor(211, 84, 0)
	blockColor6Even = tcell.NewRGBColor(197, 35, 221)
	blockColor6Odd  = tcell.NewRGBColor(169, 30, 191)
)

type GraphicalUserInterface struct {
	screen tcell.Screen
}

func NewGUI() (*GraphicalUserInterface, error) {
	s, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}
	if err := s.Init(); err != nil {
		return nil, err
	}

	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	s.SetStyle(defStyle)
	s.Clear()

	return &GraphicalUserInterface{screen: s}, nil
}

func (gui *GraphicalUserInterface) Draw(status models.GameStatus) {
	gui.drawBackground(status.Height, status.Width)
	gui.drawBorder(status.Height, status.Width)
	gui.drawText(offsetX, status.Height+offsetY+borderWidth+2, status.Title)
	gui.drawBlocks(status.Blocks)
	gui.drawPaddle(status.Paddle)
	gui.drawBall(status.Ball)

	gui.drawYAxis(status.Height, status.Width)

	gui.screen.Sync()
}

func (gui GraphicalUserInterface) drawBall(ball models.Ball) {
	style := tcell.StyleDefault.Foreground(tcell.ColorDefault).Background(tcell.ColorDefault)
	gui.screen.SetContent(ball.PositionA.X+offsetX+borderWidth, ball.PositionA.Y+offsetY+borderWidth, '⚪', nil, style)
	// gui.screen.SetContent(ball.PositionB.X+offsetX+borderWidth, ball.PositionB.Y+offsetY+borderWidth, '⚪', nil, style) //symbol takes two spaces

	// gui.drawText(offsetX, 28+offsetY+borderWidth+4, fmt.Sprintf("X:%v Y:%v", ball.Position.X, ball.Position.Y))
}

func (gui GraphicalUserInterface) drawPaddle(paddle models.Paddle) {
	style := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)

	for i := paddle.PositionA.X + offsetX + borderWidth; i <= paddle.PositionB.X+offsetX+borderWidth; i++ {
		gui.screen.SetContent(i, paddle.PositionA.Y+offsetY+borderWidth, '█', nil, style)
	}

	gui.drawText(offsetX, 28+offsetY+borderWidth+4, fmt.Sprintf("X:%v Y:%v", paddle.PositionA.X, paddle.PositionA.Y))
	gui.drawText(offsetX+20, 28+offsetY+borderWidth+4, fmt.Sprintf("X:%v Y:%v", paddle.PositionB.X, paddle.PositionB.Y))
}

func (gui *GraphicalUserInterface) drawBlocks(blocks []models.Block) {
	for index, block := range blocks {
		gui.drawSingleBlock(block.PositionA.X+offsetX+borderWidth, block.PositionA.Y+offsetY+borderWidth, block.PositionB.X+offsetX+borderWidth, block.Color, index)
	}
}

func (gui *GraphicalUserInterface) drawSingleBlock(x, y, width, color, blockIndex int) {
	style := tcell.StyleDefault.Foreground(tcell.ColorDefault).Background(tcell.ColorBlack)
	oddEven := blockIndex % 2

	switch color {
	case 1:
		if oddEven == 0 {
			style = style.Foreground(blockColor1Even)
		} else {
			style = style.Foreground(blockColor1Odd)
		}
	case 2:
		if oddEven == 0 {
			style = style.Foreground(blockColor2Even)
		} else {
			style = style.Foreground(blockColor2Odd)
		}
	case 3:
		if oddEven == 0 {
			style = style.Foreground(blockColor3Even)
		} else {
			style = style.Foreground(blockColor3Odd)
		}
	case 4:
		if oddEven == 0 {
			style = style.Foreground(blockColor4Even)
		} else {
			style = style.Foreground(blockColor4Odd)
		}
	case 5:
		if oddEven == 0 {
			style = style.Foreground(blockColor5Even)
		} else {
			style = style.Foreground(blockColor5Odd)
		}
	case 6:
		if oddEven == 0 {
			style = style.Foreground(blockColor6Even)
		} else {
			style = style.Foreground(blockColor6Odd)
		}
	}

	for i := x; i <= width; i++ {
		gui.screen.SetContent(i, y, '█', nil, style)
		gui.screen.SetContent(i, y+1, '█', nil, style)
	}
}

func (gui *GraphicalUserInterface) drawText(x int, y int, text string) {
	style := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
	for index, char := range text {
		gui.screen.SetContent(x+index, y, rune(char), nil, style)
	}
}

func (gui *GraphicalUserInterface) drawBackground(height, width int) {
	borderStyle := tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorBlack)

	for y := offsetY + 1; y <= height+offsetY+1; y++ {
		for x := offsetX + 1; x <= width+1; x++ {
			gui.screen.SetContent(x, y, '█', nil, borderStyle)
		}
	}
}

func (gui *GraphicalUserInterface) drawBorder(height, width int) {
	borderStyle := tcell.StyleDefault.Foreground(tcell.ColorReset).Background(tcell.ColorGray)

	for x := offsetX; x <= width+offsetX+borderWidth; x++ {
		gui.screen.SetContent(x, offsetY, ' ', nil, borderStyle)
		gui.screen.SetContent(x, height+offsetY+borderWidth, ' ', nil, borderStyle)
	}

	for y := offsetY; y <= height+offsetY+borderWidth; y++ {
		gui.screen.SetContent(offsetX, y, ' ', nil, borderStyle)
		gui.screen.SetContent(width+offsetX+borderWidth, y, ' ', nil, borderStyle)
	}
}

func (gui GraphicalUserInterface) drawYAxis(height, width int) {
	offset := offsetY + borderWidth
	for i := 0; i < height+offsetY+borderWidth; i++ {
		gui.drawText(width+offsetX+borderWidth+2, i, fmt.Sprintf("%v", i-offset))
	}
}
