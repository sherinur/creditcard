package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// generates combinations possible credit card numbers
// by replacing asterisks (*) with digits
// uses pickFlag variable to support --pick flag to randomly pick a single entry
func Generate(cardTemplate string, pickFlag bool) {
	cardTemplate = checkNumber(cardTemplate, true)
	// error handling
	asterisks, err := checkAsterisks(cardTemplate)
	if err != nil {
		os.Exit(1)
	}

	brandLengths := map[string][]int{
		"4":  {13, 14, 15, 16},
		"51": {16},
		"52": {16},
		"53": {16},
		"54": {16},
		"55": {16},
		"34": {15},
		"37": {15},
		"30": {14, 16},
	}

	// length check for the new number
	isCorrectLen := false
	var prefixOfNumber string

	for key := range brandLengths {
		if strings.HasPrefix(cardTemplate, key) {
			prefixOfNumber = key
			break
		}
	}

	if _, exists := brandLengths[prefixOfNumber]; exists {
		for _, correctLength := range brandLengths[prefixOfNumber] {
			if len(cardTemplate) == correctLength {
				isCorrectLen = true
			}
		}
	} else {
		isCorrectLen = true
	}

	if !isCorrectLen {
		os.Exit(1)
	}

	// logic
	base := cardTemplate[0 : len(cardTemplate)-asterisks]
	maxCombinations := powOfTen(asterisks) // number of max combinations
	var combinations []string

	var newCombination string
	for i := 0; i < maxCombinations; i++ {
		newCombination = base + padWithZeros(i, asterisks)

		if LuhnTest(newCombination) {
			combinations = append(combinations, newCombination)
			if !pickFlag {
				fmt.Fprintln(os.Stdout, newCombination)
			}
		}
	}

	if pickFlag {
		randomIndex := rand.Intn(len(combinations))
		fmt.Fprintln(os.Stdout, combinations[randomIndex])
	}
}

// counts asterisks and checks for correct number
// returns the number of asterisks and error if it is
func checkAsterisks(cardTemplate string) (int, error) {
	if len(cardTemplate) <= 0 {
		os.Exit(1)
	}

	counter := 0
	isInterrupted := false

	for i := len(cardTemplate) - 1; i >= 0; i-- {
		if counter > 4 {
			return counter, fmt.Errorf("The number of asterisks more than 4.")
		}

		if cardTemplate[i] == '*' {
			if isInterrupted {
				// error handling
				return 0, fmt.Errorf("Asterisks should be at the end of the number.")
			}
			counter++
		} else if cardTemplate[i] >= 48 && cardTemplate[i] <= 57 {
			isInterrupted = true
		} else {
			return 0, fmt.Errorf("The number is not correct.")
		}
	}
	return counter, nil
}

// returns power of ten
func powOfTen(num int) int {
	result := 1
	for i := 0; i < num; i++ {
		result *= 10
	}
	return result
}

// takes int and returns string in which number is padded with zeros
func padWithZeros(number, asterisks int) string {
	return fmt.Sprintf("%0*d", asterisks, number)
}
