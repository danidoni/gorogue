package main

type viewport struct {
	x, y          int
	width, height int
	world         *world
}

func centeredViewport(centeredX, centeredY, width, height int, world *world) *viewport {
	x := centeredX - width/2
	y := centeredY - height/2
	return &viewport{x, y, width, height, world}
}

// TODO; Make private
func (v *viewport) GetTile(viewportX, viewportY int) *tile {
	worldY := v.y + viewportY
	worldX := v.x + viewportX

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
