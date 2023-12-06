package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	cubeconundrum "github.com/pcova/advent-of-code/day2/cube_conundrum"
)

func main() {
	sum := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		game, err := cubeconundrum.ParseGameString(line)
		if err != nil {
			log.Fatalln(err)
		}

		minimalBag := game.MinimalBag()
		power := minimalBag.Power()
		fmt.Printf("%s -> minimal bag %v -> power %d\n", game.String(), minimalBag, power)

		sum += power
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Sum of valid games: %d\n", sum)
}
