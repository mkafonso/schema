package validate_string

import (
	"fmt"

	"github.com/mkafonso/schema/pkg/validator"
)

func ValidateMaxLengthExample() {
	input := "123456789123456789123456789"

	// Defining a validation schema
	maxLengthValidator := validator.NewStringValidator().
		MaxLength(10, "The string must be at most 10 characters long") // Sets maximum length constraint

	// Validating the input
	result := maxLengthValidator.Validate(input)

	// Handling validation result
	if result.IsValid {
		fmt.Println("Maximum length validation successful!") // Message for successful validation
	} else {
		fmt.Println("Maximum length validation error:", result.Errors) // Printing validation errors
	}
}
