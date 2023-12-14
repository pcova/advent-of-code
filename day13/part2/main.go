package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/pcova/advent-of-code/day13/pattern"
)

func ParsePatterns(scanner *bufio.Scanner) []pattern.Pattern {
	patterns := make([]pattern.Pattern, 0)
	p := pattern.Pattern{}
	p.FixSmudge()
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			patterns = append(patterns, p)
			p = pattern.Pattern{}
			p.FixSmudge()
			continue
		}
		p.AddRow(line)
	}
	patterns = append(patterns, p)
	return patterns
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	patterns := ParsePatterns(scanner)

	sum := 0
	for _, pattern := range patterns {
		score := pattern.Score()
		sum += score
	}

	fmt.Printf("Sum: %d\n", sum)
}
