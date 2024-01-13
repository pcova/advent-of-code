package pulsepropagation

import "strings"

type Module interface {
	Name() string
	Type() string
	AddDestination(Module)
	ProcessPulse(Module, Pulse) []Event
}

type MockModule struct {
	name string
}

func (m *MockModule) Name() string {
	return m.name
}

func (m *MockModule) Type() string {
	return "mock"
}

func (m *MockModule) AddDestination(_ Module) {
	panic("mock cannot have destinations")
}

func (m *MockModule) ProcessPulse(_ Module, pulse Pulse) []Event {
	return []Event{}
}

func CreateMockModule(name string) *MockModule {
	return &MockModule{name: name}
}

func ParseModule(line string) (Module, []string) {
	parts := strings.Split(line, "->")
	if len(parts) != 2 {
		panic("invalid module")
	}

	moduleName := strings.Trim(parts[0], " ")
	destinations := make([]string, 0)
	for _, part := range strings.Split(parts[1], ",") {
		nextModule := strings.Trim(part, " ")
		if nextModule == "" {
			continue
		}

		destinations = append(destinations, nextModule)
	}

	switch moduleName[0] {
	case '%':
		return &FlipFlop{name: moduleName[1:]}, destinations
	case '&':
		return &Conjunction{name: moduleName[1:]}, destinations
	case 'b':
		return &Broadcaster{}, destinations
	default:
		return &MockModule{name: moduleName}, destinations
	}
}
