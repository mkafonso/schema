package z_test

import (
	"testing"

	"github.com/mkafonso/go-schema/z"
	"github.com/stretchr/testify/assert"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestParseStruct_ValidData(t *testing.T) {
	data := map[string]interface{}{
		"name": "Alice",
		"age":  30,
	}

	person := Person{}
	err := z.ParseStruct(data, &person)

	assert.NoError(t, err)
	assert.Equal(t, "Alice", person.Name)
	assert.Equal(t, 30, person.Age)
}

func TestParseStruct_MissingFieldAge(t *testing.T) {
	data := map[string]interface{}{
		"name": "Bob",
	}

	person := Person{}
	err := z.ParseStruct(data, &person)

	assert.Error(t, err)
	assert.Equal(t, err.Error(), "missing field: age")
}

func TestParseStruct_InvalidFieldTypeAge(t *testing.T) {
	data := map[string]interface{}{
		"name": "Charlie",
		"age":  "invalid-type",
	}

	person := Person{}
	err := z.ParseStruct(data, &person)

	assert.Error(t, err)
	assert.Equal(t, err.Error(), "field type mismatch: age")
}

func TestParseStruct_InvalidTarget(t *testing.T) {
	data := map[string]interface{}{
		"name": "David",
		"age":  25,
	}

	// pass a non-pointer target, which should result in an error.
	person := Person{}
	err := z.ParseStruct(data, person)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "target must be a pointer to a struct")
}
