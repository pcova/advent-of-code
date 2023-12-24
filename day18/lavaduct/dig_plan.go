package lavaduct

type Direction rune

const (
	Up    Direction = 'U'
	Down  Direction = 'D'
	Left  Direction = 'L'
	Right Direction = 'R'
)

func IntToDirection(i int) Direction {
	switch i {
	case 0:
		return Right
	case 1:
		return Down
	case 2:
		return Left
	case 3:
		return Up
	default:
		panic("invalid direction")
	}
}

type Point struct {
	x, y int
}

type Instruction struct {
	direction Direction
	meters    int
}

type DigPlan []Instruction

func (d *DigPlan) AddInstruction(direction rune, meters int) {
	instruction := Instruction{
		direction: Direction(direction),
		meters:    meters,
	}
	*d = append(*d, instruction)
}

func (d DigPlan) Lagoon() Lagoon {
	vertices := make([]Point, 0)

	// Start at the origin.
	currentPoint := Point{0, 0}
	vertices = append(vertices, currentPoint)

	// Follow the dig plan.
	for _, instruction := range d {
		for i := 0; i < instruction.meters; i++ {
			switch instruction.direction {
			case Up:
				currentPoint.y++
			case Down:
				currentPoint.y--
			case Left:
				currentPoint.x--
			case Right:
				currentPoint.x++
			default:
				panic("invalid direction")
			}
			vertices = append(vertices, currentPoint)
		}
	}

	// Remove the last vertex, which is the same as the first.
	vertices = vertices[:len(vertices)-1]

	return Lagoon{
		vertices: vertices,
	}
}
