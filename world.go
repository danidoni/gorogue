package main

import (
	"container/list"
	"math/rand"
)

type tileType int

const (
	floor tileType = iota
	wall
	boundary
)

type tile struct {
	glyph rune
	kind  tileType
	color int
}

func (t tile) isWalkable() bool {
	return t.kind == floor
}

func (t tile) isDiggable() bool {
	return t.kind == wall
}

type world struct {
	width, height int
	cells         [][]*tile
	player        *player
	entities      *list.List
}

func NewTile(kind tileType) *tile {
	switch {
	case kind == floor:
		return &tile{'.', floor, 0x7}
	case kind == wall:
		return &tile{'#', wall, 0x55}
	}
	return &tile{'X', boundary, 0x91}
}

func NewWorld(width, height int) *world {
	world := &world{width: width, height: height}
	world.cells = make([][]*tile, world.height)
	for row := range world.cells {
		world.cells[row] = make([]*tile, world.width)
		for col := range world.cells[row] {
			world.cells[row][col] = NewTile(floor)
		}
	}
	RandomWorld(world)
	return world
}

func (w world) isWithinBoundaries(x, y int) bool {
	switch {
	case x < 0:
		return false
	case x >= w.width:
		return false
	case y < 0:
		return false
	case y >= w.height:
		return false
	}
	return true
}

func (w world) GetTile(x, y int) *tile {
	if w.isWithinBoundaries(x, y) {
		return w.cells[y][x]
	}
	return NewTile(boundary)
}

func (w *world) SetTile(x, y int, tile *tile) {
	if w.isWithinBoundaries(x, y) {
		w.cells[y][x] = tile
	}
}

func (w *world) dig(x, y int) {
	w.SetTile(x, y, NewTile(floor))
}

func (w *world) entitiesInside(x, y, width, height int, callback func(entity autonomous)) {
	for e := w.entities.Front(); e != nil; e = e.Next() {
		entity := e.Value.(autonomous)
		x, y := entity.Position()
		if x >= x &&
			x <= x+width &&
			y >= y &&
			y <= y+height {
			callback(entity)
		}
	}
}

// Finds a random walkable tile in the world
func (w *world) atWalkableTile() (x, y int) {
	x = rand.Intn(w.width)
	y = rand.Intn(w.height)
	for w.GetTile(x, y).isWalkable() == false {
		x = rand.Intn(w.width)
		y = rand.Intn(w.height)
	}
	return
}

/* Find an entity by coordinates
 * FIXME: Can be improved by previously indexing entities
 */
func (w *world) entityAt(x, y int) autonomous {
	for e := w.entities.Front(); e != nil; e = e.Next() {
		entity := e.Value.(autonomous)
		entityX, entityY := entity.Position()
		if entityX == x && entityY == y {
			return entity
		}
	}
	return nil
}

func (w *world) removeEntity(entity autonomous) {
	for e := w.entities.Front(); e != nil; e = e.Next() {
		if entity == e.Value.(autonomous) {
			w.entities.Remove(e)
		}
	}
}
