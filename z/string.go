package z

import (
	"errors"
	"regexp"
	"strconv"
)

type stringSchema struct {
	minLength     int
	minError      string
	maxLength     int
	maxError      string
	emailError    string
	emailRegex    *regexp.Regexp
	customMessage string
}

// NewStringSchema creates a new instance of the stringSchema for parsing and validating strings.
//
// Returns:
//   - A pointer to the newly created stringSchema instance.
func NewStringSchema() *stringSchema {
	return &stringSchema{}
}

// Parse parses and validates the input data as a string. If the input data is a valid email address,
// it returns the email string; otherwise, it returns an error with the custom error message.
//
// Parameters:
//   - data: The data to be validated as a string
//   - errorMessage (optional): The custom error message to be used when validation fails.
//
// Returns:
//   - The parsed string value if the validation succeeds.
//   - An error with the custom error message if the validation fails.
func (ss *stringSchema) Parse(data interface{}, errorMessage ...string) (string, error) {
	var str string
	emailConstraint := false

	switch val := data.(type) {
	case string:
		str = val
		// check if Email constraint was set
		if ss.emailRegex != nil {
			emailConstraint = true
		}
	default:
		message := "value must be a string"
		if ss.customMessage != "" {
			message = ss.customMessage
		}

		if len(errorMessage) > 0 {
			message = errorMessage[0]
		}

		return "", errors.New(message)
	}

	if err := ss.validateMinLength(str, errorMessage...); err != nil {
		return "", err
	}

	if err := ss.validateMaxLength(str, errorMessage...); err != nil {
		return "", err
	}

	if emailConstraint {
		if err := ss.validateEmail(str, errorMessage...); err != nil {
			return "", err
		}
	}

	return str, nil
}

// Min sets the minimum length constraint for the string.
//
// Parameters:
//   - length: The minimum required length for the string.
//   - errorMessage (optional): A custom error message to be used when validation fails.
//
// Returns:
//   - The stringSchema instance with the minimum length constraint set.
func (ss *stringSchema) Min(length int, errorMessage ...string) *stringSchema {
	ss.minLength = length
	if len(errorMessage) > 0 {
		ss.minError = errorMessage[0]
	}
	return ss
}

// helper method for string min lenght constraint
func (ss *stringSchema) validateMinLength(str string, errorMessage ...string) error {
	if ss.minLength > 0 && len(str) < ss.minLength {
		message := ss.minError
		if len(errorMessage) > 0 {
			message = errorMessage[0]
		} else if message == "" {
			message = "value must have a minimum length of " + strconv.Itoa(ss.minLength)
		}
		return errors.New(message)
	}
	return nil
}

// Max sets the maximum length constraint for the string.
//
// Parameters:
//   - length: The maximum allowed length for the string.
//   - errorMessage (optional): A custom error message to be used when validation fails.
//
// Returns:
//   - The stringSchema instance with the maximum length constraint set.
func (ss *stringSchema) Max(length int, errorMessage ...string) *stringSchema {
	ss.maxLength = length
	if len(errorMessage) > 0 {
		ss.maxError = errorMessage[0]
	}
	return ss
}

// helper method for string max lenght constraint
func (ss *stringSchema) validateMaxLength(str string, errorMessage ...string) error {
	if ss.maxLength > 0 && len(str) > ss.maxLength {
		message := ss.maxError
		if len(errorMessage) > 0 {
			message = errorMessage[0]
		} else if message == "" {
			message = "value must not exceed a maximum length of " + strconv.Itoa(ss.maxLength)
		}
		return errors.New(message)
	}
	return nil
}

// Email sets a custom error message for invalid email addresses.
//
// Parameters:
//   - errorMessage (optional): The custom error message to be used when validation fails.
//
// Returns:
//   - The stringSchema instance with the custom error message set for email validation.
func (ss *stringSchema) Email(errorMessage ...string) *stringSchema {
	ss.emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if len(errorMessage) > 0 {
		ss.emailError = errorMessage[0]
	}
	return ss
}

// helper method for email validation
func (ss *stringSchema) validateEmail(str string, errorMessage ...string) error {
	if !ss.emailRegex.MatchString(str) {
		message := ss.emailError
		if len(errorMessage) > 0 {
			message = errorMessage[0]
		}
		return errors.New(message)
	}
	return nil
}
