package main

import (
	"github.com/nsf/termbox-go"
)

type GameoverScreen struct {
	screen
}

func (s GameoverScreen) Setup(game *game) {
	game.world = nil
}

func (s GameoverScreen) Draw(game *game) {
	s.Write(0, 0, "Game over!", 0, 0)
	s.Write(0, 1, "Press any key to go to the welcome screen.", 0, 0)
}

func (s GameoverScreen) Input(game *game, event termbox.Event) []Drawable {
	return []Drawable{welcomeScreen{}}
}
