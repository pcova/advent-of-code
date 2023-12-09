package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/pcova/advent-of-code/day6/race"
)

func parseInput() ([]int, []int) {
	var times, distances []int
	var p *[]int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()

		switch word {
		case "Time:":
			p = &times
		case "Distance:":
			p = &distances
		default:
			number, err := strconv.Atoi(word)
			if err != nil {
				panic(err)
			}
			*p = append(*p, number)
		}
	}

	if len(times) != len(distances) {
		panic("invalid input")
	}

	return times, distances
}

func main() {
	times, distances := parseInput()

	var product int = 1
	for i := 0; i < len(times); i++ {
		race := race.NewRace(times[i], distances[i])

		nWaysToBeatRecord := race.CountWaysToBeatRecord()

		fmt.Printf("Number of ways to beat the record on race %v is %d\n", race, nWaysToBeatRecord)

		product *= nWaysToBeatRecord
	}

	fmt.Printf("The margin of error is %d\n", product)
}
