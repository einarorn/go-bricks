package main

import (
	game "go-bricks/internal/domain"
)

type App struct {
}

func newApp() *App {
	return &App{}
}

func (a App) start() {
	game.New(50, 50).Start()
}
