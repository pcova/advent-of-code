package contraption

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Beam struct {
	direction Direction
	x, y      int
}

func NewBeam(direction Direction, x, y int) *Beam {
	return &Beam{
		direction: direction,
		x:         x,
		y:         y,
	}
}

func (b Beam) X() int {
	return b.x
}

func (b Beam) Y() int {
	return b.y
}

/**
 * Split the beam into two beams.
 * If the beam is traveling up or down, and the tile is a vertical splitter, split the beam into a left and right beam.
 * If the beam is traveling left or right, and the tile is a horizontal splitter, split the beam into a up and down beam.
 * Otherwise, just move the current beam and return nil.
 */
func (b *Beam) Split(tileType rune) *Beam {
	var newBeam *Beam
	switch b.direction {
	case Up, Down:
		switch tileType {
		case '-':
			b.direction = Left
			newBeam = &Beam{Right, b.x + 1, b.y}
		}
	case Left, Right:
		switch tileType {
		case '|':
			b.direction = Up
			newBeam = &Beam{Down, b.x, b.y + 1}
		}
	}
	b.Move()
	return newBeam
}

/**
 * Update the direction of the beam if it hits a mirror.
 */
func (b *Beam) UpdateDirection(tileType rune) {
	switch tileType {
	case '/':
		switch b.direction {
		case Up:
			b.direction = Right
		case Down:
			b.direction = Left
		case Left:
			b.direction = Down
		case Right:
			b.direction = Up
		}
	case '\\':
		switch b.direction {
		case Up:
			b.direction = Left
		case Down:
			b.direction = Right
		case Left:
			b.direction = Up
		case Right:
			b.direction = Down
		}
	}

	b.Move()
}

/**
 * Move the beam in the current direction.
 */
func (b *Beam) Move() {
	switch b.direction {
	case Up:
		b.y = b.y - 1
	case Down:
		b.y = b.y + 1
	case Left:
		b.x = b.x - 1
	case Right:
		b.x = b.x + 1
	}
}

/**
 * BeamTravel takes a contraption and a beam and returns a map of energyzed tiles.
 * The beam starts at the initialBeam's position and travels in the initialBeam's direction.
 * The beam travels until it hits an already traveled splitter or goes out of bounds.
 * If the beam hits a splitter, it splits into two beams and continues.
 * If the beam hits a mirror, it updates its direction and continues.
 * If the beam goes out of bounds, it stops.
 * The map of energyzed tiles is a map of tile types to the number of times the beam hit that tile.
 */
func BeamTravel(contraption Contraption, initialBeam *Beam) map[Tile]int {
	beams := []*Beam{initialBeam}
	energyzedTiles := make(map[Tile]int)
	splitters := make(map[Tile]bool)
	for len(beams) > 0 {
		// Pop the first beam off the stack.
		beam := beams[0]
		beams = beams[1:]

		for {
			// If the beam is out of bounds, stop.
			if beam.X() < 0 || beam.Y() < 0 || beam.X() >= contraption.Width() || beam.Y() >= contraption.Height() {
				break
			}

			tile := contraption[beam.Y()][beam.X()]

			// If the beam has already been here, stop.
			if splitters[tile] {
				break
			}

			// increment the energyzed tile count.
			energyzedTiles[tile]++

			switch tile.Type() {
			// If the beam hits a splitter, split the beam and continue.
			case '|', '-':
				newBeam := beam.Split(tile.Type())
				if newBeam != nil {
					splitters[tile] = true
					beams = append(beams, newBeam)
				}
				continue
			// If the beam hits a mirror, update the direction and continue.
			case '/', '\\':
				beam.UpdateDirection(tile.Type())
				continue
			// If the beam hits a empty space, move the beam and continue.
			case '.':
				beam.Move()
			}
		}
	}

	return energyzedTiles
}
