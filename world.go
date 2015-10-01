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
	dimensions    *Point
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
	world := &world{dimensions: &Point{x: width, y: height}}
	world.cells = make([][]*tile, world.dimensions.y)
	for row := range world.cells {
		world.cells[row] = make([]*tile, world.dimensions.x)
		for col := range world.cells[row] {
			world.cells[row][col] = NewTile(floor)
		}
	}
	RandomWorld(world)
	return world
}

func (w world) isWithinBoundaries(point *Point) bool {
	switch {
	case point.x < 0:
		return false
	case point.x >= w.dimensions.x:
		return false
	case point.y < 0:
		return false
	case point.y >= w.dimensions.y:
		return false
	}
	return true
}

func (w world) GetTile(point *Point) *tile {
	if w.isWithinBoundaries(point) {
		return w.cells[point.y][point.x]
	}
	return NewTile(boundary)
}

func (w *world) SetTile(point *Point, tile *tile) {
	if w.isWithinBoundaries(point) {
		w.cells[point.y][point.x] = tile
	}
}

func (w *world) dig(point *Point) {
	w.SetTile(point, NewTile(floor))
}

func (w *world) entitiesInside(point *Point, width, height int, callback func(entity autonomous)) {
	for e := w.entities.Front(); e != nil; e = e.Next() {
		entity := e.Value.(autonomous)
		entityLocation := entity.Position()
		if entityLocation.x >= point.x &&
			entityLocation.x <= point.x+width &&
			entityLocation.y >= point.y &&
			entityLocation.y <= point.y+height {
			callback(entity)
		}
	}
}

// Finds a random walkable tile in the world
func (w *world) atWalkableTile() *Point {
	point := &Point{
		x: rand.Intn(w.dimensions.x),
		y: rand.Intn(w.dimensions.y),
	}
	for w.GetTile(point).isWalkable() == false {
		point = &Point{
			x: rand.Intn(w.dimensions.x),
			y: rand.Intn(w.dimensions.y),
		}
	}
	return point
}

/* Find an entity by coordinates
 * FIXME: Can be improved by previously indexing entities
 */
func (w *world) entityAt(x, y int) autonomous {
	for e := w.entities.Front(); e != nil; e = e.Next() {
		entity := e.Value.(autonomous)
		position := entity.Position()
		if position.x == x && position.y == y {
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
