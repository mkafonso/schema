package z

import "errors"

// IsNumber checks if the given data is a valid number.
//
// Parameters:
//   - data: The data to be validated.
//
// Returns:
//   - An error if the data is not a valid number; otherwise, nil.
func IsNumber(data interface{}) error {
	// Check if the data can be typecasted to a float64.
	// If not, return an error indicating that the value must be a number.
	if _, ok := data.(float64); !ok {
		return errors.New("value must be a number")
	}

	// If the data is a valid number, return nil to indicate no errors.
	return nil
}
