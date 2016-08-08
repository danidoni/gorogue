package main

type Stats struct {
	hp      int
	maxHp   int
	attack  int
	defense int
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

func (s *Stats) Attack() int {
	return s.attack
}

func (s *Stats) Defense() int {
	return s.defense
}

func (s *Stats) ApplyDamage(amount int) {
	s.hp -= amount
}
