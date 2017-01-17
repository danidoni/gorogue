package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

type playScreen struct {
	screen
}

func (s playScreen) Setup(game *game) {
	if game.world == nil {
		game.world = NewWorld(250, 100)
		game.viewport = NewViewport(
			&Point{0, 0},
			game.width,
			game.height,
			game.world)
	}

	viewport := game.viewport
	location := game.world.player.location
	viewport.center(location.x, location.y)
}

func (s playScreen) Draw(game *game) {
	world := game.world
	viewport := game.viewport

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
	playerLocation := viewport.worldToViewport(player.location)
	termbox.SetCell(playerLocation.x,
		playerLocation.y,
		player.glyph,
		termbox.Attribute(player.color),
		0)

	// Dump stats at upper-right corner
	stats := player.Stats()
	s.Write(viewport.width-15, 0, fmt.Sprintf("HP: %d/%d", stats.Hp(), stats.MaxHp()), 0, 0)
	s.Write(viewport.width-15, 1, fmt.Sprintf("AT: %d", stats.Attack()), 0, 0)
	s.Write(viewport.width-15, 2, fmt.Sprintf("DF: %d", stats.Defense()), 0, 0)

	// Show notification messages
	j := 0
	height := viewport.height - 1
	world.notifications.each(func(i int, message string) {
		s.Write(0, height-j, message, 0, 0)
		j++
	})
}

func (s playScreen) Input(game *game, event termbox.Event) []Drawable {
	viewport := game.viewport
	world := game.world
	player := world.player
	switch {
	case event.Ch == 'q':
		return []Drawable{GameoverScreen{}}
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

	world.entities.each(func(e autonomous) {
		e.update()
	})
	return []Drawable{playScreen{}}
}
