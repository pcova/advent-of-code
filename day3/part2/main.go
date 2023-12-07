package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/pcova/advent-of-code/day3/engine"
)

func LoadMap() *engine.Schematic {
	m := &engine.Schematic{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		m.AppendLine(line)
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	return m
}

func main() {
	sum := 0

	m := LoadMap()

	for _, p := range m.GetSymbols() {
		symbol := m.GetTile(p)
		if symbol == '*' {
			adjacentNumbers, err := m.GetNumbersAdjacentTo(p)
			if err != nil {
				log.Fatal(err)
			}

			// check if it is a "gear"
			if len(adjacentNumbers) == 2 {
				sum += adjacentNumbers[0] * adjacentNumbers[1]
			}
		}
	}

	fmt.Println(sum)
}
