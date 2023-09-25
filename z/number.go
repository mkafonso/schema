package z

import "errors"

// IsNumber checks if the given data is a valid number.
//
// Parameters:
//   - data: The data to be validated.
//   - customMessage: An optional custom error message to use when the validation fails.
//
// Returns:
//   - An error if the data is not a valid number; otherwise, nil.
func IsNumber(data interface{}, customMessage ...string) error {
	if _, ok := data.(float64); !ok {
		defaultMessage := "value must be a number"
		if len(customMessage) > 0 {
			defaultMessage = customMessage[0]
		}
		return errors.New(defaultMessage)
	}
	return nil
}
