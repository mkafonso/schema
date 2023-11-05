package z_test

import (
	"testing"

	"github.com/mkafonso/go-schema/z"
	"github.com/stretchr/testify/assert"
)

func TestStringSchemaTestValidString(t *testing.T) {
	schema := z.NewStringSchema()
	result, err := schema.Parse("Luna")

	assert.Equal(t, "Luna", result)
	assert.NoError(t, err)
}

func TestStringSchemaTestParseWithDefaultErrorMessage(t *testing.T) {
	schema := z.NewStringSchema()
	result, err := schema.Parse(42)

	assert.Equal(t, "", result)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "value must be a string")
}

func TestStringSchemaTestParseWithCustomErrorMessage(t *testing.T) {
	schema := z.NewStringSchema()
	result, err := schema.Parse(42, "Custom error message")

	assert.Equal(t, "", result)
	assert.Equal(t, err.Error(), "Custom error message")
}

func TestStringSchemaTestParseWithManyCustomErrorMessages(t *testing.T) {
	schema := z.NewStringSchema()
	result, err := schema.Parse(42, "First custom error message", "Second custom error message")

	assert.Equal(t, "", result)
	assert.Equal(t, err.Error(), "First custom error message")
}

func TestStringSchemaTestValidMinLengthConstraint(t *testing.T) {
	schema := z.NewStringSchema()
	result, err := schema.Min(2, "length must be at least 2 characters").Parse("Luna")

	assert.Equal(t, result, "Luna")
	assert.NoError(t, err)
}

func TestStringSchemaTestInValidMinLengthConstraint(t *testing.T) {
	schema := z.NewStringSchema()
	result, err := schema.Min(5, "length must be at least 5 characters").Parse("Hi")

	assert.Equal(t, result, "")
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "length must be at least 5 characters")
}

func TestStringSchemaTestValidMaxLengthConstraint(t *testing.T) {
	schema := z.NewStringSchema()
	result, err := schema.Max(10, "length must not exceed 2 characters").Parse("Luna")

	assert.Equal(t, result, "Luna")
	assert.NoError(t, err)
}

func TestStringSchemaTestInValidMaxLengthConstraint(t *testing.T) {
	schema := z.NewStringSchema()
	result, err := schema.Max(2, "length must not exceed 2 characters").Parse("Luna")

	assert.Equal(t, result, "")
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "length must not exceed 2 characters")
}

func TestStringSchemaTestValidEmail(t *testing.T) {
	schema := z.NewStringSchema()
	result, err := schema.Email("custom error message").Parse("email@email.com")

	assert.Equal(t, result, "email@email.com")
	assert.NoError(t, err)
}

func TestStringSchemaTestInvalidEmail(t *testing.T) {
	schema := z.NewStringSchema()
	result, err := schema.Email("custom error message").Parse("invalid-email.com")

	assert.Equal(t, result, "")
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "custom error message")
}

func TestStringSchemaTestValidEmailAndIncorrectLength(t *testing.T) {
	schema := z.NewStringSchema()
	result, err := schema.Min(100, "length custom error").Email("email error message").Parse("me@there.com")

	assert.Equal(t, result, "")
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "length custom error")
}

func TestStringSchemaTestInValidEmailAndCorrectLength(t *testing.T) {
	schema := z.NewStringSchema()
	result, err := schema.Min(1, "length custom error").Email("email error message").Parse("invalid-email.com")

	assert.Equal(t, result, "")
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "email error message")
}
