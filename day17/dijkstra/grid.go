package dijkstra

type Grid [][]int

func (g Grid) Height() int {
	return len(g)
}

func (g Grid) Width() int {
	return len(g[0])
}

func (g Grid) Get(point Point) int {
	return g[point.y][point.x]
}

func (g Grid) IsInBounds(point Point) bool {
	return point.x >= 0 && point.y >= 0 && point.x < g.Width() && point.y < g.Height()
}
