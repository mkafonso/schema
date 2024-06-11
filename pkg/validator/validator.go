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

// ValidationResult represents the result of a validation.
type ValidationResult = validator.ValidationResult
