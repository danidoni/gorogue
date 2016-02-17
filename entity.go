package main

import (
	"math/rand"
)

type entity struct {
	location *Point
	glyph rune
	color int
	world *world
	hp    int
	maxHp int
	seed  *rand.Rand
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

func (e *entity) Avatar() (rune, int) {
	return e.glyph, e.color
}

func (e *entity) Position() *Point {
	return e.location
}

func (e *entity) update() {
	// Noop
}

type autonomous interface {
	Position() *Point
	Avatar() (rune, int)
	update()
	Hp() int
	SetHp(value int)
}
