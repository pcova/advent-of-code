package hotsprings

type Spring rune

const (
	Operational Spring = '.'
	Damaged     Spring = '#'
	Unknown     Spring = '?'
)

func (s Spring) IsOperational() bool {
	return s == Operational
}

func (s Spring) IsDamaged() bool {
	return s == Damaged
}

func (s Spring) IsUnknown() bool {
	return s == Unknown
}
