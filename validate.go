package main

import (
	"fmt"
	"os"
)

// checks the card number for invalid or fake and returns boolean
func LuhnTest(number string) bool {
	if len(number) < 13 {
		return false
	}

	sum := 0
	for i, ch := range number {
		digit, err := RuneToInt(ch)
		if err != nil {
			return false
		}

		if i%2 == 0 {
			if digit*2 >= 10 {
				sum += digit*2 - 9
			} else {
				sum += digit * 2
			}
		} else {
			sum += digit
		}
	}

	if sum%10 == 0 {
		return true
	}

	return false
}

// checks validity of number and returns correct number
func checkNumber(number string, withAsterisks bool) string {
	if len(number) == 0 {
		os.Exit(1)
	}

	newNumber := ""
	for i := 0; i < len(number); i++ {
		if number[i] == ' ' {
			continue
		}

		if number[i] < '0' || number[i] > '9' {
			if number[i] == '*' && !withAsterisks {
				os.Exit(1)
			}
		}

		newNumber += string(number[i])
	}

	if len(newNumber) < 13 || len(newNumber) > 19 {
		os.Exit(1)
	}
	return newNumber
}

// converts rune to int and returns it
func RuneToInt(ch rune) (int, error) {
	if int(ch) >= 58 && int(ch) <= 47 {
		return 0, fmt.Errorf("The number is not correct.")
	}
	return int(ch - 48), nil
}

// removes quotes of the string
func RemoveQuotes(str string) {
	if len(str) == 0 {
		return
	}
	if str[0] == '"' && str[len(str)-1] == '"' {
		str = str[1 : len(str)-1]
	}
}
