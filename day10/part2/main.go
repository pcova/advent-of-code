package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	"github.com/pcova/advent-of-code/day10/maze"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	m := maze.ParseMaze(scanner)

	pipes := m.GetLoopPipes()

	// Uses the shoelace formula to calculate the area of a polygon
	a := 0
	b := 0
	for i := 0; i < len(pipes); i++ {
		nextIndex := (i + 1) % len(pipes)
		a += pipes[i].X * pipes[nextIndex].Y
		b += pipes[i].Y * pipes[nextIndex].X
	}
	area := math.Abs(float64(a-b)) / 2
	fmt.Printf("Area: %.1f\n", area)

	// Uses Pick's theorem to calculate the number of interior points
	interiorPoints := area - float64(len(pipes)/2) + 1
	fmt.Printf("Interior points: %.f\n", interiorPoints)
}
