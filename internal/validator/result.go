package validator

// ValidationResult represents the result of a validation.
type ValidationResult struct {
	IsValid bool     // Indicates if the validation was successful or not
	Errors  []string // Contains error messages if validation fails
}

// ValidatorInterface is the interface for all validators.
type ValidatorInterface interface {
	ValidateInterface(value interface{}) *ValidationResult
}
