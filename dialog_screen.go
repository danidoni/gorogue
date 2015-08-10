package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

type dialogScreen struct {
	screen
}

func (s dialogScreen) Draw(world *world, viewport *viewport) {
	s.Write(0, 0, "This is a palette dialog", 0, 0)
	s.Write(0, 1, "Press q to go back.", 0, 0)
	var color termbox.Attribute = 0x00
	for x := 0; x < 24; x++{
		for y := 0; y < 9; y++ {
			s.Write(15*y, 3+x, fmt.Sprintf("%#x", color), 0, 0)
			s.Write(15*y+5, 3+x, "   ", 0, color)
			color++
		}
	}
}

func (s dialogScreen) Input(game *game, event termbox.Event) []Drawable {
	if event.Ch == 'q' {
		return []Drawable{}
	}
	return []Drawable{dialogScreen{}}
}
