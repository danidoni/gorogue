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

type autonomous interface {
	Position() (int, int)
	Avatar() (rune, int)
	update()
	Hp() int
	SetHp(value int)
}
