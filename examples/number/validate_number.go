package number

import (
	"fmt"

	"github.com/mkafonso/schema/pkg/validator"
)

func ValidateNumberExample() {
	input := 28.0

	// Defining a validation schema
	numberValidator := validator.NewNumberValidator().
		MinValue(10, "Age must be at least 18").  // Sets minimum value constraint
		MaxValue(100, "Age must be at most 100"). // Sets maximum value constraint
		IsInteger("Age must be an integer")       // Validates as an integer

	// Validating the input
	result := numberValidator.Validate(input)

	// Handling validation result
	if result.IsValid {
		fmt.Println("Number validation successful!") // Message for successful validation
	} else {
		fmt.Println("Number validation error:", result.Errors) // Printing validation errors
	}
}
