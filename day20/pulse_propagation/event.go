package pulsepropagation

type Event struct {
	pulse       Pulse
	destination Module
	source      Module
}

func (e *Event) Pulse() Pulse {
	return e.pulse
}

func (e *Event) Destination() Module {
	return e.destination
}

func (e *Event) Source() Module {
	return e.source
}
