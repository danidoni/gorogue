package main

type entity struct {
	x, y  int
	glyph rune
	color int
	world *world
	hp    int
	maxHp int
}

func (e entity) Hp() int {
	return e.hp
}

func (e *entity) SetHp(amount int) {
	e.hp = amount
}

func (e entity) MaxHp() int {
	return e.maxHp
}

func (e *entity) SetMaxHp(value int) {
	e.maxHp = value
}

type player struct {
	entity
}

func newPlayer(world *world) *player {
	x, y := world.atWalkableTile()
	return &player{entity{
		x: x,
		y: y,
		glyph: '@',
		color: 0,
		world: world,
		hp: 100,
		maxHp: 100,
	}}
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
	Hp() int
	SetHp(value int)
}
