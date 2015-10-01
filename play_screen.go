package main

import (
	"github.com/nsf/termbox-go"
	"fmt"
)

type playScreen struct {
	screen
}

func (s playScreen) Draw(world *world, viewport *viewport) {
	// Draws all viewport elements
	viewport.iterate(func(x, y int, tile *tile) {
		termbox.SetCell(x, y,
			tile.glyph,
			termbox.Attribute(tile.color),
			termbox.ColorBlack)
	})

	// Draws all entities visible in the viewport
	viewport.entities(func(entity autonomous) {
		location := entity.Position()
		viewportLocation := viewport.worldToViewport(location)
		glyph, color := entity.Avatar()
		termbox.SetCell(viewportLocation.x,
			viewportLocation.y,
			glyph,
			termbox.Attribute(color),
			termbox.ColorBlack)
	})

	// Draws the player
	player := world.player
	playerLocation := viewport.worldToViewport(player.location)
	termbox.SetCell(playerLocation.x,
		playerLocation.y,
		player.glyph,
		termbox.Attribute(player.color),
		0)

	// Dump stats at upper-right corner
	s.Write(viewport.width - 15, 0, fmt.Sprintf("HP: %d/%d", player.Hp(), player.MaxHp()), 0, 0)
}

func (s playScreen) Input(game *game, event termbox.Event) []Drawable {
	viewport := game.viewport
	world := game.world
	player := world.player
	switch {
	case event.Ch == 'q':
		return []Drawable{welcomeScreen{}}
	case event.Ch == 'h':
		player.move(&Point{x: -1, y: 0})
	case event.Ch == 'y':
		player.move(&Point{x: -1, y: -1})
	case event.Ch == 'k':
		player.move(&Point{x: 0, y: -1})
	case event.Ch == 'u':
		player.move(&Point{x: 1, y: -1})
	case event.Ch == 'l':
		player.move(&Point{x: 1, y: 0})
	case event.Ch == 'n':
		player.move(&Point{x: 1, y: 1})
	case event.Ch == 'j':
		player.move(&Point{x: 0, y: 1})
	case event.Ch == 'b':
		player.move(&Point{x: -1, y: 1})
	case event.Ch == 's':
		SmoothCave(world)
	}
 	viewport.center(player.location.x, player.location.y)
	for e := world.entities.Front(); e != nil; e = e.Next() {
		e.Value.(autonomous).update()
	}
	return []Drawable{playScreen{}}
}
