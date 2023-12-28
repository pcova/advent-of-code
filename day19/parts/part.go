package parts

import (
	"strconv"
	"strings"
)

/**
 * Part represents a part of a workflow
 * Contains four categories: x, m, a, s
 */
type Part struct {
	x, m, a, s int
}

func (p Part) TotalRating() int {
	return p.x + p.m + p.a + p.s
}

func ParsePart(line string) Part {
	var part Part

	line = strings.TrimPrefix(line, "{")
	line = strings.TrimSuffix(line, "}")

	for _, rating := range strings.Split(line, ",") {
		keyValue := strings.Split(rating, "=")

		value, err := strconv.Atoi(keyValue[1])
		if err != nil {
			panic(err)
		}

		switch keyValue[0] {
		case "x":
			part.x = value
		case "m":
			part.m = value
		case "a":
			part.a = value
		case "s":
			part.s = value
		}
	}

	return part
}

type PartRange map[string]Range

func (pr PartRange) copy() PartRange {
	newMap := make(PartRange)

	for k, v := range pr {
		newMap[k] = v
	}

	return newMap
}

/**
 * Splits PartRange into two PartRanges by a given category and value
 *
 * [1, 10] split by i => [0, i[ and [i, 10]
 *
 * Example:
 *   given range => {x: [2, 11], m: [3, 5], ...}
 *   split by "x" and 5
 *   returns => {x: [2, 4], m: [3, 5], ...}, {x: [5, 11], m: [3, 5], ...}
 */
func (pr PartRange) splitRange(key string, value int) (PartRange, PartRange) {
	lower := pr.copy()
	upper := pr.copy()

	lower[key] = Range{pr[key].Min, value - 1}
	upper[key] = Range{value, pr[key].Max}

	return lower, upper
}

/**
 * Get the total number of possible combinations
 */
func (pr PartRange) Possibilities() int {
	combinations := 1

	for _, r := range pr {
		combinations *= r.PossibleCombinations()
	}

	return combinations
}
