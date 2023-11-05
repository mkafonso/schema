package z

import (
	"strings"
)

type enumSchema struct {
	Options map[string]bool
}

// NewEnumSchema creates a new instance of the Enum.
//
// Parameters:
//   - options: Variadic parameter that defines valid enum options as strings.
//
// Returns:
//   - A pointer to the newly created Enum instance.
func NewEnumSchema(options ...string) *enumSchema {
	enum := &enumSchema{
		Options: make(map[string]bool),
	}
	for _, option := range options {
		enum.Options[option] = true
	}
	return enum
}

// Parse checks if the provided value is a valid enum option.
// It is case-insensitive, so "apple" and "Apple" are considered equal.
//
// Parameters:
//   - value: The value to check against the enum.
//
// Returns:
//   - true if the value is a valid enum option, false otherwise.
func (e *enumSchema) Parse(value string) bool {
	_, ok := e.Options[strings.ToLower(value)]
	return ok
}

// EnumList returns a list of valid enum options.
//
// Returns:
//   - A slice of strings containing the valid enum options.
func (e *enumSchema) EnumList() []string {
	options := make([]string, 0, len(e.Options))
	for option := range e.Options {
		options = append(options, option)
	}
	return options
}
