package dijkstra

import "fmt"

type Restrictions struct {
	MaxStraight, MinStraight int
}

func inPath(path []Point, point Point) bool {
	for _, p := range path {
		if p.Equal(point) {
			return true
		}
	}
	return false
}

func printGridPath(grid *Grid, path []Point) {
	for y := 0; y < grid.Height(); y++ {
		for x := 0; x < grid.Width(); x++ {
			point := NewPoint(x, y)
			if inPath(path, point) {
				fmt.Print("#")
			} else {
				fmt.Print(grid.Get(point))
			}
		}
		fmt.Println()
	}
}

func ShortestPathDistance(grid *Grid, start, end Point, restrictions Restrictions) int {
	visited := make(map[string]bool)
	distances := make(StatesByDistance)

	// Add the initial states.
	initState1 := State{start, 1, 0, 1, []Point{start}}
	initState2 := State{start, 0, 1, 1, []Point{start}}
	distances.Add(0, initState1)
	distances.Add(0, initState2)

	for {
		minDist, states := distances.PopSmaller()
		for _, state := range states {

			// If we've reached the end, we're done.
			if state.point.Equal(end) {
				// If we haven't traveled the minimum straight distance, skip it.
				if state.distStraight < restrictions.MinStraight {
					continue
				}
				printGridPath(grid, state.path)
				return minDist
			}

			// If we've already visited this state, skip it.
			if visited[state.String()] {
				continue
			}

			// Mark this state as visited.
			visited[state.String()] = true

			// If we've traveled the minimum straight distance, add the turns.
			if state.distStraight >= restrictions.MinStraight {
				// Add the two possible turns.
				turn1 := state.point.Move(state.dy, -state.dx)
				turn2 := state.point.Move(-state.dy, state.dx)

				if grid.IsInBounds(turn1) {
					distances.Add(minDist+grid.Get(turn1), State{turn1, state.dy, -state.dx, 1, append(append([]Point{}, state.path...), turn1)})
				}

				if grid.IsInBounds(turn2) {
					distances.Add(minDist+grid.Get(turn2), State{turn2, -state.dy, state.dx, 1, append(append([]Point{}, state.path...), turn2)})
				}
			}

			// if we've already traveled the maximum straight distance, skip it.
			if state.distStraight == restrictions.MaxStraight {
				continue
			}

			// Add the straight state.
			next := state.point.Move(state.dx, state.dy)

			if grid.IsInBounds(next) {
				distances.Add(minDist+grid.Get(next), State{next, state.dx, state.dy, state.distStraight + 1, append(append([]Point{}, state.path...), next)})
			}
		}
	}
}
