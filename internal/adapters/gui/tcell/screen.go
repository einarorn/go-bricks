package tcell

import (
	"go-bricks/internal/models"

	"github.com/gdamore/tcell/v2"
)

const (
	borderWidth = 2
	offsetX     = 2
	offsetY     = 2
)

type GraphicalUserInterface struct {
	screen tcell.Screen
}

func New() (*GraphicalUserInterface, error) {
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

func (gui GraphicalUserInterface) Draw(status models.GameStatus) {
	gui.drawBorder(status.Height, status.Width)
	gui.drawText(offsetX, status.Height+offsetY+2, status.Title)
	gui.drawBlocks(status.Blocks)

	gui.screen.Show()
}

func (gui GraphicalUserInterface) drawBlocks(blocks []models.Block) {
	style := tcell.StyleDefault.Foreground(tcell.ColorGreen).Background(tcell.ColorBlack).Bold(true)

	for _, block := range blocks {
		gui.drawSingleBlock(block.PositionA.X+offsetX+borderWidth, block.PositionA.Y+offsetY+borderWidth, block.PositionB.X+offsetX+borderWidth, style)
	}
}

func (gui GraphicalUserInterface) drawSingleBlock(x, y, width int, style tcell.Style) {
	for i := x; i <= width; i++ {
		gui.screen.SetContent(i, y, '█', nil, style)
		gui.screen.SetContent(i, y+1, '▀', nil, style)
	}
}

func (gui GraphicalUserInterface) drawText(x int, y int, text string) {
	style := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
	for index, char := range text {
		gui.screen.SetContent(x+index, y, rune(char), nil, style)
	}
}

func (gui GraphicalUserInterface) drawBorder(height, width int) {
	borderStyle := tcell.StyleDefault.Foreground(tcell.ColorReset).Background(tcell.ColorGray)

	for x := offsetX; x <= width+offsetX; x++ {
		gui.screen.SetContent(x, offsetY, ' ', nil, borderStyle)
		gui.screen.SetContent(x, height+offsetY, ' ', nil, borderStyle)
	}

	for y := offsetY; y < height+offsetY; y++ {
		gui.screen.SetContent(offsetX, y, ' ', nil, borderStyle)
		gui.screen.SetContent(width+offsetX, y, ' ', nil, borderStyle)
	}
}
