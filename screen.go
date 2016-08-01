package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

type screen struct{}

func (s screen) Write(x, y int, message string, fg, bg termbox.Attribute) {
	for _, c := range message {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func (s screen) Dump(x, y int, obj interface{}) {
	s.Write(x, y, fmt.Sprintf("%+v", obj), 0, 0)
}

type ScreenWriter interface {
	Write(x, y int, message string, fg, bg termbox.Attribute)
	Dump(x, y int, obj interface{})
}

type Drawable interface {
	ScreenWriter
	Setup(game *game)
	Draw(game *game)
	Input(game *game, event termbox.Event) []Drawable
}
