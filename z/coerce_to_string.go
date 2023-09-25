package z

import "fmt"

type coerceStringSchema struct {
	customMessage string
}

// NewCoerceStringSchema creates a new instance of the coerceStringSchema for parsing and validating numbers.
//
// Returns:
//   - A pointer to the newly created coerceStringSchema instance.
func NewCoerceStringSchema() *coerceStringSchema {
	return &coerceStringSchema{}
}

func (cs *coerceStringSchema) coerceToString(data interface{}) string {
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
func (cs *coerceStringSchema) Parse(data interface{}) (string, error) {
	strData := cs.coerceToString(data)
	return strData, nil
}

// Message sets a custom error message for the coerceStringSchema instance.
//
// Parameters:
//   - customMessage: The custom error message to be used when validation fails.
//
// Returns:
//   - The coerceStringSchema instance with the custom error message set.
func (cs *coerceStringSchema) Message(customMessage string) *coerceStringSchema {
	cs.customMessage = customMessage
	return cs
}
