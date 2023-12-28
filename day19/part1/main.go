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

	// Parse parts
	var partsList []parts.Part = make([]parts.Part, 0)
	for scanner.Scan() {
		line := scanner.Text()

		part := parts.ParsePart(line)

		partsList = append(partsList, part)
	}

	approvedParts := workflows.ApprovedParts(partsList)

	sum := 0
	for _, part := range approvedParts {
		sum += part.TotalRating()
	}

	fmt.Printf("Sum: %d\n", sum)
}
