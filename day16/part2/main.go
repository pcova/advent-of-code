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

	// Search for the maximum number of energyzed tiles on the left and right sides of the contraption.
	maxEnergyzedTiles := 0
	for y := 0; y < c.Height(); y++ {
		energyzedTiles := contraption.BeamTravel(c, contraption.NewBeam(contraption.Right, 0, y))
		if len(energyzedTiles) > maxEnergyzedTiles {
			maxEnergyzedTiles = len(energyzedTiles)
		}

		energyzedTiles = contraption.BeamTravel(c, contraption.NewBeam(contraption.Left, c.Width()-1, y))
		if len(energyzedTiles) > maxEnergyzedTiles {
			maxEnergyzedTiles = len(energyzedTiles)
		}
	}

	// Search for the maximum number of energyzed tiles on the top and bottom sides of the contraption.
	for x := 0; x < c.Width(); x++ {
		energyzedTiles := contraption.BeamTravel(c, contraption.NewBeam(contraption.Down, x, 0))
		if len(energyzedTiles) > maxEnergyzedTiles {
			maxEnergyzedTiles = len(energyzedTiles)
		}

		energyzedTiles = contraption.BeamTravel(c, contraption.NewBeam(contraption.Up, x, c.Height()-1))
		if len(energyzedTiles) > maxEnergyzedTiles {
			maxEnergyzedTiles = len(energyzedTiles)
		}
	}

	fmt.Printf("# max energyzed tiles: %d\n", maxEnergyzedTiles)
}
