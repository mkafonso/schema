package z_test

import (
	"testing"

	"github.com/mkafonso/go-schema/z"
	"github.com/stretchr/testify/assert"
)

func TestNumberSchema_TestValidNumber(t *testing.T) {
	schema := z.NewNumberSchema()
	result, err := schema.Parse(42.5)

	assert.Equal(t, 42.5, result)
	assert.NoError(t, err)
}

func TestNumberSchema_TestParseWithDefaultErrorMessage(t *testing.T) {
	schema := z.NewNumberSchema()
	result, err := schema.Parse("Hi")

	assert.Equal(t, 0.0, result)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "value must be a number")
}

func TestNumberSchema_TestParseWithCustomErrorMessage(t *testing.T) {
	schema := z.NewNumberSchema()
	result, err := schema.Parse("Hi", "Custom error message")

	assert.Equal(t, 0.0, result)
	assert.EqualError(t, err, "Custom error message")
}

func TestNumberSchema_TestParseWithManyCustomErrorMessages(t *testing.T) {
	schema := z.NewNumberSchema()
	result, err := schema.Parse("Hi", "First custom error message", "Second custom error message")

	assert.Equal(t, 0.0, result)
	assert.Equal(t, err.Error(), "First custom error message")
}
