package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/pcova/advent-of-code/day11/cosmos"
)

func ParseImage(scanner *bufio.Scanner) *cosmos.Image {
	image := cosmos.Image{}
	for scanner.Scan() {
		line := scanner.Text()
		image = append(image, []rune(line))
	}
	return &image
}

func main() {
	image := ParseImage(bufio.NewScanner(os.Stdin))

	// get all galaxies
	galaxies := image.GetGalaxies(cosmos.EXPENSION_RATE_2)

	fmt.Printf("Total distance: %d\n", galaxies.CalculateDistance())
}
