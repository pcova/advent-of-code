package main

import (
	"bufio"
	"fmt"
	"os"

	camelcards "github.com/pcova/advent-of-code/day7/camel_cards"
)

func parseCard(c rune) camelcards.Card {
	switch c {
	case '2':
		return camelcards.Two
	case '3':
		return camelcards.Three
	case '4':
		return camelcards.Four
	case '5':
		return camelcards.Five
	case '6':
		return camelcards.Six
	case '7':
		return camelcards.Seven
	case '8':
		return camelcards.Eight
	case '9':
		return camelcards.Nine
	case 'T':
		return camelcards.Ten
	case 'J':
		return camelcards.Joker // only difference with day7/part1/main.go, the logic is in day7/camel_cards/hand.go
	case 'Q':
		return camelcards.Queen
	case 'K':
		return camelcards.King
	case 'A':
		return camelcards.Ace
	default:
		panic("invalid card")
	}
}

func parseHand(handStr string) (*camelcards.Hand, error) {
	var cards []camelcards.Card = make([]camelcards.Card, 5)
	var bid int
	var cardsStr string

	_, err := fmt.Sscanf(handStr, "%s %d", &cardsStr, &bid)
	if err != nil {
		return nil, fmt.Errorf("invalid hand: %v", err)
	}

	for i, c := range cardsStr {
		cards[i] = parseCard(c)
	}

	return camelcards.NewHand(cards, bid), nil
}

func main() {
	var handSet *camelcards.HandSet = camelcards.NewHandSet()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		hand, err := parseHand(line)
		if err != nil {
			panic(err)
		}

		handSet.Add(hand)
	}

	for i := 0; i < handSet.Len(); i++ {
		hand := handSet.Get(i)
		fmt.Printf("Hand %d: %v\n", i, hand)
	}

	fmt.Println("\nSorting hands...")

	handSet.SortAsc()

	sum := 0
	for i := 0; i < handSet.Len(); i++ {
		hand := handSet.Get(i)
		fmt.Printf("Hand %d: %v\n", i, hand)
		sum += hand.Bid() * (i + 1)
	}

	fmt.Printf("\nThe sum of all bids is %d\n", sum)
}
