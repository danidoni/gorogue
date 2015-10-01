package main

type Point struct {
	x, y int
}

/* Adds one point to this one and returns a new point with the result
 */
func (p *Point) Add(other *Point) *Point {
	return &Point{
		x: p.x + other.x,
		y: p.y + other.y,
	}
}

func (p *Point) Sub(other *Point) *Point {
	return &Point{
		x: p.x - other.x,
		y: p.y - other.y,
	}
}

func (p *Point) Equal(other *Point) bool {
	return p.x == other.x && p.y == other.y
}
