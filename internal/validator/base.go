package validator

// Validator is the base struct for all validators.
type Validator struct {
	stringValidations []func(string) *ValidationResult
	numberValidations []func(float64) *ValidationResult
}

// NewValidator creates a new base validator.
//
// Returns:
//   - A pointer to the newly created Validator instance.
func NewValidator() *Validator {
	return &Validator{}
}

// AddStringValidation adds a string validation function.
func (v *Validator) AddStringValidation(validation func(string) *ValidationResult) {
	v.stringValidations = append(v.stringValidations, validation)
}

// AddNumberValidation adds a number validation function.
func (v *Validator) AddNumberValidation(validation func(float64) *ValidationResult) {
	v.numberValidations = append(v.numberValidations, validation)
}

// ValidateString executes all string validations.
//
// Parameters:
//   - value: The string value to be validated.
//
// Returns:
//   - The validation result.
func (v *Validator) ValidateString(value string) *ValidationResult {
	result := &ValidationResult{IsValid: true}
	for _, validation := range v.stringValidations {
		validationResult := validation(value)
		if !validationResult.IsValid {
			result.IsValid = false
			result.Errors = append(result.Errors, validationResult.Errors...)
		}
	}
	return result
}

// ValidateNumber executes all number validations.
//
// Parameters:
//   - value: The number value to be validated.
//
// Returns:
//   - The validation result.
func (v *Validator) ValidateNumber(value float64) *ValidationResult {
	result := &ValidationResult{IsValid: true}
	for _, validation := range v.numberValidations {
		validationResult := validation(value)
		if !validationResult.IsValid {
			result.IsValid = false
			result.Errors = append(result.Errors, validationResult.Errors...)
		}
	}
	return result
}
