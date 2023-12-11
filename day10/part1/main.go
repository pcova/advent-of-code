package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/pcova/advent-of-code/day10/maze"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	m := maze.ParseMaze(scanner)

	pipes := m.GetLoopPipes()

	fmt.Printf("Farthest tile distance: %d\n", len(pipes)/2)
}
