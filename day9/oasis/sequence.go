package oasis

import (
	"strconv"
	"strings"
)

type Sequence []int

/**
 * ParseSequence parses a string of space-separated integers into a Sequence.
 *
 * Example:
 *   input := "1 2 3 4 5"
 *   seq := ParseSequence(input)
 *   fmt.Println(seq) // [1 2 3 4 5]
 */
func ParseSequence(input string) *Sequence {
	var seq Sequence

	for _, value := range strings.Split(input, " ") {
		number, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}

		seq.Add(number)
	}

	return &seq
}

func (s *Sequence) Add(value int) {
	*s = append(*s, value)
}

// Variance returns a sequence of variances between each value in the sequence.
func (s Sequence) Variance() *Sequence {
	var variances Sequence

	for i := 0; i < len(s)-1; i++ {
		variances.Add(s[i+1] - s[i])
	}

	return &variances
}

// IsZeroSequence returns true if all values in the sequence are zero.
func (s Sequence) IsZeroSequence() bool {
	isAllZero := true
	for _, value := range s {
		if value != 0 {
			isAllZero = false
			break
		}
	}

	return isAllZero
}

// NextValuePrediction returns the next value in the sequence, based on the
// previous values.
func (s Sequence) NextValuePrediction() int {
	if s.IsZeroSequence() {
		return 0
	}

	nextVariance := s.Variance()

	return s[len(s)-1] + nextVariance.NextValuePrediction()
}

// PreviousValueExtrapolation returns the previous value in the sequence, based
// on the next values.
func (s Sequence) PreviousValueExtrapolation() int {
	if s.IsZeroSequence() {
		return 0
	}

	nextVariance := s.Variance()

	return s[0] - nextVariance.PreviousValueExtrapolation()
}
