package cubeconundrum

import (
	"fmt"
	"strconv"
	"strings"
)

type Game struct {
	index    int
	cubeSets []CubeSet
}

func ParseGameString(line string) (*Game, error) {
	cubeSets := make([]CubeSet, 0)

	gameAttr := strings.Split(line, ": ")
	if len(gameAttr) != 2 {
		return nil, fmt.Errorf("error parsing game attributes: %v", gameAttr)
	}

	gameHeader := strings.Split(gameAttr[0], " ")
	if len(gameHeader) != 2 {
		return nil, fmt.Errorf("error parsing game header: %v", gameHeader)
	}

	gameIndex, err := strconv.Atoi(gameHeader[1])
	if err != nil {
		return nil, fmt.Errorf("error converting game index to int: %v", err)
	}

	gameBody := gameAttr[1]
	for _, setStr := range strings.Split(gameBody, "; ") {
		cubeSet, err := ParseCubeSetString(setStr)
		if err != nil {
			return nil, fmt.Errorf("error getting cube set: %v", err)
		}
		cubeSets = append(cubeSets, *cubeSet)
	}

	return &Game{cubeSets: cubeSets, index: gameIndex}, nil
}

func (g *Game) String() string {
	return fmt.Sprintf("Game %d: %v", g.index, g.cubeSets)
}

func (g *Game) Validate(bag CubeSet) bool {
	for _, cubeSet := range g.cubeSets {
		if !cubeSet.Validate(bag) {
			return false
		}
	}
	return true
}

func (g *Game) Index() int {
	return g.index
}

func (g *Game) MinimalBag() *CubeSet {
	if len(g.cubeSets) == 0 {
		return &CubeSet{red: 0, green: 0, blue: 0}
	}

	minimalBag := g.cubeSets[0]
	for _, cubeSet := range g.cubeSets {
		if cubeSet.red > minimalBag.red {
			minimalBag.red = cubeSet.red
		}
		if cubeSet.green > minimalBag.green {
			minimalBag.green = cubeSet.green
		}
		if cubeSet.blue > minimalBag.blue {
			minimalBag.blue = cubeSet.blue
		}
	}
	return &minimalBag
}
