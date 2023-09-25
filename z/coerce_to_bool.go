package z

import (
	"errors"
)

type coerceBoolSchema struct {
	customMessage string
}

// NewCoerceBoolSchema creates a new instance of the coerceBoolSchema for parsing and validating numbers.
//
// Returns:
//   - A pointer to the newly created coerceBoolSchema instance.
func NewCoerceBoolSchema() *coerceBoolSchema {
	return &coerceBoolSchema{}
}

func (c *coerceBoolSchema) coerceToBoolean(data interface{}) (bool, error) {
	switch val := data.(type) {
	case bool:
		return val, nil
	case string:
		if len(val) > 0 {
			return true, nil
		}
		return false, nil
	case int:
		// interpret non-zero integers as true, and zero as false
		if val != 0 {
			return true, nil
		}

		if val == 0 {
			return false, nil
		}
	case []interface{}:
		return true, nil
	case nil:
		// nil is interpreted as false
		return false, nil
	}

	message := "value must be a boolean"
	if c.customMessage != "" {
		message = c.customMessage
	}

	return false, errors.New(message)
}

// Parse coerces and validates the input data as a boolean.
//
// Parameters:
//   - data: The data to be coerced and validated as a boolean.
//
// Returns:
//   - The coerced and validated boolean value.
//   - An error if coercion or validation fails, which is typically nil in this implementation.
func (cs *coerceBoolSchema) Parse(data interface{}) (bool, error) {
	value, err := cs.coerceToBoolean(data)
	if err != nil {
		return false, err
	}
	return value, nil
}

// Message sets a custom error message for the coerceBoolSchema instance.
//
// Parameters:
//   - customMessage: The custom error message to be used when validation fails.
//
// Returns:
//   - The coerceBoolSchema instance with the custom error message set.
func (cs *coerceBoolSchema) Message(customMessage string) *coerceBoolSchema {
	cs.customMessage = customMessage
	return cs
}
