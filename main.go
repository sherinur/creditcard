package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	arguments := os.Args[1:]
	var numbers []string

	// flags
	var (
		stdinFlag, _ = contains(&arguments, "--stdin", true, false)
		pickFlag, _  = contains(&arguments, "--pick", true, false)
		feature      = identifyFeature(&arguments)
	)

	// Paths of files
	brandsFilePath := checkExtraFlag(&arguments, "--brands=")
	issuersFilePath := checkExtraFlag(&arguments, "--issuers=")

	// brand and issuer
	brand := checkExtraFlag(&arguments, "--brand=")
	issuer := checkExtraFlag(&arguments, "--issuer=")

	// reading from the pipe for stdin case
	if stdinFlag {
		// TODO: WRITE ERROR HANDLING FOR --stdin
		// ! Не работает случай когда вводят два номера без пробела между ними
		input, _ := io.ReadAll(os.Stdin)
		inputStr := string(input)
		numbers = strings.Fields(inputStr)
	} else {
		// default case
		numbers = arguments
	}

	if len(numbers) <= 0 && feature != "issue" {
		os.Exit(1)
	}

	// switch for the features
	switch feature {
	case "validate":
		if !stdinFlag {
			numbers = arguments
		}
		isValid := true // TODO: COMPLETE THIS
		for _, number := range numbers {
			RemoveQuotes(number)
			if LuhnTest(number) {
				fmt.Fprintln(os.Stdout, "OK")
			} else {
				fmt.Fprintln(os.Stderr, "INCORRECT")
				isValid = false
			}
		}

		// exiting with code
		if isValid {
			os.Exit(0)
		} else {
			os.Exit(1)
		}

	case "generate":
		// ! IS NOT COMPLETED PART
		if len(numbers) <= 0 {
			printUsage()
			os.Exit(1)
		}
		cardTemplate := numbers[0]
		Generate(cardTemplate, pickFlag)
	case "information":
		// ! НЕ ДОРАБОТАНО
		if len(brandsFilePath) <= 0 || len(issuersFilePath) <= 0 {
			printUsage()
			os.Exit(1)
		}
		Information(numbers, brandsFilePath, issuersFilePath)
		os.Exit(0)
	case "issue":
		// error handling
		if len(brandsFilePath) <= 0 || len(issuersFilePath) <= 0 {
			printUsage()
			os.Exit(1)
		}
		if len(brand) <= 0 {
			fmt.Fprintln(os.Stderr, "Error: Brand naming is incorrect.")
			os.Exit(1)
		} else if len(issuer) <= 0 {
			fmt.Fprintln(os.Stderr, "Error: Issuer naming is incorrect.")
			os.Exit(1)
		}
		Issue(brandsFilePath, issuersFilePath, brand, issuer)
	default:
		printUsage()
		os.Exit(1)
	}

	os.Exit(0)
}

func printUsage() {
	usage := `Usage:
	./creditcard validate <card_number> [<card_number> ...] or ./creditcard validate --stdin
	./creditcard generate <card_number_with_asterisks> [--pick]
	./creditcard information --brands=<brands_file> --issuers=<issuers_file> <card_number> [<card_number> ...]
	./creditcard issue --brands=<brands_file> --issuers=<issuers_file> --brand=<brand> --issuer=<issuer>`
	fmt.Println(usage)
}
