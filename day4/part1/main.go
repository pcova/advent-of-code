package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/pcova/advent-of-code/day4/scratchcards"
)

func main() {
	sum := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		card, err := scratchcards.ParseCard(line)
		if err != nil {
			log.Fatalln(err)
		}

		sum += card.Points()
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Total points: %d\n", sum)
}
