package main

import (
	"bufio"
	"os"

	"github.com/pcova/advent-of-code/day15/lens"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// Set the split function for the scanning operation.
	scanner.Split(lens.ScanStep)

	sum := 0
	for scanner.Scan() {
		step := scanner.Text()
		sum += lens.Hash(step)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	println(sum)
}
