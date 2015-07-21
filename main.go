package main

import (
	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	termbox.SetOutputMode(termbox.Output216)
	defer termbox.Close()
	termbox.HideCursor()
	termbox.Clear(termbox.ColorWhite, termbox.ColorBlack)
	game := NewGame()
	game.Run()

}
