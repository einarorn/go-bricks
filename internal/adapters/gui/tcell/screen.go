package tcell

import (
	"fmt"
	"go-bricks/internal/models"

	"github.com/gdamore/tcell/v2"
)

const (
	borderWidth = 2
	offsetX     = 2
	offsetY     = 2

	blockColor1 = tcell.ColorRed
	blockColor2 = tcell.ColorYellow
	blockColor3 = tcell.ColorGreen
	blockColor4 = tcell.ColorBlue
	blockColor5 = tcell.ColorWhite
	blockColor6 = tcell.ColorPurple
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

func (gui GraphicalUserInterface) drawBall(ball models.Ball)  {
	style := tcell.StyleDefault.Foreground(tcell.ColorDefault).Background(tcell.ColorDefault)
	gui.screen.SetContent(ball.Position.X+offsetX+borderWidth, ball.Position.Y+offsetY+borderWidth, '⚪', nil, style)

	//gui.drawText(offsetX, 28+offsetY+borderWidth+4, fmt.Sprintf("X:%v Y:%v", ball.Position.X, ball.Position.Y))
}

func (gui GraphicalUserInterface) drawPaddle(paddle models.Paddle) {
	style := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)

	for i := paddle.PositionA.X + offsetX + borderWidth; i <= paddle.PositionB.X+offsetX+borderWidth; i++ {
		gui.screen.SetContent(i, paddle.PositionA.Y+offsetY+borderWidth, '█', nil, style)
	}
}

func (gui *GraphicalUserInterface) drawBlocks(blocks []models.Block) {
	style := tcell.StyleDefault.Foreground(tcell.ColorDefault).Background(tcell.ColorBlack)

	for _, block := range blocks {
		gui.drawSingleBlock(block.PositionA.X+offsetX+borderWidth, block.PositionA.Y+offsetY+borderWidth, block.PositionB.X+offsetX+borderWidth, block.Color, style)
	}
}

func (gui *GraphicalUserInterface) drawSingleBlock(x, y, width, color int, style tcell.Style) {
	switch color {
	case 1:
		style = style.Foreground(blockColor1)
	case 2:
		style = style.Foreground(blockColor2)
	case 3:
		style = style.Foreground(blockColor3)
	case 4:
		style = style.Foreground(blockColor4)
	case 5:
		style = style.Foreground(blockColor5)
	case 6:
		style = style.Foreground(blockColor6)
	}

	for i := x; i <= width; i++ {
		gui.screen.SetContent(i, y, '█', nil, style)
		gui.screen.SetContent(i, y+1, '▀', nil, style)
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

	for y := offsetY+1; y <=height+offsetY+1; y++ {
		for x := offsetX+1; x <= width+1; x++ {
			gui.screen.SetContent(x, y, '█', nil, borderStyle)
		}
	}
}

func (gui *GraphicalUserInterface) drawBorder(height, width int) {
	borderStyle := tcell.StyleDefault.Foreground(tcell.ColorReset).Background(tcell.ColorGray)

	for x := offsetX; x <= width+offsetX; x++ {
		gui.screen.SetContent(x, offsetY, ' ', nil, borderStyle)
		gui.screen.SetContent(x, height+offsetY+borderWidth, ' ', nil, borderStyle)
	}

	for y := offsetY; y < height+offsetY+borderWidth; y++ {
		gui.screen.SetContent(offsetX, y, ' ', nil, borderStyle)
		gui.screen.SetContent(width+offsetX, y, ' ', nil, borderStyle)
	}
}

func (gui GraphicalUserInterface) drawYAxis(height, width int) {
	offset := offsetY + borderWidth
	for i := 0; i < height+offset; i++ {
		gui.drawText(width+offsetX+borderWidth, i, fmt.Sprintf("%v", i-offset))
	}
}
