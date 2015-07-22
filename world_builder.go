package main

import (
	"math/rand"
)

func RandomWorld(width, height int) *world {
	world := &world{width: width, height: height}
	world.cells = RandomCave(height, width)

	for i := 0; i < 8; i++ {
		SmoothCave(world)
	}

	return world
}

func RandomCave(height, width int) [][]*tile {
	cells := make([][]*tile, height)
	for row := range cells {
		cells[row] = make([]*tile, width)
		for col := range cells[row] {
			if rand.Float32() < 0.5 {
				cells[row][col] = NewTile(floor)
			} else {
				cells[row][col] = NewTile(wall)
			}
		}
	}

	return cells
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
