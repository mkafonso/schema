package main

import (
	"fmt"

	validate_string "github.com/mkafonso/examples/string"
)

func main() {
	fmt.Println("\nString validation")
	fmt.Println("\n---------------------------------------")

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
}
