package maze

type TileType rune

const (
	EmptyTile         TileType = '.'
	Vertical          TileType = '|'
	Horizontal        TileType = '-'
	TopRightCorner    TileType = 'F'
	TopLeftCorner     TileType = '7'
	BottomRightCorner TileType = 'L'
	BottomLeftCorner  TileType = 'J'
	Start             TileType = 'S'
)

func (t TileType) CanConnectUp(other TileType) bool {
	switch t {
	case Vertical, BottomRightCorner, BottomLeftCorner, Start:
		return other == Vertical || other == TopRightCorner || other == TopLeftCorner || other == Start
	case Horizontal, TopRightCorner, TopLeftCorner:
		return false
	}

	return false
}

func (t TileType) CanConnectDown(other TileType) bool {
	switch t {
	case Vertical, TopRightCorner, TopLeftCorner, Start:
		return other == Vertical || other == BottomRightCorner || other == BottomLeftCorner || other == Start
	case Horizontal, BottomRightCorner, BottomLeftCorner:
		return false
	}

	return false
}

func (t TileType) CanConnectLeft(other TileType) bool {
	switch t {
	case Horizontal, TopLeftCorner, BottomLeftCorner, Start:
		return other == Horizontal || other == TopRightCorner || other == BottomRightCorner || other == Start
	case Vertical, TopRightCorner, BottomRightCorner:
		return false
	}

	return false
}

func (t TileType) CanConnectRight(other TileType) bool {
	switch t {
	case Horizontal, TopRightCorner, BottomRightCorner, Start:
		return other == Horizontal || other == TopLeftCorner || other == BottomLeftCorner || other == Start
	case Vertical, TopLeftCorner, BottomLeftCorner:
		return false
	}

	return false
}
