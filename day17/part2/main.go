package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/pcova/advent-of-code/day17/dijkstra"
)

func ParseGrid() *dijkstra.Grid {
	scanner := bufio.NewScanner(os.Stdin)
	var grid dijkstra.Grid
	for scanner.Scan() {
		line := scanner.Text()
		var row []int
		for _, char := range line {
			num := int(char - '0')
			if num < 0 || num > 9 {
				panic("invalid number")
			}
			row = append(row, num)
		}
		grid = append(grid, row)
	}
	return &grid
}

func main() {
	grid := ParseGrid()

	// Create a restriction that super crucible has to travel at least 4 steps straight before turning.
	// Create a restriction that super crucible can't travel more than 10 steps straight before being forced to turn.
	restrictions := dijkstra.Restrictions{
		MinStraight: 4,
		MaxStraight: 10,
	}
	shortestPath := dijkstra.ShortestPathDistance(grid, dijkstra.NewPoint(0, 0), dijkstra.NewPoint(grid.Width()-1, grid.Height()-1), restrictions)

	fmt.Printf("least heat loss: %d\n", shortestPath)
}
