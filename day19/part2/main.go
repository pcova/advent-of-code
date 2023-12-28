package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/pcova/advent-of-code/day19/parts"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	// Parse workflows
	var workflows parts.Workflows = make(parts.Workflows)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		name, workflow := parts.ParseWorkflow(line)

		workflows[name] = *workflow
	}

	combinations := workflows.PossibleCombinations()

	fmt.Printf("Combinations: %d\n", combinations)
}
