package scratchcards

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Card struct {
	Index          int
	winingNumbers  []int
	numbersYouHave []int
	copies         int
}

func CreateCard(index int) *Card {
	return &Card{
		Index:          index,
		winingNumbers:  []int{},
		numbersYouHave: []int{},
		copies:         0,
	}
}

func (c *Card) IsWinningNumber(number int) bool {
	for _, n := range c.winingNumbers {
		if n == number {
			return true
		}
	}
	return false
}

/**
 * Returns the number of winning numbers you have
 */
func (c *Card) nMatchingNumbers() int {
	count := 0
	for _, n := range c.numbersYouHave {
		if c.IsWinningNumber(n) {
			count++
		}
	}
	return count
}

/**
 * Returns the number of points you get for this card
 */
func (c *Card) Points() int {
	count := c.nMatchingNumbers()

	if count == 0 {
		return 0
	}

	return int(math.Pow(2, float64(count-1)))
}

func (c *Card) IncreaseCopies() {
	c.copies++
}

func (c *Card) IncreaseCopiesN(n int) {
	c.copies += n
}

func parseNumbersStr(numbersStr string) ([]int, error) {
	numbers := []int{}
	for _, n := range strings.Split(numbersStr, " ") {
		if n == "" {
			continue
		}
		number, err := strconv.Atoi(n)
		if err != nil {
			return nil, fmt.Errorf("invalid number '%s'", n)
		}
		numbers = append(numbers, number)
	}
	return numbers, nil
}

/**
 * Parses a line from the input file and returns a Card
 *
 * Example line:
 * Card 1: 1 2 3 4 5 | 1 2 3 4 5
 * ------  --------------------- <- Card Body
 *  ^   -  ---------   ---------
 *  |   ^       ^             ^
 *  |   |       |             |
 *  |   |       |             +-- Numbers you have
 *  |   |       +-- Winning numbers
 *  |   +-- Card Index
 *  +-- Card Header
 */
func ParseCard(line string) (*Card, error) {
	// Split the line into the card header and body
	cardSplit := strings.Split(line, ": ")
	if len(cardSplit) != 2 {
		return nil, fmt.Errorf("invalid card line: %s", line)
	}

	// Parse the card header
	cardHeader := cardSplit[0]
	cardHeaderSplit := strings.Split(cardHeader, " ")
	if len(cardHeaderSplit) < 2 {
		return nil, fmt.Errorf("invalid card header: %s", cardHeader)
	}

	// Parse the card index from the card header
	cardIndex, err := strconv.Atoi(cardHeaderSplit[len(cardHeaderSplit)-1])
	if err != nil {
		return nil, fmt.Errorf("invalid card index '%s'", cardHeaderSplit[1])
	}

	card := CreateCard(cardIndex)

	cardBody := cardSplit[1]

	// Parse the card body
	cardBodySplit := strings.Split(cardBody, " | ")
	if len(cardBodySplit) != 2 {
		return nil, fmt.Errorf("invalid card body: %s", cardBody)
	}

	// Parse the winning numbers from the card body
	card.winingNumbers, err = parseNumbersStr(cardBodySplit[0])
	if err != nil {
		return nil, fmt.Errorf("error parsing winning numbers '%s': %v", cardBodySplit[0], err)
	}

	// Parse the numbers you have from the card body
	card.numbersYouHave, err = parseNumbersStr(cardBodySplit[1])
	if err != nil {
		return nil, fmt.Errorf("error parsing numbers you have '%s': %v", cardBodySplit[1], err)
	}

	return card, nil
}
