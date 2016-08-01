package main

type viewport struct {
	origin        *Point
	width, height int
	world         *world
}

func NewViewport(center *Point, width, height int, world *world) *viewport {
	v := &viewport{width: width, height: height, world: world}
	v.center(center.x, center.y)
	return v
}

func (v *viewport) viewportToWorld(viewportX, viewportY int) (worldX, worldY int) {
	worldX = v.origin.x + viewportX
	worldY = v.origin.y + viewportY
	return
}

func (v *viewport) worldToViewport(coords *Point) *Point {
	return coords.Sub(v.origin)
}

func (v *viewport) center(centeredX, centeredY int) {
	v.origin = &Point{
		x: centeredX - v.width/2,
		y: centeredY - v.height/2,
	}
}

// TODO; Make private
func (v *viewport) GetTile(viewportX, viewportY int) *tile {
	worldX, worldY := v.viewportToWorld(viewportX, viewportY)

	switch {
	case viewportX < 0:
		return NewTile(boundary)
	case viewportX > v.width:
		return NewTile(boundary)
	case viewportY < 0:
		return NewTile(boundary)
	case viewportY > v.height:
		return NewTile(boundary)
	case worldX > v.world.dimensions.x:
		return NewTile(boundary)
	case worldY > v.world.dimensions.y:
		return NewTile(boundary)
	}

	return v.world.GetTile(&Point{x: worldX, y: worldY})
}

// TODO: Make private
func (v *viewport) Move(direction direction, step int, w *world) {
	switch {
	case direction == left:
		v.origin = v.origin.Sub(&Point{step, 0})
	case direction == right:
		v.origin = v.origin.Add(&Point{step, 0})
	case direction == up:
		v.origin = v.origin.Sub(&Point{0, step})
	case direction == down:
		v.origin = v.origin.Add(&Point{0, step})
	}
}

func (v *viewport) iterate(callback func(x, y int, tile *tile)) {
	for y := 0; y < v.height; y++ {
		for x := 0; x < v.width; x++ {
			tile := v.GetTile(x, y)
			callback(x, y, tile)
		}
	}
}

func (v *viewport) entities(callback func(entity autonomous)) {
	v.world.entitiesInside(v.origin, v.width, v.height, callback)
}
