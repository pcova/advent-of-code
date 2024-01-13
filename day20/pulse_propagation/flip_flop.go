package pulsepropagation

type FlipFlop struct {
	name         string
	state        bool
	Destinations []Module
}

func (f *FlipFlop) Name() string {
	return f.name
}

func (f *FlipFlop) Type() string {
	return "flipflop"
}

func (f *FlipFlop) AddDestination(destination Module) {
	f.Destinations = append(f.Destinations, destination)
}

func (f *FlipFlop) ProcessPulse(_ Module, pulse Pulse) []Event {
	if pulse == HIGH {
		return []Event{}
	}

	f.state = !f.state

	events := make([]Event, 0)
	for _, destination := range f.Destinations {
		events = append(events, Event{Pulse(f.state), destination, f})
	}

	return events
}
