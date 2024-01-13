package main

import (
	"bufio"
	"fmt"
	"os"

	pg "github.com/pcova/advent-of-code/day20/pulse_propagation"
)

func parseModules() map[string]pg.Module {
	scanner := bufio.NewScanner(os.Stdin)

	destinations := make(map[string][]string)
	modules := make(map[string]pg.Module)
	for scanner.Scan() {
		line := scanner.Text()

		module, mDestinations := pg.ParseModule(line)
		modules[module.Name()] = module
		destinations[module.Name()] = mDestinations
	}

	for name, module := range modules {
		switch module.Type() {
		case "flipflop", "conjunction", "broadcaster":
			for _, destinationName := range destinations[name] {
				destination, ok := modules[destinationName]
				if !ok {
					destination = pg.CreateMockModule(destinationName)
				}

				if destination.Type() == "conjunction" {
					destination.(*pg.Conjunction).AddSource(module)
				}

				module.AddDestination(destination)
			}
		}
	}

	return modules
}

func main() {
	modules := parseModules()

	broadcaster := modules["broadcaster"].(*pg.Broadcaster)
	button := pg.CreateButton(broadcaster)

	queue := make(pg.Queue, 0)
	countHigh := 0
	countLow := 0

	for i := 0; i < 1000; i++ {
		queue.Push(button.SendPulse(pg.LOW))

		for !queue.Empty() {
			event := queue.Pop()
			events := event.Destination().ProcessPulse(event.Source(), event.Pulse())

			if event.Pulse() == pg.HIGH {
				countHigh++
			} else {
				countLow++
			}

			for _, e := range events {
				queue.Push(e)
			}
		}
	}

	fmt.Printf("High: %d, Low: %d\n", countHigh, countLow)

	fmt.Printf("Product: %d\n", countHigh*countLow)
}
