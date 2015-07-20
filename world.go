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
}

type world struct {
	width, height int
	cells         [][]*tile
}

func NewTile(kind tileType) *tile {
	switch {
	case kind == floor:
		return &tile{'.', floor}
	case kind == wall:
		return &tile{'#', wall}
	}
	return &tile{'X', boundary}
}

func NewWorld(width, height int) *world {
	return RandomWorld(width, height)
}

func (w world) Viewport(x, y, width, height int) [][]*tile {
	// TODO: How to render cells smaller than the viewport
	slice := w.cells[y : y+height]
	viewport := make([][]*tile, len(slice))
	for r := range slice {
		viewport[r] = slice[r][x : x+width]
	}
	return viewport
}
