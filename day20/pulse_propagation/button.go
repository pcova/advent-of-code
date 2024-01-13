package pulsepropagation

type Button struct {
	broadcaster *Broadcaster
}

func (b *Button) Name() string {
	return "button"
}

func (b *Button) Type() string {
	return "button"
}

func (b *Button) SendPulse(pulse Pulse) Event {
	return Event{pulse, b.broadcaster, b}
}

func (b *Button) AddDestination(_ Module) {
	panic("button cannot have destinations")
}

func (b *Button) ProcessPulse(_ Module, pulse Pulse) []Event {
	return []Event{b.SendPulse(pulse)}
}

func CreateButton(broadcaster *Broadcaster) *Button {
	return &Button{broadcaster: broadcaster}
}
