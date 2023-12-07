package scratchcards

type Game struct {
	cards map[int]*Card
}

func NewGame() *Game {
	return &Game{
		cards: make(map[int]*Card),
	}
}

func (g *Game) AddCard(card *Card) {
	g.cards[card.Index] = card
}

/**
 * Returns the total number of scratchcards you get
 */
func (g *Game) ProcessCards() int {
	sum := 0

	for i := 0; i < len(g.cards); i++ {
		card := g.cards[i+1]

		/**
		 * Increase the number of copies of the next card by the number of
		 * copies of this card plus one (the original card)
		 */
		for j := 0; j < card.nMatchingNumbers(); j++ {
			nextCard := g.cards[card.Index+j+1]
			nextCard.IncreaseCopiesN(card.copies + 1)
		}

		// Add the number of copies of this card plus one (the original card) to the total
		sum += card.copies + 1
	}

	return sum
}
