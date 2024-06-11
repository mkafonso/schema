package number

import (
	"fmt"

	"github.com/mkafonso/schema/pkg/validator"
)

func ValidatePositiveExample() {
	input := 28.0

	// Defining a validation schema
	numberValidator := validator.NewNumberValidator().
		IsPositive("The number must be positive") // Validates that the number is positive

	// Validating the input
	result := numberValidator.Validate(input)

	// Handling validation result
	if result.IsValid {
		fmt.Println("Positive number validation successful!") // Message for successful validation
	} else {
		fmt.Println("Positive number validation error:", result.Errors) // Printing validation errors
	}
}
