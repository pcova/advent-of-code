package desert

import "bufio"

type Instruction int

const (
	LEFT Instruction = iota
	RIGHT
)

func (i Instruction) String() string {
	switch i {
	case LEFT:
		return "L"
	case RIGHT:
		return "R"
	default:
		return ""
	}
}

/**
 * parseInstruction parses a rune into an Instruction.
 */
func parseInstruction(c rune) Instruction {
	switch c {
	case 'L':
		return LEFT
	case 'R':
		return RIGHT
	default:
		panic("invalid instruction")
	}
}

/**
 * ParseInstructions parses a string of instructions into a slice of Instruction.
 */
func ParseInstructions(scanner *bufio.Scanner) []Instruction {
	instructions := make([]Instruction, 0)

	scanner.Scan()
	line := scanner.Text()

	for _, c := range line {
		instructions = append(instructions, parseInstruction(c))
	}

	return instructions
}
