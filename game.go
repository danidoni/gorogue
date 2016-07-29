package main

import (
	"github.com/nsf/termbox-go"
	"log"
)

type direction int

const (
	up direction = iota
	down
	left
	right
)

type game struct {
	viewport *viewport
	screens  Stack
	world    *world
	logger   *log.Logger
}

func NewGame(logger *log.Logger) *game {
	screens := Stack{}
	screens.Push(welcomeScreen{})
	width, height := termbox.Size()
	world := NewWorld(250, 100)
	player := newPlayer(world)
	viewport := centeredViewport(player.location, width, height, world)
	world.player = player
	var entity autonomous = newFungus(world)
	world.entities.add(entity)
	game := &game{viewport, screens, world, logger}
	return game
}

func (g *game) PopLastScreen() Drawable {
	screen := g.screens.Pop()
	if screen != nil {
		return screen.(Drawable)
	} else {
		return nil
	}
}

func (g *game) Update(screen Drawable, event termbox.Event) {
	nextScreens := screen.Input(g, event)
	if len(nextScreens) > 0 {
		for _, s := range nextScreens {
			g.screens.Push(s)
		}
	}
}

func (g *game) Render(screen Drawable) {
	termbox.Clear(termbox.ColorWhite, termbox.ColorBlack)
	screen.(Drawable).Draw(g)
	termbox.Flush()
}

func (g *game) Run() {
	screen := g.PopLastScreen()
	for screen != nil {
		g.Render(screen)
		event := termbox.PollEvent()
		g.Update(screen, event)
		screen = g.PopLastScreen()
	}
}
