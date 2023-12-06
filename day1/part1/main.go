package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func getFirstDigit(input []rune) int {
	for i := 0; i < len(input); i++ {
		c := input[i]
		if unicode.IsDigit(c) {
			return int(c - '0')
		}
	}
	return 0
}

func getLastDigit(input []rune) int {
	for i := len(input) - 1; i >= 0; i-- {
		c := input[i]
		if unicode.IsDigit(c) {
			return int(c - '0')
		}
	}
	return 0
}

func main() {
	sum := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		characters := []rune(line)
		firstDigit := getFirstDigit(characters)
		lastDigit := getLastDigit(characters)

		sum += firstDigit*10 + lastDigit
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	fmt.Println(sum)
}
