package cosmos

import (
	"fmt"
	"math"
)

type Galaxy struct {
	id   int
	x, y int
}

/**
 * PathLength calculates the number of steps between two galaxies.
 * This is done by calculating the Manhattan distance between the two galaxies.
 */
func (g Galaxy) PathLength(other Galaxy) int {
	return int(math.Abs(float64(g.x-other.x))) + int(math.Abs(float64(g.y-other.y)))
}

func (g Galaxy) String() string {
	return fmt.Sprintf("%d (%d, %d)", g.id, g.x, g.y)
}

type GalaxyList []Galaxy

func (galaxies GalaxyList) Len() int {
	return len(galaxies)
}

func (galaxies *GalaxyList) CalculateDistance() int {
	distance := 0
	for i := 0; i < galaxies.Len(); i++ {
		for j := i + 1; j < galaxies.Len(); j++ {
			galaxy := (*galaxies)[i]
			other := (*galaxies)[j]

			d := galaxy.PathLength(other)

			// debug
			// fmt.Printf("%s -> %s = %d\n", galaxy, other, d)

			distance += d
		}
	}
	return distance
}
