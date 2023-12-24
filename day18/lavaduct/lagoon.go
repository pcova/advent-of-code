package lavaduct

import "math"

func shoelaceFormula(vertices []Point) int {
	var area int
	for i := 0; i < len(vertices)-1; i++ {
		area += vertices[i].x*vertices[i+1].y - vertices[i+1].x*vertices[i].y
	}
	area += vertices[len(vertices)-1].x*vertices[0].y - vertices[0].x*vertices[len(vertices)-1].y
	return int(math.Abs(float64(area)) / 2)
}

func interiorPoints(area int, boundaryPoints int) int {
	return area - boundaryPoints/2 + 1
}

type Lagoon struct {
	vertices []Point
}

func (l *Lagoon) Volume() int {
	area := shoelaceFormula(l.vertices)

	boundaryPoints := len(l.vertices)
	interiorPoints := interiorPoints(area, boundaryPoints)

	return interiorPoints + boundaryPoints
}
