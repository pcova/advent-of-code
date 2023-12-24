package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/pcova/advent-of-code/day18/lavaduct"
)

func ParseDigPlan() *lavaduct.DigPlan {
	scanner := bufio.NewScanner(os.Stdin)

	var digPlan *lavaduct.DigPlan = new(lavaduct.DigPlan)
	for scanner.Scan() {
		line := scanner.Text()

		var direction rune
		var meters int
		fmt.Sscanf(line, "%c %d", &direction, &meters)

		digPlan.AddInstruction(direction, meters)
	}

	return digPlan
}

func main() {
	digPlan := ParseDigPlan()

	lagoon := digPlan.Lagoon()

	fmt.Printf("Lagoon volume: %d\n", lagoon.Volume())
}
