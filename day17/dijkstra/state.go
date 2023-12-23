package dijkstra

import "fmt"

type State struct {
	point        Point
	dx, dy       int
	distStraight int
	path         []Point
}

func (s State) String() string {
	return fmt.Sprintf("%v-%d-%d-%d", s.point, s.dx, s.dy, s.distStraight)
}

type StatesByDistance map[int][]State

func (d StatesByDistance) Add(key int, state State) {
	slice, ok := d[key]
	if !ok {
		d[key] = []State{state}
		return
	}
	d[key] = append(slice, state)
}

func (d StatesByDistance) PopSmaller() (int, []State) {
	var minKey int = -1
	for key := range d {
		if minKey == -1 || key < minKey {
			minKey = key
		}
	}
	slice := d[minKey]
	delete(d, minKey)
	return minKey, slice
}
