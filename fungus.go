package main

import (
	"math/rand"
)

type fungus struct {
	entity
	spreadCount int
}

func newFungus(world *world) *fungus {
	x, y := world.atWalkableTile()
	return &fungus{
		entity{ x: x,
			y: y,
			glyph: 'f',
			color: 0x4b,
			world: world,
			hp: 20,
			maxHp: 20,
		},
		0}
}

func (e fungus) Position() (x, y int) {
	return e.x, e.y
}

func (e fungus) Avatar() (glyph rune, color int) {
	return e.glyph, e.color
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
