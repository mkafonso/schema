package z_test

import (
	"testing"

	"github.com/mkafonso/go-schema/z"
	"github.com/stretchr/testify/assert"
)

func TestCoerceToString_TestDifferentTypesToString(t *testing.T) {
	schema := z.NewCoerceStringSchema()
	result, err := schema.Parse("Luna")
	assert.Equal(t, result, "Luna")
	assert.NoError(t, err)

	schema2 := z.NewCoerceStringSchema()
	result2, err2 := schema2.Parse(12)
	assert.Equal(t, result2, "12")
	assert.NoError(t, err2)

	schema3 := z.NewCoerceStringSchema()
	result3, err3 := schema3.Parse(true)
	assert.Equal(t, result3, "true")
	assert.NoError(t, err3)
}
