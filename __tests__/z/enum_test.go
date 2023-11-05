package z_test

import (
	"testing"

	"github.com/mkafonso/go-schema/z"
	"github.com/stretchr/testify/assert"
)

func TestEnumSchemaTestValidValue(t *testing.T) {
	schema := z.NewEnumSchema("apple", "banana")
	result := schema.Parse("apple")

	assert.True(t, result)
}

func TestEnumSchemaTestIncensitiveCase(t *testing.T) {
	schema := z.NewEnumSchema("apple", "banana")
	result := schema.Parse("BANANA")

	assert.True(t, result)
}

func TestEnumSchemaTestInvalidValue(t *testing.T) {
	schema := z.NewEnumSchema("apple", "banana")
	result := schema.Parse("yellow")

	assert.False(t, result)
}

func TestEnumSchemaTestGetEnumList(t *testing.T) {
	schema := z.NewEnumSchema("apple", "banana")
	result := schema.EnumList()

	expectedList := []string{"apple", "banana"}

	assert.ElementsMatch(t, result, expectedList)
}
