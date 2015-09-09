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
	x, y := world.atWalkableTile()
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

type fungus struct {
	entity *entity
	spreadCount int
}

func newFungus(world *world) *fungus {
	x, y := world.atWalkableTile()
	return &fungus{&entity{x, y, 'f', 0x4b, world}, 0}
}

func (e fungus) Position() (x, y int) {
	return e.entity.x, e.entity.y
}

func (e fungus) Avatar() (glyph rune, color int) {
	return e.entity.glyph, e.entity.color
}

// Fungus entities don't move, they are stationary creatures
func (f *fungus) move(offsetX, offsetY int) {
}

func (f *fungus) update() {
	if f.spreadCount < 5 && rand.Float32() < 0.02 {
		f.spread()
	}
}

func (f *fungus) spread() *fungus {
	child := newFungus(f.entity.world)
	child.entity.x = f.entity.x + int(rand.Float32() * 11) - 5
	child.entity.y = f.entity.y + int(rand.Float32() * 11) - 5
	for f.entity.world.GetTile(child.entity.x, child.entity.y).isWalkable() == false {
		child.entity.x = f.entity.x + int(rand.Float32() * 11) - 5
		child.entity.y = f.entity.y + int(rand.Float32() * 11) - 5
	}
	f.entity.world.entities.PushBack(child)
	return child
}

type interactive interface {
	Position() (int, int)
	Avatar() (rune, int)
	update()
}
