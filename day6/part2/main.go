package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pcova/advent-of-code/day6/race"
)

func parseInput() (int, int) {
	var time, distance int

	scanner := bufio.NewScanner(os.Stdin)

	// we want to scan the first line representing the time allowed
	scanner.Scan()

	timeStr := strings.ReplaceAll(strings.TrimPrefix(scanner.Text(), "Time: "), " ", "")
	time, err := strconv.Atoi(timeStr)
	if err != nil {
		panic(fmt.Errorf("invalid input: %v", err))
	}

	// we want to scan the second line representing the record distance
	scanner.Scan()

	distanceStr := strings.ReplaceAll(strings.TrimPrefix(scanner.Text(), "Distance: "), " ", "")
	distance, err = strconv.Atoi(distanceStr)
	if err != nil {
		panic(fmt.Errorf("invalid input: %v", err))
	}

	return time, distance
}

func main() {
	time, distance := parseInput()

	race := race.NewRace(time, distance)

	fmt.Printf("The margin of error is %d\n", race.CountWaysToBeatRecord())
}
