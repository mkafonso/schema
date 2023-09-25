package z

import (
	"errors"
	"fmt"
)

type stringMinSchema struct {
	minLength     int
	customMessage string
}

// NewStringMinSchema creates a new instance of the stringMinSchema for parsing and validating strings.
//
// Returns:
//   - A pointer to the newly created stringMinSchema instance.
func NewStringMinSchema() *stringMinSchema {
	return &stringMinSchema{}
}

// Parse parses and validates the input data as a string, applying the "Min" constraint.
//
// Parameters:
//   - data: The data to be validated as a string.
//
// Returns:
//   - The parsed string value if the validation succeeds.
//   - An error with a custom error message if the validation fails due to length constraints.
func (ss *stringMinSchema) Parse(data interface{}) (string, error) {
	if str, ok := data.(string); ok {
		if ss.minLength > 0 && len(str) < ss.minLength {
			message := fmt.Sprintf("must be %d or more characters long", ss.minLength)
			if ss.customMessage != "" {
				message = ss.customMessage
			}
			return "", errors.New(message)
		}
		return str, nil
	}

	message := "value must be a string"
	if ss.customMessage != "" {
		message = ss.customMessage
	}

	return "", errors.New(message)
}

// Min sets the minimum length validation constraint for the string.
//
// Parameters:
//   - length: The minimum length that the string must have to pass validation.
//
// Returns:
//   - The stringMinSchema instance with the minimum length constraint set.
func (ss *stringMinSchema) Min(length int) *stringMinSchema {
	ss.minLength = length
	return ss
}

// Message sets a custom error message for the stringMinSchema instance.
//
// Parameters:
//   - customMessage: The custom error message to be used when validation fails.
//
// Returns:
//   - The stringMinSchema instance with the custom error message set.
func (ss *stringMinSchema) Message(customMessage string) *stringMinSchema {
	ss.customMessage = customMessage
	return ss
}
