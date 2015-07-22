package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

type basicScreen struct{}

func (s basicScreen) Write(x, y int, message string, fg, bg termbox.Attribute) {
	for _, c := range message {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func (s basicScreen) Dump(x, y int, obj interface{}) {
	s.Write(x, y, fmt.Sprintf("%+v", obj), 0, 0)
}

type ScreenWriter interface {
	Write(x, y int, message string, fg, bg termbox.Attribute)
	Dump(x, y int, obj interface{})
}

type welcomeScreen struct {
	basicScreen
}

func (s welcomeScreen) Draw(world *world, viewport *viewport) {
	s.Write(0, 0, "Welcome to gorogue v0.0!", 0, 0)
	s.Write(0, 1, "Press space to go to the play screen.", 0, 0)
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

type playScreen struct {
	basicScreen
}

func (s playScreen) Draw(world *world, viewport *viewport) {
	for y := 0; y < viewport.height; y++ {
		for x := 0; x < viewport.width; x++ {
			tile := viewport.GetTile(x, y)
			termbox.SetCell(x, y,
				tile.glyph,
				termbox.Attribute(tile.color),
				termbox.ColorBlack)
		}
	}
	s.Dump(0, 0, viewport)
}

func (s playScreen) Input(game *game, event termbox.Event) []Drawable {
	viewport := game.viewport
	world := game.world
	switch {
	case event.Ch == 'q':
		return []Drawable{welcomeScreen{}}
	case event.Ch == 'h':
		viewport.Move(left, 1, world)
		return []Drawable{playScreen{}}
	case event.Ch == 'j':
		viewport.Move(down, 1, world)
		return []Drawable{playScreen{}}
	case event.Ch == 'k':
		viewport.Move(up, 1, world)
		return []Drawable{playScreen{}}
	case event.Ch == 'l':
		viewport.Move(right, 1, world)
		return []Drawable{playScreen{}}
	}
	return []Drawable{playScreen{}}
}

type dialogScreen struct {
	basicScreen
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

type Drawable interface {
	ScreenWriter
	Draw(world *world, viewport *viewport)
	Input(game *game, event termbox.Event) []Drawable
}
