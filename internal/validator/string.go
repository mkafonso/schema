package validator

import (
	"github.com/mkafonso/schema/internal/utils"
)

// StringValidator is a validator for strings.
type StringValidator struct {
	validator *Validator
}

// NewStringValidator creates a new string validator.
//
// Returns:
//   - A pointer to the newly created StringValidator instance.
func NewStringValidator() *StringValidator {
	v := NewValidator()
	return &StringValidator{validator: v}
}

// MinLength adds a validation for minimum length.
//
// Parameters:
//   - length: The minimum length allowed for the string.
//   - errorMessage: The error message to be returned if validation fails.
//
// Returns:
//   - A pointer to the StringValidator instance.
func (v *StringValidator) MinLength(length int, errorMessage string) *StringValidator {
	v.validator.AddStringValidation(func(value string) *ValidationResult {
		if len(value) < length {
			return &ValidationResult{IsValid: false, Errors: []string{errorMessage}}
		}
		return &ValidationResult{IsValid: true}
	})
	return v
}

// MaxLength adds a validation for maximum length.
//
// Parameters:
//   - length: The maximum length allowed for the string.
//   - errorMessage: The error message to be returned if validation fails.
//
// Returns:
//   - A pointer to the StringValidator instance.
func (v *StringValidator) MaxLength(length int, errorMessage string) *StringValidator {
	v.validator.AddStringValidation(func(value string) *ValidationResult {
		if len(value) > length {
			return &ValidationResult{IsValid: false, Errors: []string{errorMessage}}
		}
		return &ValidationResult{IsValid: true}
	})
	return v
}

// IsEmail adds a validation for email format.
//
// Parameters:
//   - errorMessage: The error message to be returned if validation fails.
//
// Returns:
//   - A pointer to the StringValidator instance.
func (v *StringValidator) IsEmail(errorMessage string) *StringValidator {
	v.validator.AddStringValidation(func(value string) *ValidationResult {
		if !utils.IsValidEmail(value) {
			return &ValidationResult{IsValid: false, Errors: []string{errorMessage}}
		}
		return &ValidationResult{IsValid: true}
	})
	return v
}

// Validate executes all validations.
//
// Parameters:
//   - value: The string value to be validated.
//
// Returns:
//   - The validation result.
func (v *StringValidator) Validate(value string) *ValidationResult {
	return v.validator.ValidateString(value)
}

// ValidateInterface validates the input value, implementing ValidatorInterface.
//
// Parameters:
//   - value: The value to be validated.
//
// Returns:
//   - The validation result.
func (v *StringValidator) ValidateInterface(value interface{}) *ValidationResult {
	if str, ok := value.(string); ok {
		return v.Validate(str)
	}
	return &ValidationResult{IsValid: false, Errors: []string{"Invalid type, expected string"}}
}
