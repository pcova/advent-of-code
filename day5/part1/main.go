package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/pcova/advent-of-code/day5/farm"
)

func main() {
	// Read stdin line by line
	scanner := bufio.NewScanner(os.Stdin)
	farm, err := farm.ParseFarm(scanner)
	if err != nil {
		log.Fatal(err)
	}

	seeds := farm.Seeds()
	if len(seeds) == 0 {
		fmt.Println("Lowest location: 0")
		return
	}

	lowestLocation := farm.Location(seeds[0])
	for _, seed := range seeds[1:] {
		location := farm.Location(seed)
		if location < lowestLocation {
			lowestLocation = location
		}
	}

	fmt.Printf("Lowest location: %d\n", lowestLocation)
}
