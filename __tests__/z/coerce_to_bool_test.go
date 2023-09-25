package z_test

import (
	"testing"

	"github.com/mkafonso/go-schema/z"
	"github.com/stretchr/testify/assert"
)

func TestCoerceToBool_TestDifferentTypesToBool(t *testing.T) {
	schema := z.NewCoerceBoolSchema()
	result, err := schema.Parse("true")
	assert.Equal(t, result, true)
	assert.NoError(t, err)

	schema2 := z.NewCoerceBoolSchema()
	result2, err2 := schema2.Parse("FalSe")
	assert.Equal(t, result2, true)
	assert.NoError(t, err2)

	schema3 := z.NewCoerceBoolSchema()
	result3, err3 := schema3.Parse(true)
	assert.Equal(t, result3, true)
	assert.NoError(t, err3)

	schema4 := z.NewCoerceBoolSchema()
	result4, err4 := schema4.Parse(false)
	assert.Equal(t, result4, false)
	assert.NoError(t, err4)

	schema5 := z.NewCoerceBoolSchema()
	result5, err5 := schema5.Parse("Hello")
	assert.Equal(t, result5, true)
	assert.NoError(t, err5)

	schema6 := z.NewCoerceBoolSchema()
	result6, err6 := schema6.Parse("")
	assert.Equal(t, result6, false)
	assert.NoError(t, err6)

	schema7 := z.NewCoerceBoolSchema()
	result7, err7 := schema7.Parse(10)
	assert.Equal(t, result7, true)
	assert.NoError(t, err7)

	schema8 := z.NewCoerceBoolSchema()
	result8, err8 := schema8.Parse(nil)
	assert.Equal(t, result8, false)
	assert.NoError(t, err8)

	schema9 := z.NewCoerceBoolSchema()
	result9, err9 := schema9.Parse(1)
	assert.Equal(t, result9, true)
	assert.NoError(t, err9)

	schema10 := z.NewCoerceBoolSchema()
	result10, err10 := schema10.Parse(0)
	assert.Equal(t, result10, false)
	assert.NoError(t, err10)

	schema11 := z.NewCoerceBoolSchema()
	result11, err11 := schema11.Parse([]interface{}{})
	assert.Equal(t, result11, true)
	assert.NoError(t, err11)
}
