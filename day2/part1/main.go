package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	cubeconundrum "github.com/pcova/advent-of-code/day2/cube_conundrum"
)

var BAG = cubeconundrum.CreateCubeSet(12, 13, 14)

func main() {
	sum := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		game, err := cubeconundrum.ParseGameString(line)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%s -> ", game.String())
		if game.Validate(BAG) {
			fmt.Println("VALID")
			sum += game.Index()
		} else {
			fmt.Println("INVALID")
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Sum of valid games: %d\n", sum)
}
