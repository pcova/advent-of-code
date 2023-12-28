package parts

type Range struct {
	Min, Max int
}

func (r Range) PossibleCombinations() int {
	return r.Max - r.Min + 1
}
