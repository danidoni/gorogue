package main

type viewport struct {
	x, y          int
	width, height int
	world         *world
}

func centeredViewport(centeredX, centeredY, width, height int, world *world) *viewport {
	v := &viewport{width: width, height: height, world: world}
	v.center(centeredX, centeredY)
	return v
}

func (v *viewport) viewportToWorld(viewportX, viewportY int) (worldX, worldY int) {
	worldX = v.x + viewportX
	worldY = v.y + viewportY
	return
}

func (v *viewport) worldToViewport(worldX, worldY int) (viewportX, viewportY int) {
	viewportX = worldX - v.x
	viewportY = worldY - v.y
	return
}

func (v *viewport) center(centeredX, centeredY int) {
	v.x = centeredX - v.width/2
	v.y = centeredY - v.height/2
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
	case worldX > v.world.width:
		return NewTile(boundary)
	case worldY > v.world.height:
		return NewTile(boundary)
	}

	return v.world.GetTile(worldX, worldY)
}

// TODO: Make private
func (v *viewport) Move(direction direction, step int, w *world) {
	switch {
	case direction == left:
		updatedX := v.x - step
		v.x = updatedX
	case direction == right:
		updatedX := v.x + step
		v.x = updatedX
	case direction == up:
		updatedY := v.y - step
		v.y = updatedY
	case direction == down:
		updatedY := v.y + step
		v.y = updatedY
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
	v.world.entitiesInside(v.x, v.y, v.width, v.height, callback)
}
