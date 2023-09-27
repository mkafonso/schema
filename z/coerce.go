package z

import (
	"errors"
	"fmt"
)

type coerceSchema struct {
	coerceToStrError  string
	coerceToBoolError string
	customMessage     string
	coerceToString    bool
	coerceToBool      bool
}

// NewCoerceSchema creates a new instance of the coerceSchema for parsing and coercing values to string or bool.
//
// Returns:
//   - A pointer to the newly created coerceSchema instance.
func NewCoerceSchema() *coerceSchema {
	return &coerceSchema{}
}

// Parse coerces the input data to either a string or bool, based on the methods called before it.
//
// Parameters:
//   - data: The data to be coerced.
//
// Returns:
//   - The coerced value as either a string or bool.
//   - An error with the custom error message if coercion fails.
func (cs *coerceSchema) Parse(data interface{}) (interface{}, error) {
	if cs.coerceToString {
		return cs.coerceToStringHelper(data)
	}
	if cs.coerceToBool {
		return cs.coerceToBoolHelper(data)
	}

	message := "coercion method not specified"
	if cs.customMessage != "" {
		message = cs.customMessage
	}

	return "", errors.New(message)
}

// String configures the coercion to string and sets a custom error message for coercion failure.
//
// Parameters:
//   - errorMessage (optional): The custom error message to be used when coercion to string fails.
//
// Returns:
//   - The coerceSchema instance with coercion to string configured.
func (cs *coerceSchema) String(errorMessage ...string) *coerceSchema {
	cs.coerceToString = true
	if len(errorMessage) > 0 {
		cs.coerceToStrError = errorMessage[0]
	}
	return cs
}

// helper function for String
func (cs *coerceSchema) coerceToStringHelper(data interface{}) (string, error) {
	return fmt.Sprintf("%v", data), nil
}

// Bool configures the coercion to bool and sets a custom error message for coercion failure.
//
// Parameters:
//   - errorMessage (optional): The custom error message to be used when coercion to bool fails.
//
// Returns:
//   - The coerceSchema instance with coercion to bool configured.
func (cs *coerceSchema) Bool(errorMessage ...string) *coerceSchema {
	cs.coerceToBool = true
	if len(errorMessage) > 0 {
		cs.coerceToBoolError = errorMessage[0]
	}
	return cs
}

// helper function for Bool
func (cs *coerceSchema) coerceToBoolHelper(data interface{}) (bool, error) {
	switch val := data.(type) {
	case bool:
		return val, nil
	case string:
		if len(val) > 0 {
			return true, nil
		}
		return false, nil
	case int:
		if val != 0 {
			return true, nil
		}
		return false, nil
	case []interface{}:
		return true, nil
	case nil:
		return false, nil
	}

	message := "value must be a boolean"
	if cs.customMessage != "" {
		message = cs.customMessage
	}

	return false, errors.New(message)
}
