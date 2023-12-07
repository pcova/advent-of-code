package engine

import (
	"fmt"
	"strconv"
	"unicode"
)

type Schematic struct {
	Width   int
	Height  int
	tiles   [][]rune
	symbols []Point
}

type Point struct {
	X int
	Y int
}

func (m *Schematic) AppendLine(line string) {
	m.Width = len(line)

	for i, c := range line {
		if c != '.' && !unicode.IsDigit(c) {
			m.symbols = append(m.symbols, Point{X: i, Y: m.Height})
		}
	}

	m.tiles = append(m.tiles, []rune(line))
	m.Height++
}

func (m *Schematic) GetTile(p Point) rune {
	return m.tiles[p.Y][p.X]
}

func (m *Schematic) GetSymbols() []Point {
	return m.symbols
}

/**
 * Get the number at the given point.
 * The number is defined as the longest string of digits adjacent to the given point.
 * The number is returned as an int, along with the start and end points of the number.
 */
func (m *Schematic) GetNumber(p Point) (int, *Point, *Point, error) {
	var start, end int = p.X, p.X
	for i := p.X - 1; i >= 0; i-- {
		if !unicode.IsDigit(m.tiles[p.Y][i]) {
			break
		}
		start = i
	}

	for i := p.X + 1; i < m.Width; i++ {
		if !unicode.IsDigit(m.tiles[p.Y][i]) {
			break
		}
		end = i
	}

	number, err := strconv.Atoi(string(m.tiles[p.Y][start : end+1]))
	if err != nil {
		return 0, nil, nil, fmt.Errorf("could not convert %s to int", string(m.tiles[p.Y][start:end+1]))
	}

	return number, &Point{X: start, Y: p.Y}, &Point{X: end, Y: p.Y}, nil
}

func (m *Schematic) GetNumbersAdjacentTo(p Point) ([]int, error) {
	var numbers []int

	for j := p.Y - 1; j <= p.Y+1; j++ {
		for i := p.X - 1; i <= p.X+1; i++ {
			if i == p.X && j == p.Y {
				continue
			}

			if i < 0 || i >= m.Width {
				continue
			}

			if j < 0 || j >= m.Height {
				continue
			}

			if unicode.IsDigit(m.tiles[j][i]) {
				number, _, end, err := m.GetNumber(Point{X: i, Y: j})
				if err != nil {
					return nil, fmt.Errorf("could not get number adjacent to %v: %w", p, err)
				}
				numbers = append(numbers, number)
				i = end.X
			}
		}
	}

	return numbers, nil
}
