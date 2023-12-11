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

	// visualize the image
	// println("Before expansion:")
	// println(image.String())

	// expand the image by adding empty rows and columns by each row or column that is empty
	image.Expand(cosmos.EXPENSION_RATE_1)

	// visualize the image
	// println("After expansion:")
	// println(image.String())

	// get all galaxies
	galaxies := image.GetGalaxies(cosmos.NO_EXPENSION)

	fmt.Printf("Total distance: %d\n", galaxies.CalculateDistance())
}
