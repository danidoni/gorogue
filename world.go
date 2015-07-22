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

type world struct {
	width, height int
	cells         [][]*tile
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

func (w world) GetTile(x, y int) *tile {
	switch {
	case x < 0:
		return NewTile(boundary)
	case x >= w.width:
		return NewTile(boundary)
	case y < 0:
		return NewTile(boundary)
	case y >= w.height:
		return NewTile(boundary)
	}
	return w.cells[y][x]
}
