package z_test

import (
	"testing"

	"github.com/mkafonso/go-schema/z"
	"github.com/stretchr/testify/assert"
)

func TestString_withValidString(t *testing.T) {
	// don't expect error when string is valid
	str := z.NewStringSchema()
	result, err := str.Parse("Luna")
	assert.Equal(t, result, "Luna")
	assert.NoError(t, err)
}

func TestString_withInvalidString(t *testing.T) {
	// expect an error when string is invalid
	str := z.NewStringSchema()
	result, err := str.Parse(12)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "value must be a string")
	assert.Equal(t, result, "")

	// expect an error when string is invalid and return custom error message
	str2 := z.NewStringSchema()
	result2, err2 := str2.Message("custom error message").Parse(12)
	assert.Equal(t, err2.Error(), "custom error message")
	assert.Equal(t, result2, "")

	// expect an error when string is invalid and return the last custom error message
	str3 := z.NewStringSchema()
	result3, err3 := str3.Message("first custom error message").Message("second custom error message").Parse(12)
	assert.Equal(t, err3.Error(), "second custom error message")
	assert.Equal(t, result3, "")
}

func TestString_testMinMethod(t *testing.T) {
	// expect an error when string length exceed Max value
	str := z.NewStringSchema()
	result, err := str.Max(3).Message("custom error message").Parse("Luna")
	assert.Equal(t, result, "")
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "custom error message")

	// expect an error when string length exceed Max value
	str2 := z.NewStringSchema()
	result2, err2 := str2.Min(5).Message("custom error message").Parse("Hi")
	assert.Equal(t, result2, "")
	assert.Error(t, err2)
	assert.Equal(t, err2.Error(), "custom error message")
}
