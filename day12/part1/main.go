package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	hotsprings "github.com/pcova/advent-of-code/day12/hot_springs"
)

/**
 * ParseRow parses a line of input and returns a Row.
 *
 * The line should be in the format:
 *   <springs> <damaged springs groups>
 *
 * where <springs> is a string of '.' (operational), '#' (damaged)
 * or '?' (unknown) characters and <number of contiguous damaged springs> is a
 * comma-separated list of integers.
 *
 * For example:
 *   .#?.#..#. 2,1,1
 *
 * The function returns an error if the line is not in the expected format.
 */
func ParseRow(line string) (*hotsprings.Row, error) {
	// Split the line in two parts: the springs and the contiguous damaged springs groups.
	parts := strings.Split(line, " ")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid line: %s", line)
	}

	// Parse the springs.
	springs := []hotsprings.Spring{}
	for _, spring := range parts[0] {
		springs = append(springs, hotsprings.Spring(spring))
	}

	// Parse the contiguous damaged springs groups.
	contiguousDamagedSprings := []int{}
	for _, groupSizeStr := range strings.Split(parts[1], ",") {
		groupSize, err := strconv.Atoi(groupSizeStr)
		if err != nil {
			return nil, fmt.Errorf("invalid group size: %s", groupSizeStr)
		}
		contiguousDamagedSprings = append(contiguousDamagedSprings, groupSize)
	}

	return hotsprings.NewRow(springs, contiguousDamagedSprings), nil
}

func ParseRows(scanner *bufio.Scanner) ([]*hotsprings.Row, error) {
	rows := []*hotsprings.Row{}

	for scanner.Scan() {
		line := scanner.Text()

		row, err := ParseRow(line)
		if err != nil {
			return nil, fmt.Errorf("error parsing row '%s': %v", line, err)
		}

		rows = append(rows, row)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return rows, nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	rows, err := ParseRows(scanner)
	if err != nil {
		log.Fatal(err)
	}

	sum := 0
	for _, row := range rows {
		validRows := row.ValidAlternativeRows(true)

		sum += validRows
	}

	fmt.Printf("Sum: %d\n", sum)
}
