package z_test

import (
	"testing"

	"github.com/mkafonso/go-schema/z"
)

func TestNumberSchemaTestValidNumber(t *testing.T) {
	schema := z.NewNumberSchema()
	result, err := schema.Parse(42.5)

	if result != 42.5 {
		t.Errorf("Expected result to be 42.5, got %v", result)
	}
	if err != nil {
		t.Errorf("Expected error to be nil, got %v", err)
	}
}

func TestNumberSchemaTestParseWithDefaultErrorMessage(t *testing.T) {
	schema := z.NewNumberSchema()
	result, err := schema.Parse("Hi")

	if result != 0.0 {
		t.Errorf("Expected result to be 0.0, got %v", result)
	}
	if err == nil {
		t.Error("Expected an error, but got nil")
	} else if err.Error() != "value must be a number" {
		t.Errorf("Expected error message to be 'value must be a number', got '%s'", err.Error())
	}
}

func TestNumberSchemaTestParseWithCustomErrorMessage(t *testing.T) {
	schema := z.NewNumberSchema()
	result, err := schema.Parse("Hi", "Custom error message")

	if result != 0.0 {
		t.Errorf("Expected result to be 0.0, got %v", result)
	}
	if err == nil {
		t.Error("Expected an error, but got nil")
	} else if err.Error() != "Custom error message" {
		t.Errorf("Expected error message to be 'Custom error message', got '%s'", err.Error())
	}
}

func TestNumberSchemaTestParseWithManyCustomErrorMessages(t *testing.T) {
	schema := z.NewNumberSchema()
	result, err := schema.Parse("Hi", "First custom error message", "Second custom error message")

	if result != 0.0 {
		t.Errorf("Expected result to be 0.0, got %v", result)
	}
	if err == nil {
		t.Error("Expected an error, but got nil")
	} else if err.Error() != "First custom error message" {
		t.Errorf("Expected error message to be 'First custom error message', got '%s'", err.Error())
	}
}
