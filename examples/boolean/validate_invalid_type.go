package main

import (
	"fmt"

	"github.com/mkafonso/schema/pkg/validator"
)

func ValidateInvalidTypeExample() {
	boolValidator := validator.NewBooleanValidator()

	result := boolValidator.ValidateInterface("not a boolean")
	fmt.Println("Example - Validating invalid type:")
	fmt.Printf("Result: %v\n", result.IsValid)
	if !result.IsValid {
		fmt.Printf("Errors: %v\n", result.Errors)
	}
}
