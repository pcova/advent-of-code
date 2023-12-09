package camelcards

import "fmt"

type Card int

const (
	Joker Card = iota + 1 // for part 2
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack  // 11
	Queen // 12
	King  // 13
	Ace   // 14
)

func (c Card) String() string {
	if c >= Two && c <= Nine {
		return fmt.Sprintf("%d", c)
	}

	switch c {
	case Ten:
		return "T"
	case Jack, Joker:
		return "J"
	case Queen:
		return "Q"
	case King:
		return "K"
	case Ace:
		return "A"
	default:
		panic("invalid card")
	}
}
