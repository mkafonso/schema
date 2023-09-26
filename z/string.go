package z

import "errors"

type stringSchema struct {
	minLength     int
	maxLength     int
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
func (ss *stringSchema) Parse(data interface{}) (string, error) {
	if str, ok := data.(string); ok {
		if ss.minLength > 0 && len(str) < ss.minLength {
			return "", errors.New(ss.customMessage)
		}
		if ss.maxLength > 0 && len(str) > ss.maxLength {
			return "", errors.New(ss.customMessage)
		}
		return str, nil
	}

	message := "value must be a string"
	if ss.customMessage != "" {
		message = ss.customMessage
	}

	return "", errors.New(message)
}

// Min sets the minimum length for the string.
//
// Parameters:
//   - minLength: The minimum length that the string must have.
//
// Returns:
//   - The stringSchema instance with the minimum length set.
func (ss *stringSchema) Min(minLength int) *stringSchema {
	ss.minLength = minLength
	return ss
}

// Max sets the maximum length for the string.
//
// Parameters:
//   - maxLength: The maximum length that the string must not exceed.
//
// Returns:
//   - The stringSchema instance with the maximum length set.
func (ss *stringSchema) Max(maxLength int) *stringSchema {
	ss.maxLength = maxLength
	return ss
}

// Message sets a custom error message for the stringSchema instance.
//
// Parameters:
//   - customMessage: The custom error message to be used when validation fails.
//
// Returns:
//   - The stringSchema instance with the custom error message set.
func (ss *stringSchema) Message(customMessage string) *stringSchema {
	ss.customMessage = customMessage
	return ss
}
