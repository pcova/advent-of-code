package farm

import "fmt"

type Range struct {
	start  int
	length int
}

func NewRange(start, length int) *Range {
	return &Range{
		start:  start,
		length: length,
	}
}

func (r *Range) Start() int {
	return r.start
}

func (r *Range) End() int {
	return r.start + r.length - 1
}

func (r *Range) Contains(index int) bool {
	return index >= r.start && index < r.start+r.length
}

// Intersects returns the intersection between the range and the other range
func (r *Range) Intersects(other *Range) *Range {
	if r.start > other.End() || r.End() < other.start {
		return nil
	}

	intersectionStart := r.start
	if other.start > r.start {
		intersectionStart = other.start
	}

	intersectionEnd := r.End()
	if other.End() < r.End() {
		intersectionEnd = other.End()
	}

	return NewRange(intersectionStart, intersectionEnd-intersectionStart+1)
}

/**
 * NotIntersects returns a list of ranges that are the result of
 * removing the intersection between the range and the other range
 *
 * Example:
 *  - range: (0-10)
 *  - other: (2-4)
 *  - result: [(0-1), (5-10)]
 */
func (r *Range) NotIntersects(other *Range) []*Range {
	if r.start > other.End() || r.End() < other.start {
		return []*Range{r}
	}

	var ranges []*Range

	if r.start < other.start {
		ranges = append(ranges, NewRange(r.start, other.start-r.start))
	}

	if r.End() > other.End() {
		ranges = append(ranges, NewRange(other.End()+1, r.End()-other.End()))
	}

	return ranges
}

// MapperRange represents a range of values that can be converted
type MapperRange struct {
	source      *Range
	destination *Range
}

/**
 * ParseRange parses a mapper range from a string
 *
 * Example:
 *  - input: "0 10 5"
 *  - result: (10-15)->-10->(0-5)
 *
 *  - input: "10 0 5"
 *  - result: (0-5)->+10->(10-15)
 */
func ParseRange(input string) (*MapperRange, error) {
	var sourceStart, destinationStart, length int
	_, err := fmt.Sscanf(input, "%d %d %d", &destinationStart, &sourceStart, &length)
	if err != nil {
		return nil, err
	}

	return &MapperRange{
		source:      &Range{start: sourceStart, length: length},
		destination: &Range{start: destinationStart, length: length},
	}, nil
}

func (r *MapperRange) contains(index int) bool {
	return r.source.Contains(index)
}

// mapIndex maps a value from the source range to the destination range
func (r *MapperRange) mapIndex(index int) (int, error) {
	if !r.contains(index) {
		return 0, fmt.Errorf("index %d is not in range %d-%d", index, r.source.start, r.source.start+r.source.length)
	}

	return r.destination.start + (index - r.source.start), nil
}

// Mapper represents a mapping of values
type Mapper struct {
	name    string
	mRanges []*MapperRange
}

func NewMapper() *Mapper {
	return &Mapper{
		mRanges: make([]*MapperRange, 0),
	}
}

// ParseMapper parses a mapper from a list of strings
func ParseMapper(input []string) (*Mapper, error) {
	mapper := NewMapper()

	for i, line := range input {
		// The first line is the mapper name, not a range (e.g. "mapper1 map:")
		if i == 0 {
			_, err := fmt.Sscanf(line, "%s map:", &mapper.name)
			if err != nil {
				return nil, fmt.Errorf("failed to parse mapper name: %w", err)
			}
			continue
		}

		r, err := ParseRange(line)
		if err != nil {
			return nil, err
		}

		mapper.AddRange(r)
	}

	return mapper, nil
}

func (m *Mapper) AddRange(r *MapperRange) {
	m.mRanges = append(m.mRanges, r)
}

// Apply iterates over the mapper ranges and applies the mapping to the index
func (m *Mapper) Apply(index int) int {
	for _, r := range m.mRanges {
		mappedIndex, err := r.mapIndex(index)
		if err == nil {
			return mappedIndex
		}
	}

	return index
}

/**
 * ApplyToRange applies the mapper to a range
 * It returns a list of ranges that are the result of the mapping
 * If the range is not fully mapped, it will return the original range
 *
 * Example:
 *  - range: (0-10)
 *  - mapper: (0-10)->+1->(1-11)
 *  - result: [(1-10)]
 *
 *  - range: (0-10)
 *  - mapper: (2-4)->+12->(14-16), (5-10)->+1->(6-11)
 *  - result: [(14-16), (6-11), (0-1)]
 *
 */
func (m *Mapper) ApplyToRange(r *Range) []*Range {
	// remainingRanges is the list of ranges that are not yet mapped (initially, it's the range itself)
	remainingRanges := []*Range{r}

	// ranges is the list of mapped ranges
	var ranges []*Range
	for _, mr := range m.mRanges {
		var err error
		var start, end int

		// Find the intersection between the range and the mapper range
		intersection := mr.source.Intersects(r)
		if intersection == nil {
			continue
		}

		// Remove the intersection from the remaining ranges
		for i := 0; i < len(remainingRanges); i++ {
			remainingRange := remainingRanges[i]
			notIntersected := remainingRange.NotIntersects(intersection)

			if len(notIntersected) == 0 {
				remainingRanges = append(remainingRanges[:i], remainingRanges[i+1:]...)
			} else if i == len(remainingRanges)-1 {
				remainingRanges = append(remainingRanges[:i], notIntersected...)
			} else {
				remainingRanges = append(remainingRanges[:i], append(notIntersected, remainingRanges[i+1:]...)...)
			}
		}

		start, err = mr.mapIndex(intersection.start)
		if err != nil {
			start = intersection.Start()
		}

		end, err = mr.mapIndex(intersection.End())
		if err != nil {
			end = intersection.End()
		}

		// Add the mapped range to the list of ranges
		ranges = append(ranges, NewRange(start, end-start+1))
	}

	// Add remaining ranges that are not mapped by any mapper range
	ranges = append(ranges, remainingRanges...)

	return ranges
}

func (m *Mapper) String() string {
	return fmt.Sprintf("%s: %+v", m.name, m.mRanges)
}
