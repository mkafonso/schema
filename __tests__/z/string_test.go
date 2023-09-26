package z_test

import (
	"testing"

	"github.com/mkafonso/go-schema/z"
	"github.com/stretchr/testify/assert"
)

func TestStringSchema_TestValidString(t *testing.T) {
	schema := z.NewStringSchema()

	result, err := schema.Parse("Luna")
	assert.Equal(t, "Luna", result)
	assert.NoError(t, err)
}

func TestStringSchema_TestParseWithCustomErrorMessage(t *testing.T) {
	schema := z.NewStringSchema()

	result, err := schema.Parse(42, "Custom error message")
	assert.Equal(t, "", result)
	assert.EqualError(t, err, "Custom error message")
}

func TestStringSchema_TestParseWithManyCustomErrorMessages(t *testing.T) {
	schema := z.NewStringSchema()

	result, err := schema.Parse(42, "First custom error message", "Second custom error message")
	assert.Equal(t, "", result)
	assert.EqualError(t, err, "First custom error message")
}

func TestStringSchema_TestValidMinLengthConstraint(t *testing.T) {
	schema := z.NewStringSchema()
	result, err := schema.Min(2, "length must be at least 2 characters").Parse("Luna")

	assert.Equal(t, result, "Luna")
	assert.NoError(t, err)
}

func TestStringSchema_TestInValidMinLengthConstraint(t *testing.T) {
	schema := z.NewStringSchema()
	result, err := schema.Min(5, "length must be at least 5 characters").Parse("Hi")

	assert.Equal(t, result, "")
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "length must be at least 5 characters")
}

func TestStringSchema_TestValidMaxLengthConstraint(t *testing.T) {
	schema := z.NewStringSchema()
	result, err := schema.Max(10, "length must not exceed 2 characters").Parse("Luna")

	assert.Equal(t, result, "Luna")
	assert.NoError(t, err)
}

func TestStringSchema_TestInValidMaxLengthConstraint(t *testing.T) {
	schema := z.NewStringSchema()
	result, err := schema.Max(2, "length must not exceed 2 characters").Parse("Luna")

	assert.Equal(t, result, "")
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "length must not exceed 2 characters")
}

func TestStringSchema_TestValidEmail(t *testing.T) {
	schema := z.NewStringSchema()
	result, err := schema.Email("custom error message").Parse("email@email.com")

	assert.Equal(t, result, "email@email.com")
	assert.NoError(t, err)
}

func TestStringSchema_TestInvalidEmail(t *testing.T) {
	schema := z.NewStringSchema()
	result, err := schema.Email("custom error message").Parse("invalid-email.com")

	assert.Equal(t, result, "")
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "custom error message")
}

func TestStringSchema_TestValidEmailAndIncorrectLength(t *testing.T) {
	schema := z.NewStringSchema()
	result, err := schema.Min(100, "length custom error").Email("email error message").Parse("me@there.com")

	assert.Equal(t, result, "")
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "length custom error")
}

func TestStringSchema_TestInValidEmailAndCorrectLength(t *testing.T) {
	schema := z.NewStringSchema()
	result, err := schema.Min(1, "length custom error").Email("email error message").Parse("invalid-email.com")

	assert.Equal(t, result, "")
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "email error message")
}
