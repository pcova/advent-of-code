package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/pcova/advent-of-code/day9/oasis"
)

func main() {
	sum := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		seq := oasis.ParseSequence(line)

		nextValue := seq.NextValuePrediction()

		fmt.Printf("Sequence: %v Next value: %d\n", seq, nextValue)

		sum += nextValue
	}

	fmt.Printf("Sum: %d\n", sum)
}
