package pulsepropagation

type Conjunction struct {
	name         string
	lastPulses   map[Module]Pulse
	Destinations []Module
}

func (c *Conjunction) Name() string {
	return c.name
}

func (c *Conjunction) Type() string {
	return "conjunction"
}

func (c *Conjunction) AddDestination(destination Module) {
	c.Destinations = append(c.Destinations, destination)
}

func (c *Conjunction) AddSource(source Module) {
	if c.lastPulses == nil {
		c.lastPulses = make(map[Module]Pulse)
	}

	c.lastPulses[source] = LOW
}

func (c *Conjunction) ProcessPulse(source Module, pulse Pulse) []Event {
	c.lastPulses[source] = pulse

	allHigh := true
	for _, lastPulse := range c.lastPulses {
		if lastPulse == LOW {
			allHigh = false
			break
		}
	}

	events := make([]Event, 0)
	for _, destination := range c.Destinations {
		events = append(events, Event{Pulse(!allHigh), destination, c})
	}

	return events
}
