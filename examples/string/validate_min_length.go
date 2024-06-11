package string

import (
	"fmt"

	"github.com/mkafonso/schema/pkg/validator"
)

func ValidateMinLengthExample() {
	input := "short"

	// Defining a validation schema
	minLengthValidator := validator.NewStringValidator().
		MinLength(6, "The string must be at least 6 characters long") // Sets minimum length constraint

	// Validating the input
	result := minLengthValidator.Validate(input)

	// Handling validation result
	if result.IsValid {
		fmt.Println("Minimum length validation successful!") // Message for successful validation
	} else {
		fmt.Println("Minimum length validation error:", result.Errors) // Printing validation errors
	}
}
