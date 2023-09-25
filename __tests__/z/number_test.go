package z_test

import (
	"testing"

	"github.com/mkafonso/go-schema/z"
	"github.com/stretchr/testify/assert"
)

func TestNumber_withValidNumber(t *testing.T) {
	// don't expect error when number is valid
	number := z.NewNumberSchema()
	result, err := number.Parse(42.5)
	assert.Equal(t, result, 42.5)
	assert.NoError(t, err)
}

func TestNumber_withInvalidNumber(t *testing.T) {
	// expect an error when number is invalid
	number := z.NewNumberSchema()
	result, err := number.Parse("not a number")
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "value must be a number")
	assert.Equal(t, result, 0.0)

	// expect an error when number is invalid and return custom error message
	number2 := z.NewNumberSchema()
	result2, err2 := number2.Message("custom error message").Parse("not a number")
	assert.Equal(t, err2.Error(), "custom error message")
	assert.Equal(t, result2, 0.0)

	// expect an error when number is invalid and return the last custom error message
	number3 := z.NewNumberSchema()
	result3, err3 := number3.Message("first custom error message").Message("second custom error message").Parse("not a number")
	assert.Equal(t, err3.Error(), "second custom error message")
	assert.Equal(t, result3, 0.0)
}
