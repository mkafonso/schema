package validator

// BooleanValidator is a validator for boolean values.
type BooleanValidator struct {
	validations []func(bool) *ValidationResult
}

// NewBooleanValidator creates a new boolean validator.
//
// Returns:
//   - A pointer to the newly created BooleanValidator instance.
func NewBooleanValidator() *BooleanValidator {
	return &BooleanValidator{}
}

// addValidation adds a custom validation function.
func (v *BooleanValidator) addValidation(validation func(bool) *ValidationResult) {
	v.validations = append(v.validations, validation)
}

// IsTrue adds a validation to check if the boolean is true.
//
// Parameters:
//   - errorMessage: The error message to be returned if validation fails.
//
// Returns:
//   - A pointer to the BooleanValidator instance.
func (v *BooleanValidator) IsTrue(errorMessage string) *BooleanValidator {
	v.addValidation(func(value bool) *ValidationResult {
		if !value {
			return &ValidationResult{IsValid: false, Errors: []string{errorMessage}}
		}
		return &ValidationResult{IsValid: true}
	})
	return v
}

// IsFalse adds a validation to check if the boolean is false.
//
// Parameters:
//   - errorMessage: The error message to be returned if validation fails.
//
// Returns:
//   - A pointer to the BooleanValidator instance.
func (v *BooleanValidator) IsFalse(errorMessage string) *BooleanValidator {
	v.addValidation(func(value bool) *ValidationResult {
		if value {
			return &ValidationResult{IsValid: false, Errors: []string{errorMessage}}
		}
		return &ValidationResult{IsValid: true}
	})
	return v
}

// Validate executes all validations.
//
// Parameters:
//   - value: The boolean value to be validated.
//
// Returns:
//   - The validation result.
func (v *BooleanValidator) Validate(value bool) *ValidationResult {
	result := &ValidationResult{IsValid: true}
	for _, validation := range v.validations {
		validationResult := validation(value)
		if !validationResult.IsValid {
			result.IsValid = false
			result.Errors = append(result.Errors, validationResult.Errors...)
		}
	}
	return result
}

// ValidateInterface validates the input value, implementing ValidatorInterface.
//
// Parameters:
//   - value: The value to be validated.
//
// Returns:
//   - The validation result.
func (v *BooleanValidator) ValidateInterface(value interface{}) *ValidationResult {
	if boolVal, ok := value.(bool); ok {
		return v.Validate(boolVal)
	}
	return &ValidationResult{IsValid: false, Errors: []string{"Invalid type, expected boolean"}}
}
