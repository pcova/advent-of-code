package pulsepropagation

type Pulse bool

const (
	HIGH Pulse = true
	LOW  Pulse = false
)

func (p Pulse) String() string {
	if p == HIGH {
		return "high"
	}

	return "low"
}
