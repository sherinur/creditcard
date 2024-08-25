package main

import (
	"math/rand"
	"os"
	"strings"
)

// The issue feature generates a random valid credit card number
// for a specified brand and issuer.
func Issue(brandsFilePath, issuersFilePath, brand, issuer string) {
	// content of file - array of strings
	brandsContent := readFile(brandsFilePath)
	issuersContent := readFile(issuersFilePath)

	// map that created via content of file
	brandsMap := createMap(brandsContent)
	issuersMap := createMap(issuersContent)

	// error handling for maps and contents
	// TODO: Дописать
	if len(brandsContent) <= 0 || len(issuersContent) <= 0 {
		os.Exit(1)
	}

	// map of possible credit card number lenghts of brands
	numberLengths := map[string][]int{
		"VISA":       {13, 14, 15, 16},
		"MASTERCARD": {16},
		"AMEX":       {15},
		"DINERSCLUB": {14, 16},
	}

	// checking does brand support issuer
	isCorrectIssuer := false
	issuersArray, exists := issuersMap[issuer]
	if !exists || len(issuersArray) == 0 {
		os.Exit(1)
	}

	for _, brandCode := range brandsMap[brand] {
		if checkIssuer(brandCode, issuersMap[issuer][0]) {
			isCorrectIssuer = true
			break
		}
	}
	if !isCorrectIssuer {
		os.Exit(1)
	}

	// random generating credit card number
	var newNumberLen int
	// generating length for the credit card number
	if _, exists := numberLengths[brand]; exists {
		randomIndex := rand.Intn(len(numberLengths[brand]))
		newNumberLen = numberLengths[brand][randomIndex]
	} else {
		newNumberLen = 16
	}

	rangeOfRandom := powOfTen(newNumberLen-10) - 1
	randomIssuerCode := issuersMap[issuer][rand.Intn(len(issuersMap[issuer]))]
	newNumber := randomIssuerCode + padWithZeros((rand.Intn(rangeOfRandom)), newNumberLen-10) + "****"

	Generate(newNumber, true)
}

// checking does brand support issuer
func checkIssuer(brand, issuer string) bool {
	if len(brand) <= 0 || len(issuer) <= 0 {
		os.Exit(1)
	}

	return strings.HasPrefix(issuer, brand)
}
