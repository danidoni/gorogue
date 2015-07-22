package main

import (
	"github.com/nsf/termbox-go"
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
}

func NewGame() *game {
	screens := Stack{}
	screens.Push(welcomeScreen{})
	width, height := termbox.Size()
	world := NewWorld(250, 100)
	viewport := &viewport{0, 0, width, height, world}
	game := &game{viewport, screens, world}
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

func (g *game) Run() {
	screen := g.PopLastScreen()
	for screen != nil {
		termbox.Clear(termbox.ColorWhite, termbox.ColorBlack)
		screen.(Drawable).Draw(g.world, g.viewport)
		termbox.Flush()
		event := termbox.PollEvent()
		g.Update(screen, event)
		screen = g.PopLastScreen()
	}
}
