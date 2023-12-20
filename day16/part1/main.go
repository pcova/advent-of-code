package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/pcova/advent-of-code/day16/contraption"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Read the contraption from stdin.
	var c contraption.Contraption
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		var row []contraption.Tile
		for x, tileType := range line {
			row = append(row, *contraption.NewTile(x, y, tileType))
		}
		c = append(c, row)
		y++
	}

	energyzedTiles := contraption.BeamTravel(c, contraption.NewBeam(contraption.Right, 0, 0))

	fmt.Printf("# energyzed tiles: %d\n", len(energyzedTiles))
}
