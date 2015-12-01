package main

import (
	"github.com/nsf/termbox-go"
	"os"
	"log"
	"time"
)

func main() {
	file, err := os.OpenFile("gorogue.log", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", os.Stderr, ":", err)
	}
	defer file.Close()
	logger := log.New(file, "logger: ", log.Lshortfile)
	logger.Printf("%s: Gorogue started!", time.Now())

	err = termbox.Init()
	if err != nil {
		panic(err)
	}
	termbox.SetOutputMode(termbox.Output216)
	defer termbox.Close()
	termbox.HideCursor()
	termbox.Clear(termbox.ColorWhite, termbox.ColorBlack)
	game := NewGame(logger)
	game.Run()

}
