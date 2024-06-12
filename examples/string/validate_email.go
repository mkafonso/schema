package validate_string

import (
	"fmt"

	"github.com/mkafonso/schema/pkg/validator"
)

func ValidateEmailExample() {
	input := "mkafonso.dev@gmail.com"

	// Defining a validation schema
	stringValidator := validator.NewStringValidator().
		MinLength(5, "A string deve ter no mínimo 5 caracteres").   // Sets minimum length constraint
		MaxLength(50, "A string deve ter no máximo 50 caracteres"). // Sets maximum length constraint
		IsEmail("Formato de email inválido")                        // Validates as an email format

	// Validating the input
	result := stringValidator.Validate(input)

	// Handling validation result
	if result.IsValid {
		fmt.Println("Validation successful!") // Message for successful validation
	} else {
		fmt.Println("Validation error:", result.Errors) // Printing validation errors
	}
}
