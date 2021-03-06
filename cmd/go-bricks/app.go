package main

import (
	"go-bricks/internal/adapters/gui/tcell"
	game "go-bricks/internal/domain"
	"go-bricks/internal/ports"
)

type App struct {
	UserInterface ports.UserOutput
}

func newApp() (App, error) {
	tcellUI, err := tcell.NewGUI()
	if err != nil {
		return App{}, err
	}
	return App{UserInterface: tcellUI}, nil
}

func (a App) start() {
	game.NewGame(28, 90).Start(a.UserInterface)
}
