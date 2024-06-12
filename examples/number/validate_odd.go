package validate_number

import (
	"fmt"

	"github.com/mkafonso/schema/pkg/validator"
)

func ValidateOddExample() {
	input := 21.0

	// Defining a validation schema
	numberValidator := validator.NewNumberValidator().
		IsOdd("The number must be odd") // Validates that the number is odd

	// Validating the input
	result := numberValidator.Validate(input)

	// Handling validation result
	if result.IsValid {
		fmt.Println("Odd number validation successful!") // Message for successful validation
	} else {
		fmt.Println("Odd number validation error:", result.Errors) // Printing validation errors
	}
}
