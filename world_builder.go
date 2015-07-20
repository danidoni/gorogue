package main

import (
	"math/rand"
)

func RandomWorld(width, height int) *world {
	world := &world{width: width, height: height}
	world.cells = make([][]*tile, height)
	for r := range world.cells {
		world.cells[r] = make([]*tile, width)
		for c := range world.cells[r] {
			if rand.Float32() < 0.5 {
				world.cells[r][c] = NewTile(floor)
			} else {
				world.cells[r][c] = NewTile(wall)
			}
		}
	}
	return world
}
