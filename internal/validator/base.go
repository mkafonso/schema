package validator

// Validator defines the base structure of the validator.
type Validator struct {
	validations []func(string) *ValidationResult // List of validation functions
}

// NewValidator creates a new base validator.
//
// Returns:
//   - A pointer to the newly created Validator instance.
func NewValidator() *Validator {
	return &Validator{
		validations: []func(string) *ValidationResult{},
	}
}

// AddValidation adds a new validation to the validator.
//
// Parameters:
//   - validation: The validation function to be added.
func (v *Validator) AddValidation(validation func(string) *ValidationResult) {
	v.validations = append(v.validations, validation)
}

// Validate executes all validations.
//
// Parameters:
//   - value: The value to be validated.
//
// Returns:
//   - The validation result.
func (v *Validator) Validate(value string) *ValidationResult {
	for _, validation := range v.validations {
		result := validation(value)
		if !result.IsValid {
			return result
		}
	}
	return &ValidationResult{IsValid: true}
}
