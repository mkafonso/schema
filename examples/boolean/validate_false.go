package main

import (
	"fmt"

	"github.com/mkafonso/schema/pkg/validator"
)

func ValidateFalseExample() {
	boolValidator := validator.NewBooleanValidator().
		IsFalse("Value must be false")

	result := boolValidator.Validate(false)
	fmt.Println("Example - Validating false:")
	fmt.Printf("Result: %v\n", result.IsValid)
	if !result.IsValid {
		fmt.Printf("Errors: %v\n", result.Errors)
	}
}
