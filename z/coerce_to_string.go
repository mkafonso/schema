package z

import "fmt"

type coerceSchema struct {
	customMessage string
}

// NewCoerceSchema creates a new instance of the coerceSchema for parsing and validating numbers.
//
// Returns:
//   - A pointer to the newly created coerceSchema instance.
func NewCoerceSchema() *coerceSchema {
	return &coerceSchema{}
}

func (cs *coerceSchema) coerceToString(data interface{}) string {
	return fmt.Sprintf("%v", data)
}

// Parse coerces and validates the input data as a string.
//
// Parameters:
//   - data: The data to be coerced and validated as a string.
//
// Returns:
//   - The coerced and validated string value.
//   - An error if coercion or validation fails, which is typically nil in this implementation.
func (cs *coerceSchema) Parse(data interface{}) (string, error) {
	strData := cs.coerceToString(data)
	return strData, nil
}

// Message sets a custom error message for the coerceSchema instance.
//
// Parameters:
//   - customMessage: The custom error message to be used when validation fails.
//
// Returns:
//   - The coerceSchema instance with the custom error message set.
func (cs *coerceSchema) Message(customMessage string) *coerceSchema {
	cs.customMessage = customMessage
	return cs
}
