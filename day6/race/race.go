package race

import (
	"fmt"
	"math"
)

// findRoots returns the roots of the quadratic equation ax^2 + bx + c = 0
func findRoots(a, b, c int) (float64, float64) {
	discriminant := float64(b*b - 4*a*c)
	if discriminant < 0 {
		panic("no real roots")
	}

	sqrt := func(x float64) float64 {
		return math.Sqrt(x)
	}

	return (-float64(b) + sqrt(discriminant)) / (2 * float64(a)),
		(-float64(b) - sqrt(discriminant)) / (2 * float64(a))
}

type Race struct {
	timeAllowed    int
	recordDistance int
}

func NewRace(timeAllowed, recordDistance int) *Race {
	return &Race{timeAllowed, recordDistance}
}

func (r *Race) String() string {
	return fmt.Sprintf("{time=%d, distance=%d}", r.timeAllowed, r.recordDistance)
}

// NWaysToBeatRecord returns the number of ways to beat the record
func (r *Race) CountWaysToBeatRecord() int {
	max, min := findRoots(1, -r.timeAllowed, r.recordDistance)

	maxInt := int(math.Floor(max))
	if math.Floor(max) == max { // if max is an integer, we don't want to count it
		maxInt--
	}

	minInt := int(math.Ceil(min))
	if math.Ceil(min) == min { // if min is an integer, we don't want to count it
		minInt++
	}

	// we want to count the number of integers between min and max, inclusive
	return (maxInt - minInt) + 1
}
