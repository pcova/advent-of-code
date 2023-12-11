package cosmos

import "strings"

const (
	NO_EXPENSION     = 0
	EXPENSION_RATE_1 = 2
	EXPENSION_RATE_2 = 1000000
)

type Image [][]rune

func (i Image) Width() int {
	return len(i[0])
}

func (i Image) Height() int {
	return len(i)
}

// empty row is a row of dots
func (i *Image) AddEmptyRow(y int) {
	row := make([]rune, i.Width())
	for x := range row {
		row[x] = '.'
	}
	*i = append((*i)[:y], append([][]rune{row}, (*i)[y:]...)...)
}

// empty column is a column of dots
func (i *Image) AddEmptyColumn(x int) {
	for y := range *i {
		(*i)[y] = append((*i)[y][:x], append([]rune{'.'}, (*i)[y][x:]...)...)
	}
}

func (i *Image) GetRow(y int) []rune {
	return (*i)[y]
}

func (i *Image) GetColumn(x int) []rune {
	column := make([]rune, i.Height())
	for y := range *i {
		column[y] = (*i)[y][x]
	}
	return column
}

func (i *Image) IsRowEmpty(y int) bool {
	return !strings.Contains(string(i.GetRow(y)), "#")
}

func (i *Image) IsColumnEmpty(x int) bool {
	return !strings.Contains(string(i.GetColumn(x)), "#")
}

/**
 * Expand the image by adding empty rows and columns by each row or column that is empty.
 * This is done by checking if a row or column contains a #. If it does, then it's not empty.
 * If it doesn't, then it's empty and we add a row or column of dots.
 *
 * Note:
 *   This method is much less efficient to use than using directly GetGalaxies with a non-zero expension rate.
 *   This method is only used for debugging/visualization purposes.
 */
func (i *Image) Expand(expensionRate int) {
	// iterate over the image rows
	for y := 0; y < i.Height(); y++ {
		row := string((*i)[y])
		// if row contains a #, then it's not empty
		if strings.Contains(row, "#") {
			continue
		}

		// add (expensionRate - 1) rows of dots to the image
		for j := 0; j < expensionRate-1; j++ {
			i.AddEmptyRow(y)
		}
		y += expensionRate - 1
	}

	// iterate over the image columns
	for x := 0; x < i.Width(); x++ {
		column := make([]rune, i.Height())
		for y := range *i {
			column[y] = (*i)[y][x]
		}
		// if column contains a #, then it's not empty
		if strings.Contains(string(column), "#") {
			continue
		}

		// add (expensionRate - 1) columns of dots to the image
		for j := 0; j < expensionRate-1; j++ {
			i.AddEmptyColumn(x)
		}
		x += expensionRate - 1
	}
}

func (i Image) String() string {
	s := ""
	for y := range i {
		s += string(i[y]) + "\n"
	}
	return s
}

/**
 * GetGalaxies returns a list of all galaxies in the image.
 * The expension rate is used to virtually expand the image by adding n empty rows and columns by each row or column that is empty.
 * This is done by checking if a row or column contains a #. If it does, then it's not empty.
 * If it doesn't, then it's empty and we add an ammount of rows or columns of dots equal to the expension rate.
 *
 * Note: this method is much more efficient to use than using Expand() and then GetGalaxies with a zero expension rate.
 */
func (i Image) GetGalaxies(expensionRate int) *GalaxyList {
	galaxies := GalaxyList{}

	id := 0
	yOffset := 0
	for y := range i {
		// if row is empty, then increase the offset by the expension rate - 1 (if expension rate is > 0)
		if expensionRate > 0 && i.IsRowEmpty(y) {
			yOffset += (expensionRate - 1)
			continue
		}

		xOffset := 0
		for x := range i[y] {
			// if column is empty, then increase the offset by the expension rate - 1 (if expension rate is > 0)
			if expensionRate > 0 && i.IsColumnEmpty(x) {
				xOffset += (expensionRate - 1)
				continue
			}

			// if the current position is a #, then add a new galaxy
			if i[y][x] == '#' {
				id++

				// the x and y coordinates are increased by
				// the expension rate times the number of rows or columns
				// that are empty before the current row or column (respectively)
				galaxies = append(galaxies, Galaxy{id, x + xOffset, y + yOffset})
			}
		}
	}
	return &galaxies
}
