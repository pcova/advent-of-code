package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/pcova/advent-of-code/day5/farm"
)

// GetLowestLocationForRange returns the lowest location for a given range of seeds
func GetLowestLocationForRange(farm *farm.Farm, seedRange *farm.Range) int {
	locationRanges := farm.LocationForRange(seedRange)

	if len(locationRanges) == 0 {
		return -1
	}

	// Get lowest location for all location ranges
	lowestLocation := locationRanges[0].Start()
	for _, locationRange := range locationRanges[1:] {
		if locationRange.Start() < lowestLocation {
			lowestLocation = locationRange.Start()
		}
	}

	return lowestLocation
}

func main() {
	// Read stdin line by line
	scanner := bufio.NewScanner(os.Stdin)
	f, err := farm.ParseFarm(scanner)
	if err != nil {
		log.Fatal(err)
	}

	seeds := f.Seeds()
	if len(seeds) == 0 {
		fmt.Println("Lowest location: 0")
		return
	}

	// Get lowest location for each range
	var lowestLocations []int
	for i := 0; i < len(seeds); i += 2 {
		seedRange := farm.NewRange(seeds[i], seeds[i+1])

		lowestLocationForRange := GetLowestLocationForRange(f, seedRange)
		if lowestLocationForRange == -1 {
			log.Fatalf("Failed to get lowest location for range %v", seedRange)
		}

		lowestLocations = append(lowestLocations, lowestLocationForRange)
	}

	if len(lowestLocations) == 0 {
		log.Fatal("Failed to get lowest location for any range")
	}

	// Get lowest location for all ranges
	lowestLocation := lowestLocations[0]
	for _, location := range lowestLocations[1:] {
		if location < lowestLocation {
			lowestLocation = location
		}
	}

	fmt.Printf("Lowest location: %d\n", lowestLocation)
}
