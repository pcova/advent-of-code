package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/pcova/advent-of-code/day8/desert"
)

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func checkStartNode(node string) bool {
	return node[2] == 'A'
}

func checkEndNode(node string) bool {
	return node[2] == 'Z'
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	instructions := desert.ParseInstructions(scanner)

	// skip empty line
	scanner.Scan()

	m := desert.ParseNetwork(scanner)

	// find all starting nodes (i.e. nodes that end with A)
	pathsToFollow := make([]string, 0)
	for node := range m {
		if checkStartNode(node) {
			pathsToFollow = append(pathsToFollow, node)
		}
	}

	stepsByPath := make([]int, len(pathsToFollow))
	end := false
	for !end {
		end = true

		// for each path, follow the left/right instruction
		for i, currentNode := range pathsToFollow {

			// if we're already on an end node, skip this path
			if checkEndNode(currentNode) {
				continue
			}

			// it chooses the next node for this path based on the current left/right instruction
			currentInstruction := instructions[stepsByPath[i]%len(instructions)]
			pathsToFollow[i] = m[currentNode][int(currentInstruction)]

			stepsByPath[i]++

			// stop when all paths currently being followed are at end nodes (i.e. end with Z)
			end = end && checkEndNode(pathsToFollow[i])
		}
	}

	// the total number of steps is the LCM of the number of steps for each path
	steps := LCM(stepsByPath[0], stepsByPath[1], stepsByPath[2:]...)

	fmt.Printf("By following the left/right instructions, you end up entirely on nodes that end in Z after %v steps\n", steps)
}
