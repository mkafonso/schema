package z_test

import (
	"testing"

	"github.com/mkafonso/go-verify/z"
)

func TestEnumSchemaTestValidValue(t *testing.T) {
	schema := z.NewEnumSchema("apple", "banana")
	result := schema.Parse("apple")

	if result != true {
		t.Error("Expected true, got false")
	}
}

func TestEnumSchemaTestIncensitiveCase(t *testing.T) {
	schema := z.NewEnumSchema("apple", "banana")
	result := schema.Parse("BANANA")

	if result != true {
		t.Error("Expected true, got false")
	}
}

func TestEnumSchemaTestInvalidValue(t *testing.T) {
	schema := z.NewEnumSchema("apple", "banana")
	result := schema.Parse("yellow")

	if result != false {
		t.Error("Expected false, got true")
	}
}

func TestEnumSchemaTestGetEnumList(t *testing.T) {
	schema := z.NewEnumSchema("apple", "banana")
	result := schema.EnumList()

	expectedList := []string{"apple", "banana"}

	if len(result) != len(expectedList) {
		t.Errorf("Expected result length of %d, got %d", len(expectedList), len(result))
	} else {
		for i, val := range result {
			if val != expectedList[i] {
				t.Errorf("Expected '%s', got '%s'", expectedList[i], val)
			}
		}
	}
}
