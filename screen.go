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

	player := world.player
	playerX, playerY := viewport.worldToViewport(player.x, player.y)
	termbox.SetCell(playerX, playerY, player.glyph, termbox.Attribute(player.color), 0)
	s.Dump(0, 0, viewport)
	s.Dump(0, 1, world.player)
}

func (s playScreen) Input(game *game, event termbox.Event) []Drawable {
	viewport := game.viewport
	world := game.world
	player := world.player
	switch {
	case event.Ch == 'q':
		return []Drawable{welcomeScreen{}}
	case event.Ch == 'h':
		player.move(-1, 0)
	case event.Ch == 'j':
		player.move(0, 1)
	case event.Ch == 'k':
		player.move(0, -1)
	case event.Ch == 'l':
		player.move(1, 0)
	case event.Ch == 's':
		SmoothCave(world)
	}
 	viewport.center(player.x, player.y)
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
