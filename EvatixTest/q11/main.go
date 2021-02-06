package main

import (
	"strconv"
)

func currencyConvert(input string, currencySymbol string, decimalPlaces int) string {
	isEmpty(input)
	if !isInteger(input) || isDecimal(input) {
		return input
	}
	currencySymbol, decimalPlaces = checkAgainstDefault(currencySymbol, decimalPlaces)
	integerPartOfInput := getIntegerPart(input)
	decimalPartOfInput := getDecimalPart(input)
	return convertToCurrency(integerPartOfInput, decimalPartOfInput, currencySymbol)
}
func isEmpty(input string) bool {
	return input == "" || len(input) == 0
}
func isInteger(input string) bool {
	_, err := strconv.Atoi(input)
	if err != nil {
		return false
	} else {
		return true
	}
}
func isDecimal(input string) bool {
	if _, err := strconv.ParseFloat(input, 10); err != nil {
		return false
	} else {
		return true
	}
}
func main() {

}
