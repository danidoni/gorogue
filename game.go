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
	world := NewWorld(250, 100)

	width, height := termbox.Size()
	viewport := centeredViewport(world.player.location, width, height, world)

	screens := Stack{}
	screens.Push(welcomeScreen{})
	game := &game{
		viewport: viewport,
		screens: screens,
		world: world,
		logger: logger,
	}
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

func (g *game) Update(screen Drawable) {
	event := termbox.PollEvent()
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
		g.Update(screen)
		screen = g.PopLastScreen()
	}
}
