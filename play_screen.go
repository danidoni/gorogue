package main

import (
	"github.com/nsf/termbox-go"
)

type playScreen struct {
	screen
}

func (s playScreen) Draw(world *world, viewport *viewport) {
	viewport.iterate(func(x, y int, tile *tile) {
		termbox.SetCell(x, y,
			tile.glyph,
			termbox.Attribute(tile.color),
			termbox.ColorBlack)
	})

	viewport.entities(func(entity interactive) {
		x, y := entity.Position()
		translatedX, translatedY := viewport.worldToViewport(x, y)
		glyph, color := entity.Avatar()
		termbox.SetCell(translatedX,
			translatedY,
			glyph,
			termbox.Attribute(color),
			termbox.ColorBlack)
	})

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
	case event.Ch == 'y':
		player.move(-1, -1)
	case event.Ch == 'k':
		player.move(0, -1)
	case event.Ch == 'u':
		player.move(1, -1)
	case event.Ch == 'l':
		player.move(1, 0)
	case event.Ch == 'n':
		player.move(1, 1)
	case event.Ch == 'j':
		player.move(0, 1)
	case event.Ch == 'b':
		player.move(-1, 1)
	case event.Ch == 's':
		SmoothCave(world)
	}
 	viewport.center(player.x, player.y)
	for e := world.entities.Front(); e != nil; e = e.Next() {
		e.Value.(interactive).update()
	}
	return []Drawable{playScreen{}}
}
