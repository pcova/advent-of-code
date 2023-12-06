package cubeconundrum

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type CubeSet struct {
	red   int
	green int
	blue  int
}

func CreateCubeSet(red int, green int, blue int) CubeSet {
	return CubeSet{red: red, green: green, blue: blue}
}

func ParseCubeSetString(setStr string) (*CubeSet, error) {
	cubeSet := CubeSet{red: 0, green: 0, blue: 0}

	for _, cubeStr := range strings.Split(setStr, ", ") {
		cubeAttributes := strings.Split(cubeStr, " ")
		if len(cubeAttributes) != 2 {
			return nil, fmt.Errorf("error parsing cube attributes: %v", cubeAttributes)
		}

		qty, err := strconv.Atoi(cubeAttributes[0])
		if err != nil {
			return nil, fmt.Errorf("error converting cube quantity to int: %v", err)
		}

		switch cubeAttributes[1] {
		case "red":
			cubeSet.red = qty
		case "green":
			cubeSet.green = qty
		case "blue":
			cubeSet.blue = qty
		default:
			log.Fatalln("Unknown cube color")
		}
	}

	return &cubeSet, nil
}

func (c *CubeSet) Validate(bag CubeSet) bool {
	return c.red <= bag.red && c.green <= bag.green && c.blue <= bag.blue
}

func (c *CubeSet) Power() int {
	return c.red * c.green * c.blue
}
