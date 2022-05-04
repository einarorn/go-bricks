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
}

func drawBorder(s tcell.Screen, height, width int) {
	borderStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorGray)

	for w := 1 + 1; w < height; w++ {
		s.SetContent(1, w, tcell.RuneVLine, nil, borderStyle)
		s.SetContent(width, w, tcell.RuneVLine, nil, borderStyle)
	}

	for h := 1; h <= width; h++ {
		s.SetContent(h, 1, tcell.RuneHLine, nil, borderStyle)
		s.SetContent(h, height, tcell.RuneHLine, nil, borderStyle)
	}

	s.Show()
}


