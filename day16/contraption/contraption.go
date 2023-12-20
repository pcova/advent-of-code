package contraption

type Tile struct {
	x, y     int
	tileType rune
}

func NewTile(x, y int, tileType rune) *Tile {
	return &Tile{
		x:        x,
		y:        y,
		tileType: tileType,
	}
}

func (t Tile) Type() rune {
	return t.tileType
}

type Contraption [][]Tile

func (c Contraption) Height() int {
	return len(c)
}

func (c Contraption) Width() int {
	return len(c[0])
}
