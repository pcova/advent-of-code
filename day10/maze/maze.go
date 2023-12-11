package maze

import "bufio"

type Maze struct {
	grid         *Grid
	startingPipe *Tile
}

func NewMaze(grid *Grid, startingPipe *Tile) *Maze {
	return &Maze{
		grid:         grid,
		startingPipe: startingPipe,
	}
}

func ParseMaze(scanner *bufio.Scanner) *Maze {
	var grid Grid
	var startingPipe *Tile

	for scanner.Scan() {
		line := scanner.Text()

		var row []*Tile
		for _, char := range line {
			pipe := &Tile{
				Type: TileType(char),
				X:    len(row),
				Y:    len(grid),
			}

			if pipe.Type == Start {
				startingPipe = pipe
			}

			row = append(row, pipe)
		}

		grid = append(grid, row)
	}

	grid.MakeConnections()

	return NewMaze(&grid, startingPipe)
}

func (m *Maze) GetLoopPipes() []*Tile {
	var pipes []*Tile

	traveledTiles := make(map[string]bool)

	// Start with one of the next tiles of the starting pipe
	nextTiles := append([]*Tile{}, m.startingPipe.NextTiles[0])
	// Mark the starting tile as traveled
	traveledTiles[m.startingPipe.Id()] = true

	pipes = append(pipes, m.startingPipe)

	for len(nextTiles) > 0 {
		var nextNextTiles []*Tile

		for _, tile := range nextTiles {
			// Mark the tile as traveled
			traveledTiles[tile.Id()] = true

			pipes = append(pipes, tile)

			for _, nextTile := range tile.NextTiles {
				// If we've already traveled this tile, skip it
				if traveledTiles[nextTile.Id()] {
					continue
				}
				nextNextTiles = append(nextNextTiles, nextTile)
			}
		}

		// Set the next tiles for the next iteration
		nextTiles = nextNextTiles
	}

	return pipes
}
