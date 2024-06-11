package number

import (
	"fmt"

	"github.com/mkafonso/schema/pkg/validator"
)

func ValidateNegativeExample() {
	input := -1.0

	// Defining a validation schema
	numberValidator := validator.NewNumberValidator().
		IsNegative("The number must be negative") // Validates that the number is negative

	// Validating the input
	result := numberValidator.Validate(input)

	// Handling validation result
	if result.IsValid {
		fmt.Println("Negative number validation successful!") // Message for successful validation
	} else {
		fmt.Println("Negative number validation error:", result.Errors) // Printing validation errors
	}
}
