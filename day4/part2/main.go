package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/pcova/advent-of-code/day4/scratchcards"
)

func main() {
	game := scratchcards.NewGame()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		card, err := scratchcards.ParseCard(line)
		if err != nil {
			log.Fatalln(err)
		}

		game.AddCard(card)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Total scratchcards: %d\n", game.ProcessCards())
}
