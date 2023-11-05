package z_test

import (
	"strings"
	"testing"

	"github.com/mkafonso/go-schema/z"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestParseStructValidData(t *testing.T) {
	data := map[string]interface{}{
		"name": "Alice",
		"age":  30,
	}

	person := Person{}
	err := z.ParseStruct(data, &person)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if person.Name != "Alice" {
		t.Errorf("Expected Name to be 'Alice', but got: %s", person.Name)
	}
	if person.Age != 30 {
		t.Errorf("Expected Age to be 30, but got: %d", person.Age)
	}
}

func TestParseStructMissingFieldAge(t *testing.T) {
	data := map[string]interface{}{
		"name": "Bob",
	}

	person := Person{}
	err := z.ParseStruct(data, &person)

	if err == nil || err.Error() != "missing field: age" {
		t.Errorf("Expected 'missing field: age' error, but got: %v", err)
	}
}

func TestParseStructInvalidFieldTypeAge(t *testing.T) {
	data := map[string]interface{}{
		"name": "Charlie",
		"age":  "invalid-type",
	}

	person := Person{}
	err := z.ParseStruct(data, &person)

	if err == nil || err.Error() != "field type mismatch: age" {
		t.Errorf("Expected 'field type mismatch: age' error, but got: %v", err)
	}
}

func TestParseStructInvalidTarget(t *testing.T) {
	data := map[string]interface{}{
		"name": "David",
		"age":  25,
	}

	// Pass a non-pointer target, which should result in an error.
	person := Person{}
	err := z.ParseStruct(data, person)

	if err == nil || !strings.Contains(err.Error(), "target must be a pointer to a struct") {
		t.Errorf("Expected 'target must be a pointer to a struct' error, but got: %v", err)
	}
}
