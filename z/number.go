package z

import (
	"errors"
	"strconv"
)

type NumberSchema struct {
	customMessage string
}

// NewNumberSchema creates a new instance of the NumberSchema for parsing and validating numbers.
//
// Returns:
//   - A pointer to the newly created NumberSchema instance.
func NewNumberSchema() *NumberSchema {
	return &NumberSchema{}
}

// Parse parses and validates the input data as a number.
//
// Parameters:
//   - data: The data to be validated as a number.
//   - errorMessage (optional): The custom error message to be used when validation fails.
//
// Returns:
//   - The parsed number value if the validation succeeds.
//   - An error with the custom error message if the validation fails.
func (ns *NumberSchema) Parse(data interface{}, errorMessage ...string) (float64, error) {
	var num float64

	switch val := data.(type) {
	case float64:
		num = val
	case int:
		num = float64(val)
	case string:
		parsedNum, err := strconv.ParseFloat(val, 64)
		if err != nil {
			message := "value must be a number"
			if ns.customMessage != "" {
				message = ns.customMessage
			}

			if len(errorMessage) > 0 {
				message = errorMessage[0]
			}

			return 0, errors.New(message)
		}
		num = parsedNum
	default:
		message := "value must be a number"
		if ns.customMessage != "" {
			message = ns.customMessage
		}
		if len(errorMessage) > 0 {
			message = errorMessage[0]
		}
		return 0, errors.New(message)
	}

	return num, nil
}
