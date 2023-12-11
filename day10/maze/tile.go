package maze

import "fmt"

type Tile struct {
	NextTiles []*Tile
	Type      TileType
	X, Y      int
}

func (t *Tile) String() string {
	return fmt.Sprintf("(%d, %d)", t.X, t.Y)
}

func (t *Tile) Id() string {
	return fmt.Sprintf("%d-%d", t.X, t.Y)
}

func (t *Tile) CanConnect(other *Tile) bool {
	if t.X == other.X && t.Y == other.Y {
		return false
	}

	if t.X == other.X {
		if t.Y > other.Y {
			return t.Type.CanConnectUp(other.Type)
		} else {
			return t.Type.CanConnectDown(other.Type)
		}
	} else if t.Y == other.Y {
		if t.X > other.X {
			return t.Type.CanConnectLeft(other.Type)
		} else {
			return t.Type.CanConnectRight(other.Type)
		}
	}

	return false
}
