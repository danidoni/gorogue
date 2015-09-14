package main

import (
	"math/rand"
)

func RandomWorld(world *world) {
	RandomCave(world)

	for i := 0; i < 8; i++ {
		SmoothCave(world)
	}
}

func RandomCave(world *world) {
	for row := range world.cells {
		for col := range world.cells[row] {
			if rand.Float32() < 0.5 {
				world.cells[row][col] = NewTile(floor)
			} else {
				world.cells[row][col] = NewTile(wall)
			}
		}
	}
}

func SmoothCave(world *world) {
	var smoothedCells [][]*tile = make([][]*tile, world.height)

	for row := range world.cells {
		smoothedCells[row] = make([]*tile, world.width)
		for col := range world.cells[row] {
			floors := 0
			walls := 0

			for y := -1; y < 2; y++ {
				for x := -1; x < 2; x++ {
					if world.GetTile(x+col, y+row).kind == floor {
						floors++
					} else {
						walls++
					}
				}
			}

			if floors >= walls {
				smoothedCells[row][col] = NewTile(floor)
			} else {
				smoothedCells[row][col] = NewTile(wall)
			}
		}
	}
	world.cells = smoothedCells
}
