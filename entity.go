package main

type entity struct {
	x, y  int
	glyph rune
	color int
	world *world
}

type player entity

func newPlayer(world *world) *player {
	x, y := world.atWalkableTile()
	return &player{x, y, '@', 0, world}
}

func (p *player) move(offsetX, offsetY int) {
	newX := p.x + offsetX
	newY := p.y + offsetY
	tile := p.world.GetTile(newX, newY)
	entity := p.world.entityAt(newX, newY)
	if entity != nil {
		p.world.removeEntity(entity)
	} else if tile.isWalkable() {
		p.x = newX
		p.y = newY
	} else if tile.isDiggable() {
		p.world.dig(newX, newY)
	}
}

type autonomous interface {
	Position() (int, int)
	Avatar() (rune, int)
	update()
}
