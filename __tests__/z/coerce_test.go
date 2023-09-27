package z_test

import (
	"testing"

	"github.com/mkafonso/go-schema/z"
	"github.com/stretchr/testify/assert"
)

func TestCoerceSchema_TestWithStringToString(t *testing.T) {
	schema := z.NewCoerceSchema()
	result, err := schema.String().Parse("Luna")

	assert.Equal(t, "Luna", result)
	assert.NoError(t, err)
}

func TestCoerceSchema_TestWithNumberToString(t *testing.T) {
	schema := z.NewCoerceSchema()
	result, err := schema.String().Parse(12)

	assert.Equal(t, "12", result)
	assert.NoError(t, err)
}

func TestCoerceSchema_TestWithBoolToString(t *testing.T) {
	schema := z.NewCoerceSchema()
	result, err := schema.String().Parse(true)

	assert.Equal(t, "true", result)
	assert.NoError(t, err)
}

func TestCoerceSchema_TestWithStringToBool(t *testing.T) {
	schema := z.NewCoerceSchema()
	result, err := schema.Bool().Parse("Luna")

	assert.Equal(t, true, result)
	assert.NoError(t, err)
}

func TestCoerceSchema_TestWithStringToBool_1(t *testing.T) {
	schema := z.NewCoerceSchema()
	result, err := schema.Bool().Parse("true")

	assert.Equal(t, true, result)
	assert.NoError(t, err)
}

func TestCoerceSchema_TestWithStringToBool_2(t *testing.T) {
	schema := z.NewCoerceSchema()
	result, err := schema.Bool().Parse("FalSe")

	assert.Equal(t, true, result)
	assert.NoError(t, err)
}

func TestCoerceSchema_TestWithEmptyStringToBool(t *testing.T) {
	schema := z.NewCoerceSchema()
	result, err := schema.Bool().Parse("")

	assert.Equal(t, false, result)
	assert.NoError(t, err)
}

func TestCoerceSchema_TestWithNilToBool(t *testing.T) {
	schema := z.NewCoerceSchema()
	result, err := schema.Bool().Parse(nil)

	assert.Equal(t, false, result)
	assert.NoError(t, err)
}

func TestCoerceSchema_Test1ToBool(t *testing.T) {
	schema := z.NewCoerceSchema()
	result, err := schema.Bool().Parse(1)

	assert.Equal(t, true, result)
	assert.NoError(t, err)
}

func TestCoerceSchema_Test0ToBool(t *testing.T) {
	schema := z.NewCoerceSchema()
	result, err := schema.Bool().Parse(0)

	assert.Equal(t, false, result)
	assert.NoError(t, err)
}

func TestCoerceSchema_TestFalseToBool(t *testing.T) {
	schema := z.NewCoerceSchema()
	result, err := schema.Bool().Parse(false)

	assert.Equal(t, false, result)
	assert.NoError(t, err)
}

func TestCoerceSchema_TestTrueToBool(t *testing.T) {
	schema := z.NewCoerceSchema()
	result, err := schema.Bool().Parse(true)

	assert.Equal(t, true, result)
	assert.NoError(t, err)
}
