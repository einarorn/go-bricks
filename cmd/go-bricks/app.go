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
	tcellUI, err := tcell.New()
	if err != nil {
		return App{}, err
	}
	return App{UserInterface: tcellUI}, nil
}

func (a App) start() {
	game.New(24, 100).Start(a.UserInterface)
}
