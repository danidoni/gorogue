package main

import (
	"math/rand"
	"fmt"
)

type player struct {
	entity
}

func newPlayer(world *world) *player {
	point := world.atWalkableTile()
	stats := &Stats{hp: 100, maxHp: 100, attack: 20, defense: 5}
	return &player{
		entity{
			location: point,
			glyph: '@',
			color: 0,
			world: world,
			stats: stats,
			name: "Player",
		}}
}

func (p *player) move(offset *Point) {
	newLocation := p.location.Add(offset)
	tile := p.world.GetTile(newLocation)
	entity := p.world.entityAt(newLocation.x, newLocation.y)
	if entity != nil {
		p.Attack(entity)
	} else if tile.isWalkable() {
		p.location.x = newLocation.x
		p.location.y = newLocation.y
	} else if tile.isDiggable() {
		p.world.dig(newLocation)
	}
}

func (p *player) Attack(foe autonomous) {
	damage := p.Stats().Attack() - foe.Stats().Defense()

	amount := 0
	if damage > 0 {
		amount = damage
	}
	amount = rand.Intn(amount + 1)
	p.world.notifications.notify(fmt.Sprintf("You attack the '%s' for %d damage.", foe.Name(), amount))
	foe.Stats().ApplyDamage(amount)
	if foe.Stats().Hp() <= 0 {
		p.OnKill(foe)
	}
}

func (p *player) OnKill(entity autonomous) {
	p.world.entities.remove(entity)
	p.world.notifications.notify("You have slain the creep!")
}

func (p *player) Name() string {
	return p.name
}
