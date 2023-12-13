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

const (
	UNFOLD_FACTOR = 5
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

	// Unfolds the list of spring conditions. For example, if the line is:
	//   .#?.#..#. 2,1,1
	// then the springs will be:
	//   .#?.#..#.?.#?.#..#.?.#?.#..#.?.#?.#..#.?.#?.#..#.
	springs := []hotsprings.Spring{}
	for i := 0; i < UNFOLD_FACTOR; i++ {
		// Parse the springs.
		for _, spring := range parts[0] {
			springs = append(springs, hotsprings.Spring(spring))
		}
		// Add a separator.
		if i != UNFOLD_FACTOR-1 {
			springs = append(springs, hotsprings.Unknown)
		}
	}

	// Unfolds the list of blocks sizes. For example, if the line is:
	//   .#?.#..#. 2,1,1
	// then the blocks will be:
	//   2,1,1,2,1,1,2,1,1,2,1,1,2,1,1
	blocks := ""
	for i := 0; i < UNFOLD_FACTOR; i++ {
		blocks += parts[1]
		if i != UNFOLD_FACTOR-1 {
			blocks += ","
		}
	}
	contiguousDamagedSprings := []int{}
	// Parse the blocks.
	for _, groupSizeStr := range strings.Split(blocks, ",") {
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
		validRows := row.ValidAlternativeRows(false)
		hotsprings.CacheClear()

		sum += validRows
	}

	fmt.Printf("Sum: %d\n", sum)
}
