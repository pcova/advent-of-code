package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"unicode"
)

var DIGITS = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func getSpelledWithLettersDigit(input string) int {
	for spelledDigit := range DIGITS {
		if match, _ := regexp.MatchString(".*"+spelledDigit+".*", input); match {
			return DIGITS[spelledDigit]
		}
	}
	return -1
}

func getFirstDigit(input []rune) int {
	var word string
	for i := 0; i < len(input); i++ {
		c := input[i]
		if unicode.IsDigit(c) {
			return int(c - '0')
		}
		word += string(c)
		if getSpelledWithLettersDigit(word) != -1 {
			return getSpelledWithLettersDigit(word)
		}
	}
	return -1
}

func getLastDigit(input []rune) int {
	var word string
	for i := len(input) - 1; i >= 0; i-- {
		c := input[i]
		if unicode.IsDigit(c) {
			return int(c - '0')
		}
		word = string(c) + word
		if getSpelledWithLettersDigit(word) != -1 {
			return getSpelledWithLettersDigit(word)
		}
	}
	return -1
}

func main() {
	sum := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		characters := []rune(line)
		firstDigit := getFirstDigit(characters)
		lastDigit := getLastDigit(characters)

		if firstDigit == -1 || lastDigit == -1 {
			log.Fatalln("Error: couldn't find digit", line)
		}

		sum += firstDigit*10 + lastDigit
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	fmt.Println(sum)
}
