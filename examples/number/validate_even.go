package number

import (
	"fmt"

	"github.com/mkafonso/schema/pkg/validator"
)

func ValidateEvenExample() {
	input := 28.0

	// Defining a validation schema
	numberValidator := validator.NewNumberValidator().
		IsEven("The number must be even") // Validates that the number is even

	// Validating the input
	result := numberValidator.Validate(input)

	// Handling validation result
	if result.IsValid {
		fmt.Println("Even number validation successful!") // Message for successful validation
	} else {
		fmt.Println("Even number validation error:", result.Errors) // Printing validation errors
	}
}
