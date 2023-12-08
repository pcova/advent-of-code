package farm

import (
	"bufio"
	"fmt"
	"strings"
)

type Farm struct {
	seeds                []int
	seed2Soil            *Mapper
	soil2Fertilizer      *Mapper
	fertilizer2Water     *Mapper
	water2Light          *Mapper
	light2Temperature    *Mapper
	temperature2Humidity *Mapper
	humidity2Location    *Mapper
}

// AddMapper adds a mapper to the farm based on its name
func (f *Farm) AddMapper(m *Mapper) {
	switch m.name {
	case "seed-to-soil":
		f.seed2Soil = m
	case "soil-to-fertilizer":
		f.soil2Fertilizer = m
	case "fertilizer-to-water":
		f.fertilizer2Water = m
	case "water-to-light":
		f.water2Light = m
	case "light-to-temperature":
		f.light2Temperature = m
	case "temperature-to-humidity":
		f.temperature2Humidity = m
	case "humidity-to-location":
		f.humidity2Location = m
	}
}

func (f *Farm) AddSeed(seed int) {
	f.seeds = append(f.seeds, seed)
}

func (f *Farm) Seeds() []int {
	return f.seeds
}

func (f *Farm) Location(seed int) int {
	// Apply mappers
	soil := f.seed2Soil.Apply(seed)
	fertilizer := f.soil2Fertilizer.Apply(soil)
	water := f.fertilizer2Water.Apply(fertilizer)
	light := f.water2Light.Apply(water)
	temperature := f.light2Temperature.Apply(light)
	humidity := f.temperature2Humidity.Apply(temperature)
	location := f.humidity2Location.Apply(humidity)

	return location
}

// LocationForRange maps a range of seeds to a range of locations
func (f *Farm) LocationForRange(seedRange *Range) []*Range {
	// Apply mappers
	soilRange := f.seed2Soil.ApplyToRange(seedRange)

	var fertilizerRange []*Range
	for _, sr := range soilRange {
		newRange := f.soil2Fertilizer.ApplyToRange(sr)
		fertilizerRange = append(fertilizerRange, newRange...)
	}

	var waterRange []*Range
	for _, fr := range fertilizerRange {
		waterRange = append(waterRange, f.fertilizer2Water.ApplyToRange(fr)...)
	}

	var lightRange []*Range
	for _, wr := range waterRange {
		lightRange = append(lightRange, f.water2Light.ApplyToRange(wr)...)
	}

	var temperatureRange []*Range
	for _, lr := range lightRange {
		temperatureRange = append(temperatureRange, f.light2Temperature.ApplyToRange(lr)...)
	}

	var humidityRange []*Range
	for _, tr := range temperatureRange {
		humidityRange = append(humidityRange, f.temperature2Humidity.ApplyToRange(tr)...)
	}

	var locationRange []*Range
	for _, hr := range humidityRange {
		locationRange = append(locationRange, f.humidity2Location.ApplyToRange(hr)...)
	}

	return locationRange
}

func NewFarm() *Farm {
	return &Farm{}
}

/**
 * parseSeeds parses a string of seeds
 *
 * Example:
 *  - input: "0 10 5 5"
 *  - seeds: [0 10 5 5]
 */
func parseSeeds(input string) ([]int, error) {
	var seeds []int
	for _, s := range strings.Split(input, " ") {
		var seed int
		_, err := fmt.Sscanf(s, "%d", &seed)
		if err != nil {
			return nil, fmt.Errorf("failed to parse seed: %w", err)
		}

		seeds = append(seeds, seed)
	}

	return seeds, nil
}

/**
 * ParseFarm parses a farm from a scanner
 *
 * Example:
 *  - input:
 *    seeds: 0 10 25 5
 *    seed-to-soil map:
 *    0 10 5
 *    soil-to-fertilizer map:
 *    10 0 5
 *    fertilizer-to-water map:
 *    0 10 5
 *	  10 6 3
 *    water-to-light map:
 *    10 0 5
 *    light-to-temperature map:
 *    0 10 5
 *    temperature-to-humidity map:
 *    10 0 5
 *    humidity-to-location map:
 *    0 10 5
 *  - result:
 *    seeds: [0 10 25 5]
 *    seed2soil: [(10-15)->-10->(0-5)]
 *    soil2fertilizer: [(0-5)->+10->(10-15)]
 *    fertilizer2water: [(10-15)->-10->(0-5) (6-9)->+4->(10-13)]
 *    water2light: [(0-5)->+10->(10-15)]
 *    light2temperature: [(10-15)->-10->(0-5)]
 *    temperature2humidity: [(0-5)->+10->(10-15)]
 *    humidity2location: [(10-15)->-10->(0-5)]
 */
func ParseFarm(s *bufio.Scanner) (*Farm, error) {
	farm := NewFarm()

	// Parse seeds
	if s.Scan() {
		seedsStr := strings.TrimPrefix(s.Text(), "seeds: ")

		var err error
		farm.seeds, err = parseSeeds(seedsStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse seeds: %w", err)
		}
	}

	// Parse mappers
	var mapperStr []string
	for s.Scan() {
		line := s.Text()

		// Empty line means we're done with this mapper
		if line == "" {
			// on the first line, we don't have a mapper to add
			if len(mapperStr) > 0 {
				// Parse mapper
				mapper, err := ParseMapper(mapperStr)
				if err != nil {
					return nil, fmt.Errorf("failed to parse mapper: %w", err)
				}
				farm.AddMapper(mapper)

				// reset mapperStr
				mapperStr = []string{}
			}
			continue
		}

		// Add line to mapperStr to be parsed later
		mapperStr = append(mapperStr, line)
	}

	// Check for errors in scanner (other than EOF)
	if err := s.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan input: %w", err)
	}

	// Parse last mapper if there is one
	// (if there is no empty line at the end of the input)
	if len(mapperStr) > 0 {
		mapper, err := ParseMapper(mapperStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse mapper: %w", err)
		}
		farm.AddMapper(mapper)
	}

	fmt.Printf("seeds: %v\n", farm.seeds)

	return farm, nil
}
