package main

import (
	"math/rand"
)

type entity struct {
	x, y  int
	glyph rune
	color int
	world *world
}

type player entity

func newPlayer(world *world) *player {
	x, y := atWalkableTile(world)
	return &player{x, y, '@', 0, world}
}

func (p *player) move(offsetX, offsetY int) {
	newX := p.x + offsetX
	newY := p.y + offsetY
	tile := p.world.GetTile(newX, newY)
	if tile.isWalkable() {
		p.x = newX
		p.y = newY
	} else if tile.isDiggable() {
		p.world.dig(newX, newY)
	}
}

type fungus entity

func newFungus(world *world) *fungus {
	x, y := atWalkableTile(world)
	return &fungus{x, y, 'f', 0x4b, world}
}

func (e fungus) Position() (x, y int) {
	return e.x, e.y
}

func (e fungus) Avatar() (glyph rune, color int) {
	return e.glyph, e.color
}

// Finds a random walkable tile in the world
func atWalkableTile(world *world) (x, y int) {
	x = rand.Intn(world.width)
	y = rand.Intn(world.height)
	for world.GetTile(x, y).isWalkable() == false {
		x = rand.Intn(world.width)
		y = rand.Intn(world.height)
	}
	return
}

// Fungus entities don't move, they are stationary creatures
func (f *fungus) move(offsetX, offsetY int) {
}

type renderable interface {
	Position() (int, int)
	Avatar() (rune, int)
}
