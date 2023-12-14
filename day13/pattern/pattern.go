package pattern

import (
	"strings"
)

/**
 * Returns a slice of indexes where the strings differ.
 * The strings must be of equal length.
 */
func diff(a, b string) []int {
	if len(a) != len(b) {
		panic("strings must be of equal length")
	}

	diff := make([]int, 0)
	for i, c := range a {
		if c != rune(b[i]) {
			diff = append(diff, i)
		}
	}

	return diff
}

type Pattern struct {
	rows    []string
	columns []string
	smudge  bool // if true, the pattern should be fixed by changing one symbol
}

/**
 * Returns a string representation of the pattern.
 */
func (p Pattern) String() string {
	pStr := ""

	hRef := p.HReflectionLine()
	vRef := p.VReflectionLine()

	for i, row := range p.rows {
		if hRef > 0 && i == hRef {
			pStr += strings.Repeat("-", len(row)) + "\n"
		}
		for j, c := range row {
			if vRef > 0 && j == vRef {
				pStr += "|"
			}
			pStr += string(c)
		}
		if i < len(p.rows)-1 {
			pStr += "\n"
		}
	}

	return pStr
}

/**
 * Sets the smudge flag to true.
 */
func (p *Pattern) FixSmudge() {
	p.smudge = true
}

/**
 * Adds a row to the pattern and updates the columns.
 */
func (p *Pattern) AddRow(row string) {
	// add row to rows
	p.rows = append(p.rows, row)

	// add columns
	for i, c := range row {
		if len(p.columns) <= i {
			p.columns = append(p.columns, string(c))
		} else {
			p.columns[i] += string(c)
		}
	}
}

/**
 * Returns the horizontal/vertical reflection line.
 *
 * list: the list of rows/columns
 * smudge: if true, the pattern *should* be fixed by changing one symbol
 * start: the index to start the search from
 */
func reflectionLine(list []string, smudge bool, start int) int {
	// if smudged is true, the pattern was fixed by changing one symbol
	var smudged bool = false

	// finds the first line that is equal to the line after it (reflection line)
	// or the first line that is equal to the line after it if one symbol is changed
	// the search starts at the given 'start' index
	var lineIdx int = -1
	for i := start; i < len(list)-1; i++ {
		diff := diff(list[i], list[i+1])
		if len(diff) == 0 || (!smudged && smudge && len(diff) == 1) {
			lineIdx = i + 1
			break
		}
	}

	// if no line was found, then no reflection line exists
	if lineIdx == -1 {
		return 0
	}

	// check if the current reflection line is valid
	// the reflection line is valid if all the pair of lines before and after it
	// are equal or can be fixed by changing one symbol when smudge flag is true
	//
	// Example:
	// #..##.#..# <- pair 2
	// #..#...#.# <- pair 1
	// ---------- <- reflection line
	// #..#...#.# <- pair 1
	// #..##.#..# <- pair 2
	for i := 0; i < len(list); i++ {
		// the indexes of the pair of lines to compare
		bIdx := lineIdx - i - 1
		aIdx := lineIdx + i

		// if the indexes are out of bounds, then the validation is complete
		if bIdx < 0 || aIdx >= len(list) {
			break
		}

		// check if the lines are equal or if they can be fixed by changing one symbol when smudge flag is true
		diff := diff(list[bIdx], list[aIdx])
		if len(diff) == 0 || (!smudged && smudge && len(diff) == 1) {
			// if there is only one symbol that differs, then the pattern was smudged
			if len(diff) == 1 {
				smudged = true
			}
			// continue to the next pair of lines
			continue
		}

		// if the lines are not equal or cannot be fixed by changing one symbol when smudge flag is true,
		// then the current reflection line is not valid
		// try to find a new reflection line starting from the current line
		return reflectionLine(list, smudge, lineIdx)
	}

	// if smudge flag is true and the pattern was not smudged,
	// then try to find a new reflection line starting from the current line
	if smudge && !smudged {
		return reflectionLine(list, smudge, lineIdx)
	}

	// if the current reflection line is valid, then return it
	return lineIdx
}

// Returns the horizontal reflection line.
func (p *Pattern) HReflectionLine() int {
	return reflectionLine(p.rows, p.smudge, 0)
}

// Returns the vertical reflection line.
func (p *Pattern) VReflectionLine() int {
	return reflectionLine(p.columns, p.smudge, 0)
}

// Calculates the score of the pattern.
func (p *Pattern) Score() int {
	hLine := p.HReflectionLine()

	if hLine == 0 {
		return p.VReflectionLine()
	}

	return hLine * 100
}
