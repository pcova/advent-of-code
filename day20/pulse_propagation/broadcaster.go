package pulsepropagation

type Broadcaster struct {
	Destinations []Module
}

func (b *Broadcaster) Name() string {
	return "broadcaster"
}

func (b *Broadcaster) Type() string {
	return "broadcaster"
}

func (b *Broadcaster) AddDestination(destination Module) {
	b.Destinations = append(b.Destinations, destination)
}

func (b *Broadcaster) ProcessPulse(_ Module, pulse Pulse) []Event {
	events := make([]Event, 0)
	for _, destination := range b.Destinations {
		events = append(events, Event{pulse, destination, b})
	}

	return events
}
