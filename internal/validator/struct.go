package validator

import (
	"reflect"
)

// StructValidator is a validator for structs.
type StructValidator struct {
	fields map[string]ValidatorInterface
}

// NewStructValidator creates a new struct validator.
//
// Returns:
//   - A pointer to the newly created StructValidator instance.
func NewStructValidator() *StructValidator {
	return &StructValidator{fields: make(map[string]ValidatorInterface)}
}

// Field adds a field validator to the struct validator.
//
// Parameters:
//   - fieldName: The name of the field to validate.
//   - fieldValidator: The validator to use for the field.
//
// Returns:
//   - A pointer to the StructValidator instance.
func (v *StructValidator) Field(fieldName string, fieldValidator ValidatorInterface) *StructValidator {
	v.fields[fieldName] = fieldValidator
	return v
}

// Validate executes all validations on the struct.
//
// Parameters:
//   - s: The struct to be validated.
//
// Returns:
//   - The validation result.
func (v *StructValidator) Validate(s interface{}) *ValidationResult {
	result := &ValidationResult{IsValid: true}

	val := reflect.ValueOf(s)
	if val.Kind() != reflect.Struct {
		result.IsValid = false
		result.Errors = append(result.Errors, "Input is not a struct")
		return result
	}

	for fieldName, validator := range v.fields {
		field := val.FieldByName(fieldName)
		if !field.IsValid() {
			result.IsValid = false
			result.Errors = append(result.Errors, "Field "+fieldName+" does not exist in struct")
			continue
		}

		fieldResult := validator.ValidateInterface(field.Interface())
		if !fieldResult.IsValid {
			result.IsValid = false
			result.Errors = append(result.Errors, fieldResult.Errors...)
		}
	}

	return result
}
