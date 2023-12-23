package dijkstra

import "fmt"

type Point struct {
	x, y int
}

func NewPoint(x, y int) Point {
	return Point{
		x: x,
		y: y,
	}
}

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

func (p Point) Move(dx, dy int) Point {
	return Point{
		x: p.x + dx,
		y: p.y + dy,
	}
}

func (p Point) Equal(other Point) bool {
	return p.x == other.x && p.y == other.y
}
