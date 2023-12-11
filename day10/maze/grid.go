package maze

import (
	"math"
)

type Grid [][]*Tile

func (g Grid) Width() int {
	return len(g[0])
}

func (g Grid) Height() int {
	return len(g)
}

func (g Grid) GetTile(x, y int) *Tile {
	return g[y][x]
}

/**
 * GetAdjacentTiles returns a slice of tiles adjacent to the given tile.
 *
 * Example:
 *   grid =
 *   [ 1 2 3 4 5 ]
 *   [ 6 7 8 9 0 ]
 *   [ a b c d e ]
 *   [ f g h i j ]
 *   [ k l m n o ]
 *	 get adjacent tiles for (2, 2) (c):
 *   [ 7 8 9 b c d g h i ]
 */
func (g Grid) GetAdjacentTiles(x, y int) []*Tile {
	var cells []*Tile

	startY := int(math.Max(0, float64(y-1)))
	endY := int(math.Min(float64(g.Height()-1), float64(y+1)))
	startX := int(math.Max(0, float64(x-1)))
	endX := int(math.Min(float64(g.Width()-1), float64(x+1)))

	rows := g[startY : endY+1]
	for _, row := range rows {
		cells = append(cells, row[startX:endX+1]...)
	}

	return cells
}

func (g *Grid) MakeConnections() {
	for y, row := range *g {
		for x, tile := range row {
			if tile.Type == EmptyTile {
				continue
			}

			// Reset NextTile
			tile.NextTiles = make([]*Tile, 0)

			for _, otherTile := range g.GetAdjacentTiles(x, y) {
				if tile.CanConnect(otherTile) {
					tile.NextTiles = append(tile.NextTiles, otherTile)
				}
			}
		}
	}
}
