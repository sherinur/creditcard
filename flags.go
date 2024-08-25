package main

import "os"

// functions and methods to implement flags

// checks additional flags in format "--flag=", and returns boolean,
// removes flag from the argument list and gives the value of the flag
func checkExtraFlag(arguments *[]string, extraFlag string) string {
	// checking does arguments contain this flag
	isThisFlag, index := contains(arguments, extraFlag, false, true)

	if isThisFlag && index >= 0 {
		// storing value of the flag
		valueWithFlag := (*arguments)[index]
		value := valueWithFlag[len(extraFlag):]

		// delete after finding
		*arguments = append((*arguments)[:index], (*arguments)[index+1:]...)
		return value
	}

	return ""
}

// checks via contains every item of the arguments string for feature
func identifyFeature(arguments *[]string) string {
	features := []string{"validate", "generate", "information", "issue"}

	for _, feature := range features {
		isThisFeature, _ := contains(arguments, feature, true, false)
		if isThisFeature {
			return feature
		}
	}
	return ""
}

// checks if array contains the item, or item that starts alike,
// removes it and returns boolean and index of item.
// has a indicator delete to a delete item if array contains that item
func contains(arguments *[]string, item string, delete bool, isFilePath bool) (bool, int) {
	for index, argument := range *arguments {
		if len(item) > len(argument) {
			continue
		}

		match := true
		for i := 0; i < len(item); i++ {
			if item[i] != argument[i] {
				match = false
				break
			}
		}

		if match {
			if !isFilePath && len(argument) > len(item) {
				printUsage()
				os.Exit(1)
			}

			if delete {
				*arguments = append((*arguments)[:index], (*arguments)[index+1:]...)
			}
			return true, index
		}
	}
	return false, -1
}
