package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/pcova/advent-of-code/day8/desert"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	instructions := desert.ParseInstructions(scanner)

	// skip empty line
	scanner.Scan()

	m := desert.ParseNetwork(scanner)

	steps := 0
	currentNode := "AAA"

	// "ZZZ" is the last node
	for currentNode != "ZZZ" {
		// choose the next node based on the current left/right instruction
		currentInstruction := instructions[steps%len(instructions)]
		currentNode = m[currentNode][int(currentInstruction)]

		steps++
	}

	// fmt.Print("\n\n")
	fmt.Printf("By following the left/right instructions, you reach ZZZ in %d steps\n", steps)
}
