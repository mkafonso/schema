package validate_struct

import (
	"fmt"

	"github.com/mkafonso/schema/pkg/validator"
)

type User struct {
	Name  string
	Age   int
	Email string
}

func ValidateStructExample() {
	user := User{
		Name:  "Dolores Madrigal",
		Age:   21,
		Email: "dolores.madrigal@disney.com",
	}

	// Defining a validation schema
	structValidator := validator.NewStructValidator().
		Field("Name", validator.NewStringValidator().MinLength(2, "Name must be at least 2 characters long")).
		Field("Age", validator.NewNumberValidator().MinValue(18, "Age must be at least 18")).
		Field("Email", validator.NewStringValidator().IsEmail("Invalid email format"))

	// Validating the input
	result := structValidator.Validate(user)

	// Handling validation result
	if result.IsValid {
		fmt.Println("Validation successful!") // Message for successful validation
	} else {
		fmt.Println("Validation error:", result.Errors) // Printing validation errors
	}
}
