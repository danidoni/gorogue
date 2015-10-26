package main

import (
	"math/rand"
	"time"
)

type fungus struct {
	entity
	spreadCount int
}

func newFungus(world *world) *fungus {
	point := world.atWalkableTile()
	return &fungus{
		entity{
			location: point,
			glyph: 'f',
			color: 0x4b,
			world: world,
			hp: 20,
			maxHp: 20,
			seed: rand.New(rand.NewSource(time.Now().UnixNano())),
		},
		0}
}

func (e fungus) Position() *Point {
	return e.location
}

func (e fungus) Avatar() (glyph rune, color int) {
	return e.glyph, e.color
}

// Fungus entities don't move, they are stationary creatures
func (f *fungus) move(offset *Point) {
}

func (f *fungus) update() {
	if f.spreadCount < 5 && rand.Float32() < 0.02 {
		f.spread()
	}
}

func (f *fungus) spread() *fungus {
	child := newFungus(f.entity.world)
	randomPoint := &Point{
		x: int(rand.Float32() * 11) - 5,
		y: int(rand.Float32() * 11) - 5,
	}
	child.entity.location.Add(randomPoint)
	for f.entity.world.GetTile(child.entity.location).isWalkable() == false {
		randomPoint := &Point{
			x: int(rand.Float32() * 11) - 5,
			y: int(rand.Float32() * 11) - 5,
		}
		child.entity.location.Add(randomPoint)
	}
	f.entity.world.entities.PushBack(child)
	return child
}
