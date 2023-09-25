package z

import "errors"

type stringSchema struct {
	customMessage string
}

// NewStringSchema creates a new instance of the stringSchema for parsing and validating strings.
//
// Returns:
//   - A pointer to the newly created stringSchema instance.
func NewStringSchema() *stringSchema {
	return &stringSchema{}
}

// Parse parses and validates the input data as a string.
//
// Parameters:
//   - data: The data to be validated as a string.
//
// Returns:
//   - The parsed string value if the validation succeeds, or an empty string.
//   - An error with a custom error message if the validation fails.
func (ns *stringSchema) Parse(data interface{}) (string, error) {
	if str, ok := data.(string); ok {
		return str, nil
	}

	message := "value must be a string"
	if ns.customMessage != "" {
		message = ns.customMessage
	}

	return "", errors.New(message)
}

// Message sets a custom error message for the stringSchema instance.
//
// Parameters:
//   - customMessage: The custom error message to be used when validation fails.
//
// Returns:
//   - The stringSchema instance with the custom error message set.
func (ns *stringSchema) Message(customMessage string) *stringSchema {
	ns.customMessage = customMessage
	return ns
}
