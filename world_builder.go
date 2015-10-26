package main

import (
	"math/rand"
	"time"
)

func RandomWorld(world *world) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	RandomCave(world, r)

	for i := 0; i < 8; i++ {
		SmoothCave(world)
	}
}

func RandomCave(world *world, r *rand.Rand) {
	for row := range world.cells {
		for col := range world.cells[row] {
			if r.Float32() < 0.5 {
				world.cells[row][col] = NewTile(floor)
			} else {
				world.cells[row][col] = NewTile(wall)
			}
		}
	}
}

func SmoothCave(world *world) {
	var smoothedCells [][]*tile = make([][]*tile, world.dimensions.y)

	for row := range world.cells {
		smoothedCells[row] = make([]*tile, world.dimensions.x)
		for col := range world.cells[row] {
			floors := 0
			walls := 0

			for y := -1; y < 2; y++ {
				for x := -1; x < 2; x++ {
					point := &Point{
						x: x+col,
						y: y+row,
					}
					if world.GetTile(point).kind == floor {
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
