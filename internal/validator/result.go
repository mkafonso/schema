package validator

// ValidationResult represents the result of a validation.
type ValidationResult struct {
	IsValid bool     // Indicates if the validation was successful or not
	Errors  []string // Contains error messages if validation fails
}
