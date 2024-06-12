package validator

// NumberValidator is a validator for numbers.
type NumberValidator struct {
	validations []func(float64) *ValidationResult
}

// NewNumberValidator creates a new number validator.
//
// Returns:
//   - A pointer to the newly created NumberValidator instance.
func NewNumberValidator() *NumberValidator {
	return &NumberValidator{}
}

// addValidation adds a custom validation function.
func (v *NumberValidator) addValidation(validation func(float64) *ValidationResult) {
	v.validations = append(v.validations, validation)
}

// MinValue adds a validation for minimum value.
//
// Parameters:
//   - min: The minimum value allowed.
//   - errorMessage: The error message to be returned if validation fails.
//
// Returns:
//   - A pointer to the NumberValidator instance.
func (v *NumberValidator) MinValue(min float64, errorMessage string) *NumberValidator {
	v.addValidation(func(value float64) *ValidationResult {
		if value < min {
			return &ValidationResult{IsValid: false, Errors: []string{errorMessage}}
		}
		return &ValidationResult{IsValid: true}
	})
	return v
}

// MaxValue adds a validation for maximum value.
//
// Parameters:
//   - max: The maximum value allowed.
//   - errorMessage: The error message to be returned if validation fails.
//
// Returns:
//   - A pointer to the NumberValidator instance.
func (v *NumberValidator) MaxValue(max float64, errorMessage string) *NumberValidator {
	v.addValidation(func(value float64) *ValidationResult {
		if value > max {
			return &ValidationResult{IsValid: false, Errors: []string{errorMessage}}
		}
		return &ValidationResult{IsValid: true}
	})
	return v
}

// IsInteger adds a validation to check if the number is an integer.
//
// Parameters:
//   - errorMessage: The error message to be returned if validation fails.
//
// Returns:
//   - A pointer to the NumberValidator instance.
func (v *NumberValidator) IsInteger(errorMessage string) *NumberValidator {
	v.addValidation(func(value float64) *ValidationResult {
		if value != float64(int(value)) {
			return &ValidationResult{IsValid: false, Errors: []string{errorMessage}}
		}
		return &ValidationResult{IsValid: true}
	})
	return v
}

// IsPositive adds a validation to check if the number is positive.
//
// Parameters:
//   - errorMessage: The error message to be returned if validation fails.
//
// Returns:
//   - A pointer to the NumberValidator instance.
func (v *NumberValidator) IsPositive(errorMessage string) *NumberValidator {
	v.addValidation(func(value float64) *ValidationResult {
		if value <= 0 {
			return &ValidationResult{IsValid: false, Errors: []string{errorMessage}}
		}
		return &ValidationResult{IsValid: true}
	})
	return v
}

// IsNegative adds a validation to check if the number is negative.
//
// Parameters:
//   - errorMessage: The error message to be returned if validation fails.
//
// Returns:
//   - A pointer to the NumberValidator instance.
func (v *NumberValidator) IsNegative(errorMessage string) *NumberValidator {
	v.addValidation(func(value float64) *ValidationResult {
		if value >= 0 {
			return &ValidationResult{IsValid: false, Errors: []string{errorMessage}}
		}
		return &ValidationResult{IsValid: true}
	})
	return v
}

// IsEven adds a validation to check if the number is even.
//
// Parameters:
//   - errorMessage: The error message to be returned if validation fails.
//
// Returns:
//   - A pointer to the NumberValidator instance.
func (v *NumberValidator) IsEven(errorMessage string) *NumberValidator {
	v.addValidation(func(value float64) *ValidationResult {
		if int(value)%2 != 0 {
			return &ValidationResult{IsValid: false, Errors: []string{errorMessage}}
		}
		return &ValidationResult{IsValid: true}
	})
	return v
}

// IsOdd adds a validation to check if the number is odd.
//
// Parameters:
//   - errorMessage: The error message to be returned if validation fails.
//
// Returns:
//   - A pointer to the NumberValidator instance.
func (v *NumberValidator) IsOdd(errorMessage string) *NumberValidator {
	v.addValidation(func(value float64) *ValidationResult {
		if int(value)%2 == 0 {
			return &ValidationResult{IsValid: false, Errors: []string{errorMessage}}
		}
		return &ValidationResult{IsValid: true}
	})
	return v
}

// Validate executes all validations.
//
// Parameters:
//   - value: The number value to be validated.
//
// Returns:
//   - The validation result.
func (v *NumberValidator) Validate(value float64) *ValidationResult {
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
func (v *NumberValidator) ValidateInterface(value interface{}) *ValidationResult {
	switch val := value.(type) {
	case float64:
		return v.Validate(val)
	case int:
		return v.Validate(float64(val))
	default:
		return &ValidationResult{IsValid: false, Errors: []string{"Invalid type, expected number"}}
	}
}
