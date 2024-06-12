package validate_string

import (
	"fmt"

	"github.com/mkafonso/schema/pkg/validator"
)

func ValidateStringExample() {
	input := "Luna"

	// Defining a validation schema
	stringValidator := validator.NewStringValidator().
		MinLength(5, "The string must be at least 5 characters long"). // Sets minimum length constraint
		MaxLength(50, "The string must be at most 50 characters long") // Sets maximum length constraint

	// Validating the input
	result := stringValidator.Validate(input)

	// Handling validation result
	if result.IsValid {
		fmt.Println("String validation successful!") // Message for successful validation
	} else {
		fmt.Println("String validation error:", result.Errors) // Printing validation errors
	}
}
