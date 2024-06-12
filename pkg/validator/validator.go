package validator

import (
	"github.com/mkafonso/schema/internal/validator"
)

// NewStringValidator creates a new instance of the StringValidator.
//
// Returns:
//   - A pointer to the newly created StringValidator instance.
func NewStringValidator() *validator.StringValidator {
	return validator.NewStringValidator()
}

// NewNumberValidator creates a new instance of the NumberValidator.
//
// Returns:
//   - A pointer to the newly created NumberValidator instance.
func NewNumberValidator() *validator.NumberValidator {
	return validator.NewNumberValidator()
}

// NewStructValidator creates a new instance of the StructValidator.
//
// Returns:
//   - A pointer to the newly created StructValidator instance.
func NewStructValidator() *validator.StructValidator {
	return validator.NewStructValidator()
}

// ValidationResult represents the result of a validation.
type ValidationResult = validator.ValidationResult
