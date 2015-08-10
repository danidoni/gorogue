package main

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
	return RandomWorld(width, height)
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
