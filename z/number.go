package z

import "errors"

type numberSchema struct {
	customMessage string
}

// NewNumberSchema creates a new instance of the numberSchema for parsing and validating numbers.
//
// Returns:
//   - A pointer to the newly created numberSchema instance.
func NewNumberSchema() *numberSchema {
	return &numberSchema{}
}

// Parse parses and validates the input data as a number.
//
// Parameters:
//   - data: The data to be validated as a number.
//
// Returns:
//   - The parsed number value if the validation succeeds, or 0.0.
//   - An error with a custom error message if the validation fails.
func (ns *numberSchema) Parse(data interface{}) (float64, error) {
	if num, ok := data.(float64); ok {
		return num, nil
	}

	message := "value must be a number"
	if ns.customMessage != "" {
		message = ns.customMessage
	}

	return 0, errors.New(message)
}

// Message sets a custom error message for the numberSchema instance.
//
// Parameters:
//   - customMessage: The custom error message to be used when validation fails.
//
// Returns:
//   - The numberSchema instance with the custom error message set.
func (ns *numberSchema) Message(customMessage string) *numberSchema {
	ns.customMessage = customMessage
	return ns
}
