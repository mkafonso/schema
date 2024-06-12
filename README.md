# Schema - A Go Validation Library

This is an open-source schema validation toolkit inspired by zod.dev. It provides a simple and expressive API for validating strings, numbers, struct, ... with various constraints. Schema is designed to be easy to use, extensible, and suitable for a wide range of applications.

Feel free to use and modify it for any purpose. If you have any suggestions for improvement, please don't hesitate to reach out. Thank you!

---

Features include:

-   String Validation: Min/Max length, email format, and more.
-   Number Validation: Min/Max value, integer check, positive/negative, even/odd, divisibility, and more.
-   Custom Error Messages: Provide custom error messages for failed validations.
-   Composability: Combine multiple validations easily.
-   Extensible: Easy to add custom validation logic.

---

## Installation

To install this package, use `go get`:

    go get github.com/mkafonso/schema

---

## Get started:

-   Install schema with [one line of code](#installation)
-   Check out the API Documentation https://pkg.go.dev/github.com/mkafonso/schema

## [`String Validation`](https://pkg.go.dev/github.com/mkafonso/schema/string "API documentation") package

-   Minimum Length: Ensure the string has a minimum length.
-   Maximum Length: Ensure the string does not exceed a maximum length.
-   Email Format: Validate the string as an email format.

See it in action:

```go
package yours

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
```

---

# License

This project is licensed under the terms of the MIT license.
