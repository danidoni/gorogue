package main

type player struct {
	entity
}

func newPlayer(world *world) *player {
	point := world.atWalkableTile()
	stats := &Stats{hp: 100, maxHp: 100}
	return &player{
		entity{
			location: point,
			glyph: '@',
			color: 0,
			world: world,
			stats: stats,
		}}
}

func (p *player) move(offset *Point) {
	newLocation := p.location.Add(offset)
	tile := p.world.GetTile(newLocation)
	entity := p.world.entityAt(newLocation.x, newLocation.y)
	if entity != nil {
		p.OnKill(entity)
	} else if tile.isWalkable() {
		p.location.x = newLocation.x
		p.location.y = newLocation.y
	} else if tile.isDiggable() {
		p.world.dig(newLocation)
	}
}

func (p *player) OnKill(entity autonomous) {
	p.world.entities.remove(entity)
	p.world.notifications.notify("You have slain the creep!")
}
