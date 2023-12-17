package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/pcova/advent-of-code/day15/lens"
)

type Operation int

const (
	Add Operation = iota
	Remove
)

/**
 * Parse a step string into a key, operation, and value.
 */
func parseStep(step string) (key string, operation Operation, value int) {
	// If the step ends with a '-', it's a remove operation.
	if step[len(step)-1] == '-' {
		key = step[:len(step)-1]
		operation = Remove
		return
	}

	// Otherwise, it's an add operation.
	parts := strings.Split(step, "=")
	key = parts[0]
	operation = Add

	// Parse the value.
	value, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}

	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// Set the split function for the scanning operation.
	scanner.Split(lens.ScanStep)

	boxes := map[int]*lens.Box{}
	for scanner.Scan() {
		step := scanner.Text()

		key, operation, value := parseStep(step)

		// Get the box for this key.
		hash := lens.Hash(key)

		// If the box doesn't exist, create it.
		box, ok := boxes[hash]
		if !ok {
			box = lens.NewBox(hash)
			boxes[hash] = box
		}

		// Perform the operation.
		if operation == Remove {
			box.RemoveLen(key)
			continue
		}

		if _, err := box.GetLen(key); err == nil {
			box.UpdateLen(key, value)
		} else {
			box.AddLen(key, value)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	sum := 0
	for _, box := range boxes {
		focusingPower := box.FocusingPower()
		sum += focusingPower
	}

	println(sum)
}
