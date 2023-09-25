package z_test

import (
	"testing"

	"github.com/mkafonso/go-schema/z"
	"github.com/stretchr/testify/assert"
)

func TestIsNumber_withValidNumber(t *testing.T) {
	validNumber := 42.5
	err := z.IsNumber(validNumber)
	assert.NoError(t, err)
}

func TestIsNumber_withInvalidNumber(t *testing.T) {
	invalidValue := "not a number"
	err := z.IsNumber(invalidValue)
	assert.Error(t, err, "Expected an error for an invalid value")
	assert.Contains(t, err.Error(), "value must be a number")
}
