package tcell

import (
	"github.com/gdamore/tcell/v2"
	"go-bricks/internal/models"
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
	drawBorder(gui.screen, status.Height, status.Width)
	drawText(gui.screen, 2, status.Height+2, status.Title)

	gui.screen.Show()
}

func drawText(s tcell.Screen,x int, y int, text string) {
	style := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
	for index, char := range text {
		s.SetContent(x+index, y, rune(char), nil, style)
	}
}

func drawBorder(s tcell.Screen, height, width int) {
	borderStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorGray)

	for w := 1 + 1; w < height; w++ {
		s.SetContent(1, w, ' ', nil, borderStyle)
		s.SetContent(width, w, ' ', nil, borderStyle)
	}

	for h := 1; h <= width; h++ {
		s.SetContent(h, 1, ' ', nil, borderStyle)
		s.SetContent(h, height, ' ', nil, borderStyle)
	}
}


