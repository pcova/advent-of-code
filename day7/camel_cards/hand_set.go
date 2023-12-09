package camelcards

import "sort"

type HandSet struct {
	hands []*Hand
}

func NewHandSet() *HandSet {
	return &HandSet{hands: make([]*Hand, 0)}
}

func (hs *HandSet) Len() int {
	return len(hs.hands)
}

func (hs *HandSet) Add(hand *Hand) {
	hs.hands = append(hs.hands, hand)
}

func (hs *HandSet) Get(i int) *Hand {
	return hs.hands[i]
}

// SortAsc sorts the hands in ascending order (i.e. the first hand is the worst)
func (hs *HandSet) SortAsc() {
	sort.Slice(hs.hands, func(i, j int) bool {
		return hs.hands[i].Compare(hs.hands[j]) < 0
	})
}
