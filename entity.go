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

func (p *player) move(offsetX, offsetY int) {
	if p.world.GetTile(p.x + offsetX, p.y + offsetY).isWalkable() {
		p.x += offsetX
		p.y += offsetY
	}
}
