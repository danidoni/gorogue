package main

import (
	"math/rand"
)

type entity struct {
	location *Point
	glyph rune
	color int
	world *world
	stats *Stats
	seed  *rand.Rand
	name string
}

func (e *entity) Avatar() (rune, int) {
	return e.glyph, e.color
}

func (e *entity) Position() *Point {
	return e.location
}

func (e *entity) update() {
	// Noop
}

func (e *entity) Stats() *Stats {
	return e.stats
}

type autonomous interface {
	Position() *Point
	Avatar() (rune, int)
	Name() string
	update()
	Stats() *Stats
}
