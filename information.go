package main

import (
	"bufio"
	"fmt"
	"os"
)

func Information(numbers []string, brandsFilePath, issuersFilePath string) {
	// TODO: написать комменты
	// content of file - array of strings
	brandsContent := readFile(brandsFilePath)
	issuersContent := readFile(issuersFilePath)

	// map that created via content of file
	brandsMap := createMap(brandsContent)
	issuersMap := createMap(issuersContent)

	// lengths of credit card numbers of the brands
	numberLengths := map[string][]int{
		"VISA":       {13, 14, 15, 16},
		"MASTERCARD": {16},
		"AMEX":       {15},
		"DINERSCLUB": {14, 16},
	}

	// processing every number via for loop
	for _, number := range numbers {
		identifiedBrand := identifyValue(brandsMap, number)
		identifiedIssuer := identifyValue(issuersMap, number)
		isCorrectNumber := false

		// length checker
		for _, correctLength := range numberLengths[identifiedBrand] {
			if len(number) == correctLength {
				isCorrectNumber = true
				break
			}
		}

		isCorrectNumber = isCorrectNumber && LuhnTest(number)
		printInformation(number, isCorrectNumber, identifiedBrand, identifiedIssuer)
	}
}

// prints information about the card number
// ! Не принтит через stdout
func printInformation(number string, isCorrect bool, brand string, issuer string) {
	fmt.Println(number)
	if isCorrect {
		fmt.Println("Correct: yes")
		fmt.Println("Card Brand: " + brand)
		fmt.Println("Card Issuer: " + issuer)
	} else {
		fmt.Println("Correct: no")
		fmt.Println("Card Brand: -")
		fmt.Println("Card Issuer: -")
	}
	fmt.Println()
}

func identifyValue(brandsMap map[string][]string, number string) string {
	if len(number) <= 0 {
		os.Exit(1)
	}

	for brand, codes := range brandsMap {
		for _, code := range codes {
			for i := 0; i < len(code); i++ {
				if len(code) <= 0 {
					break
				}

				if code[i] != number[i] {
					break
				} else if i == len(code)-1 {
					return brand
				}
			}
		}
	}

	return "-"
}

// opens and reads lines from the file and returns array of strings
// TODO: Дописать error handling
func readFile(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	content, err := readLines(file)
	if err != nil {
		os.Exit(1)
	}
	return content
}

// TODO: Разобраться как работает эта функция
func readLines(file *os.File) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

// modifies array of the strings (content) to map with key and value via Split() func
// TODO: Оптимизировать и сделать код читабельным
func createMap(content []string) map[string][]string {
	m := make(map[string][]string)

	var splittedLine []string
	for _, line := range content {
		if len(line) <= 0 || line == "" || line == "\n" {
			continue
		}

		splittedLine = Split(line, ':')
		if len(splittedLine) == 0 {
			continue
		}
		m[splittedLine[0]] = append(m[splittedLine[0]], splittedLine[1])
	}

	return m
}

// splits and trim spaces, returns an array of the strings
func Split(s string, separator byte) []string {
	var result []string
	currentWord := ""

	for i := 0; i < len(s); i++ {
		if s[i] != separator && len(s) != 0 {
			currentWord += string(s[i])
		} else {
			result = append(result, currentWord)
			currentWord = ""
		}
	}
	if currentWord != "" {
		result = append(result, currentWord)
	}

	return result
}
