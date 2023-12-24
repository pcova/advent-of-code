package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pcova/advent-of-code/day18/lavaduct"
)

func ParseDigPlan() *lavaduct.DigPlan {
	scanner := bufio.NewScanner(os.Stdin)

	var digPlan *lavaduct.DigPlan = new(lavaduct.DigPlan)
	for scanner.Scan() {
		line := scanner.Text()

		var d rune
		var m int
		var hex string
		fmt.Sscanf(line, "%c %d %s", &d, &m, &hex)

		hex = strings.Trim(hex[1:len(hex)-1], "#")

		direction := lavaduct.IntToDirection(int(hex[len(hex)-1] - '0'))
		meters, err := strconv.ParseInt(hex[:len(hex)-1], 16, 64)
		if err != nil {
			panic(err)
		}

		digPlan.AddInstruction(rune(direction), int(meters))
	}

	return digPlan
}

func main() {
	digPlan := ParseDigPlan()

	lagoon := digPlan.Lagoon()

	fmt.Printf("Lagoon volume: %d\n", lagoon.Volume())
}
