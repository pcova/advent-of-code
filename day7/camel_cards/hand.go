package camelcards

import "fmt"

type HandType int

const (
	HighCard HandType = iota
	OnePair
	TwoPairs
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func (ht HandType) String() string {
	switch ht {
	case HighCard:
		return "High Card"
	case OnePair:
		return "One Pair"
	case TwoPairs:
		return "Two Pairs"
	case ThreeOfAKind:
		return "Three of a Kind"
	case FullHouse:
		return "Full House"
	case FourOfAKind:
		return "Four of a Kind"
	case FiveOfAKind:
		return "Five of a Kind"
	default:
		panic("invalid hand type")
	}
}

type Hand struct {
	cards []Card
	bid   int
}

func NewHand(cards []Card, bid int) *Hand {
	return &Hand{cards: cards, bid: bid}
}

func (h *Hand) Bid() int {
	return h.bid
}

func (h *Hand) Type() HandType {
	var counts map[int]int = make(map[int]int)
	var nJokers int = 0

	for _, c := range h.cards {
		if c == Joker {
			nJokers++
			continue
		}
		counts[int(c)]++
	}

	var nPairs int = 0
	var nTriples int = 0
	var nQuadruples int = 0
	var nQuintuples int = 0

	for _, count := range counts {
		switch count {
		case 2:
			nPairs++
		case 3:
			nTriples++
		case 4:
			nQuadruples++
		case 5:
			nQuintuples++
		}
	}

	// logic for part 2 (jokers)
	switch nJokers {
	case 5, 4:
		return FiveOfAKind
	case 3:
		if nPairs == 1 {
			return FiveOfAKind
		}
		return FourOfAKind
	case 2:
		if nTriples == 1 {
			return FiveOfAKind
		}
		if nPairs == 1 {
			return FourOfAKind
		}
		return ThreeOfAKind
	case 1:
		if nQuadruples == 1 {
			return FiveOfAKind
		}
		if nTriples == 1 {
			return FourOfAKind
		}
		if nPairs == 2 {
			return FullHouse
		}
		if nPairs == 1 {
			return ThreeOfAKind
		}
		return OnePair
	}

	// if there are no jokers
	if nQuintuples == 1 {
		return FiveOfAKind
	}

	if nQuadruples == 1 {
		return FourOfAKind
	}

	if nTriples == 1 && nPairs == 1 {
		return FullHouse
	}

	if nTriples == 1 {
		return ThreeOfAKind
	}

	if nPairs == 2 {
		return TwoPairs
	}

	if nPairs == 1 {
		return OnePair
	}

	return HighCard
}

// Compare returns a positive number if h is better than other, a negative number if h is worse than other, and 0 if they are equal
func (h *Hand) Compare(other *Hand) int {
	if h.Type() != other.Type() {
		return int(h.Type()) - int(other.Type())
	}

	// if the hands are of the same type, we need to compare the cards
	for i := 0; i < len(h.cards); i++ {
		if int(h.cards[i]) != int(other.cards[i]) {
			return int(h.cards[i]) - int(other.cards[i])
		}
	}

	return 0
}

func (h *Hand) String() string {
	return fmt.Sprintf("{%v, bid=%d, type=%v}", h.cards, h.bid, h.Type())
}
