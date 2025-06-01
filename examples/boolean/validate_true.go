package main

import (
	"fmt"

	"github.com/mkafonso/schema/pkg/validator"
)

func ValidateTrueExample() {
	boolValidator := validator.NewBooleanValidator().
		IsTrue("Value must be true")

	result := boolValidator.Validate(true)
	fmt.Println("Example - Validating true:")
	fmt.Printf("Result: %v\n", result.IsValid)
	if !result.IsValid {
		fmt.Printf("Errors: %v\n", result.Errors)
	}
}
