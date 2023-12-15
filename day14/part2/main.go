package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/pcova/advent-of-code/day14/platform"
)

func ParsePlatform(scanner *bufio.Scanner) *platform.Platform {
	var p = platform.Platform{}
	for scanner.Scan() {
		line := scanner.Text()
		var row []platform.Space
		for _, char := range line {
			row = append(row, platform.Space(char))
		}
		p.AddRow(row)
	}
	return &p
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	p := ParsePlatform(scanner)

	nCycles := 1000000000

	iterationsIndexes := map[string]int{}
	iterations := []*platform.Platform{}
	loopStart := -1
	loopEnd := -1
	for i := 0; i < nCycles; i++ {
		// if current iteration is already in the map, we found a loop
		if prev, ok := iterationsIndexes[p.String()]; ok {
			loopStart = prev
			loopEnd = i
			break
		}

		// otherwise, add it to the map
		iterationsIndexes[p.String()] = i
		iterations = append(iterations, p.Copy())

		// and spin cycle again
		p.SpinCycle()
	}

	// if we found a loop, we can skip the iterations that repeat
	// and just calculate the rest
	loopLength := loopEnd - loopStart
	restLoops := (nCycles - loopStart) % loopLength

	// get the platform at the end of nCycles iterations
	p = iterations[loopStart+restLoops]

	fmt.Println(p.TotalLoad())
}
