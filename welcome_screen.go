package main

import (
	"github.com/nsf/termbox-go"
)

type welcomeScreen struct {
	screen
}

func (s welcomeScreen) Setup(game *game) {

}

func (s welcomeScreen) Draw(game *game) {
	s.Write(0, 0, "Welcome to gorogue v0.0!", 0, 0)
	s.Write(0, 1, "Press space to create a new game.", 0, 0)
	s.Write(0, 2, "Press i to launch a dialog.", 0, 0)
	s.Write(0, 3, "Press q to exit.", 0, 0)
}

func (s welcomeScreen) Input(game *game, event termbox.Event) []Drawable {
	if event.Key == termbox.KeySpace {
		return []Drawable{playScreen{}}
	}
	if event.Ch == 'i' {
		return []Drawable{s, dialogScreen{}}
	}
	if event.Ch == 'q' {
		return []Drawable{}
	}
	return []Drawable{welcomeScreen{}}
}
