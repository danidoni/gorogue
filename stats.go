package main

type Stats struct {
	hp    int
	maxHp int
}

func (s Stats) Hp() int {
	return s.hp
}

func (s *Stats) SetHp(amount int) {
	s.hp = amount
}

func (s Stats) MaxHp() int {
	return s.maxHp
}

func (s *Stats) SetMaxHp(value int) {
	s.maxHp = value
}
