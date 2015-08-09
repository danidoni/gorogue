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
	x := rand.Intn(world.width)
	y := rand.Intn(world.height)
	for world.GetTile(x, y).isWalkable() == false {
		x = rand.Intn(world.width)
		y = rand.Intn(world.height)
	}
	return &player{x, y, '@', 0, world}
}

func (p *player) move(direction direction, step int) {
	switch direction {
	case left:
		updatedX := p.x - step
		if p.world.GetTile(updatedX, p.y).isWalkable() {
			p.x = updatedX
		}
	case right:
		updatedX := p.x + step
		if p.world.GetTile(updatedX, p.y).isWalkable() {
			p.x = updatedX
		}
	case up:
		updatedY := p.y - step
		if p.world.GetTile(p.x, updatedY).isWalkable() {
			p.y = updatedY
		}
	case down:
		updatedY := p.y + step
		if p.world.GetTile(p.x, updatedY).isWalkable() {
			p.y = updatedY
		}
	}
}
