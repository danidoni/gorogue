package main

type viewport struct {
	x, y          int
	width, height int
	world         *world
}

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

func (v *viewport) Move(direction direction, step int, w *world) {
	switch {
	case direction == left:
		updatedX := v.x - step
		// if updatedX >= 0 {
			v.x = updatedX
		// }
	case direction == right:
		updatedX := v.x + step
		// if updatedX+v.width <= w.width {
			v.x = updatedX
		// }
	case direction == up:
		updatedY := v.y - step
		// if updatedY >= 0 {
			v.y = updatedY
		// }
	case direction == down:
		updatedY := v.y + step
		// if updatedY+v.height <= w.height {
			v.y = updatedY
		// }
	}
}
