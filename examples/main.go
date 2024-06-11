package main

import (
	"fmt"

	validate_number "github.com/mkafonso/examples/number"
	validate_string "github.com/mkafonso/examples/string"
)

func main() {
	fmt.Println("\n---------------------------------------")
	fmt.Println("\n---------------------------------------")
	fmt.Println("\n---------------------------------------")
	fmt.Print("\nString validation\n\n")

	// Example of string validation
	fmt.Println("Example of string validation:")
	validate_string.ValidateStringExample()

	// Example of email validation
	fmt.Println("\nExample of email validation:")
	validate_string.ValidateEmailExample()

	// Example of minimum length validation
	fmt.Println("\nExample of minimum length validation:")
	validate_string.ValidateMinLengthExample()

	// Example of maximum length validation
	fmt.Println("\nExample of maximum length validation:")
	validate_string.ValidateMaxLengthExample()

	fmt.Println("\n---------------------------------------")
	fmt.Println("\n---------------------------------------")
	fmt.Println("\n---------------------------------------")
	fmt.Print("\nNumber validation\n\n")

	// Example of number validation
	fmt.Println("\nExample of number validation:")
	validate_number.ValidateNumberExample()

	// Example of even number validation
	fmt.Println("\nExample of even number validation:")
	validate_number.ValidateEvenExample()

	// Example of odd number validation
	fmt.Println("\nExample of odd number validation:")
	validate_number.ValidateOddExample()

	// Example of positive number validation
	fmt.Println("\nExample of positive number validation:")
	validate_number.ValidatePositiveExample()

	// Example of negative number validation
	fmt.Println("\nExample of negative number validation:")
	validate_number.ValidateNegativeExample()

	fmt.Println("\n---------------------------------------")
}
