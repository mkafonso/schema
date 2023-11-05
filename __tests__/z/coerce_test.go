package z_test

import (
	"testing"

	"github.com/mkafonso/go-schema/z"
	"github.com/stretchr/testify/assert"
)

func TestCoerceSchemaTestWithStringToString(t *testing.T) {
	schema := z.NewCoerceSchema()
	result, err := schema.String().Parse("Luna")

	assert.Equal(t, "Luna", result)
	assert.NoError(t, err)
}

func TestCoerceSchemaTestWithNumberToString(t *testing.T) {
	schema := z.NewCoerceSchema()
	result, err := schema.String().Parse(12)

	assert.Equal(t, "12", result)
	assert.NoError(t, err)
}

func TestCoerceSchemaTestWithBoolToString(t *testing.T) {
	schema := z.NewCoerceSchema()
	result, err := schema.String().Parse(true)

	assert.Equal(t, "true", result)
	assert.NoError(t, err)
}

func TestCoerceSchemaTestWithStringToBool(t *testing.T) {
	schema := z.NewCoerceSchema()
	result, err := schema.Bool().Parse("Luna")

	assert.Equal(t, true, result)
	assert.NoError(t, err)
}

func TestCoerceSchemaTestWithStringToBool1(t *testing.T) {
	schema := z.NewCoerceSchema()
	result, err := schema.Bool().Parse("true")

	assert.Equal(t, true, result)
	assert.NoError(t, err)
}

func TestCoerceSchemaTestWithStringToBool2(t *testing.T) {
	schema := z.NewCoerceSchema()
	result, err := schema.Bool().Parse("FalSe")

	assert.Equal(t, true, result)
	assert.NoError(t, err)
}

func TestCoerceSchemaTestWithEmptyStringToBool(t *testing.T) {
	schema := z.NewCoerceSchema()
	result, err := schema.Bool().Parse("")

	assert.Equal(t, false, result)
	assert.NoError(t, err)
}

func TestCoerceSchemaTestWithNilToBool(t *testing.T) {
	schema := z.NewCoerceSchema()
	result, err := schema.Bool().Parse(nil)

	assert.Equal(t, false, result)
	assert.NoError(t, err)
}

func TestCoerceSchemaTest1ToBool(t *testing.T) {
	schema := z.NewCoerceSchema()
	result, err := schema.Bool().Parse(1)

	assert.Equal(t, true, result)
	assert.NoError(t, err)
}

func TestCoerceSchemaTest0ToBool(t *testing.T) {
	schema := z.NewCoerceSchema()
	result, err := schema.Bool().Parse(0)

	assert.Equal(t, false, result)
	assert.NoError(t, err)
}

func TestCoerceSchemaTestFalseToBool(t *testing.T) {
	schema := z.NewCoerceSchema()
	result, err := schema.Bool().Parse(false)

	assert.Equal(t, false, result)
	assert.NoError(t, err)
}

func TestCoerceSchemaTestTrueToBool(t *testing.T) {
	schema := z.NewCoerceSchema()
	result, err := schema.Bool().Parse(true)

	assert.Equal(t, true, result)
	assert.NoError(t, err)
}
