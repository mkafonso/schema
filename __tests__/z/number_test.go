package z_test

import (
	"testing"

	"github.com/mkafonso/go-schema/z"
	"github.com/stretchr/testify/assert"
)

func TestIsNumber_withValidNumber(t *testing.T) {
	// expect not error when number is valid
	value := 12
	err := z.IsNumber(value)
	assert.NoError(t, err)
}

func TestIsNumber_withInvalidNumber(t *testing.T) {
	// expect an error when number is invalid
	value := "not a number"
	err := z.IsNumber(value)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "value must be a number")

	// expect an error when number is invalid and return custom error message
	value2 := "not a number"
	err = z.IsNumber(value2, "número inválido")
	assert.Equal(t, err.Error(), "número inválido")

	// expect an error when number is invalid and return only the first custom error message
	invalidValue := "not a number"
	err = z.IsNumber(invalidValue, "número inválido", "another error message")
	assert.Equal(t, err.Error(), "número inválido")
}
