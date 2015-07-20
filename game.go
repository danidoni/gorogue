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

type viewport struct {
	x, y          int
	width, height int
}

type game struct {
	viewport *viewport
	screens  Stack
	world    *world
}

func NewGame() *game {
	screens := Stack{}
	screens.Push(welcomeScreen{})
	width, height := termbox.Size()
	viewport := &viewport{0, 0, width, height}
	world := NewWorld(250, 100)
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

func (v *viewport) Move(direction direction, step int, w *world) {
	switch {
	case direction == left:
		updatedX := v.x - step
		if updatedX >= 0 {
			v.x = updatedX
		}
	case direction == right:
		updatedX := v.x + step
		if updatedX+v.width <= w.width {
			v.x = updatedX
		}
	case direction == up:
		updatedY := v.y - step
		if updatedY >= 0 {
			v.y = updatedY
		}
	case direction == down:
		updatedY := v.y + step
		if updatedY+v.height <= w.height {
			v.y = updatedY
		}
	}
}
