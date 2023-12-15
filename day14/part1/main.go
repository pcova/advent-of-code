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
	platform := ParsePlatform(scanner)
	platform.TiltUp()
	fmt.Println(platform.TotalLoad())
}
