package z_test

import (
	"testing"

	"github.com/mkafonso/go-schema/z"
)

func TestStringSchemaTestValidString(t *testing.T) {
	schema := z.NewStringSchema()
	result, err := schema.Parse("Luna")

	if result != "Luna" {
		t.Errorf("Expected result to be 'Luna', got '%s'", result)
	}
	if err != nil {
		t.Errorf("Expected error to be nil, got %v", err)
	}
}

func TestStringSchemaTestParseWithDefaultErrorMessage(t *testing.T) {
	schema := z.NewStringSchema()
	result, err := schema.Parse(42)

	if result != "" {
		t.Errorf("Expected result to be empty, got '%s'", result)
	}
	if err == nil {
		t.Error("Expected an error, but got nil")
	} else if err.Error() != "value must be a string" {
		t.Errorf("Expected error message to be 'value must be a string', got '%s'", err.Error())
	}
}

func TestStringSchemaTestParseWithCustomErrorMessage(t *testing.T) {
	schema := z.NewStringSchema()
	result, err := schema.Parse(42, "Custom error message")

	if result != "" {
		t.Errorf("Expected result to be empty, got '%s'", result)
	}
	if err == nil {
		t.Error("Expected an error, but got nil")
	} else if err.Error() != "Custom error message" {
		t.Errorf("Expected error message to be 'Custom error message', got '%s'", err.Error())
	}
}

func TestStringSchemaTestParseWithManyCustomErrorMessages(t *testing.T) {
	schema := z.NewStringSchema()
	result, err := schema.Parse(42, "First custom error message", "Second custom error message")

	if result != "" {
		t.Errorf("Expected result to be empty, got '%s'", result)
	}
	if err == nil {
		t.Error("Expected an error, but got nil")
	} else if err.Error() != "First custom error message" {
		t.Errorf("Expected error message to be 'First custom error message', got '%s'", err.Error())
	}
}

func TestStringSchemaTestValidMinLengthConstraint(t *testing.T) {
	schema := z.NewStringSchema()
	result, err := schema.Min(2, "length must be at least 2 characters").Parse("Luna")

	if result != "Luna" {
		t.Errorf("Expected result to be 'Luna', got '%s'", result)
	}
	if err != nil {
		t.Errorf("Expected error to be nil, got %v", err)
	}
}

func TestStringSchemaTestInValidMinLengthConstraint(t *testing.T) {
	schema := z.NewStringSchema()
	result, err := schema.Min(5, "length must be at least 5 characters").Parse("Hi")

	if result != "" {
		t.Errorf("Expected result to be empty, got '%s'", result)
	}
	if err == nil {
		t.Error("Expected an error, but got nil")
	} else if err.Error() != "length must be at least 5 characters" {
		t.Errorf("Expected error message to be 'length must be at least 5 characters', got '%s'", err.Error())
	}
}

func TestStringSchemaTestValidMaxLengthConstraint(t *testing.T) {
	schema := z.NewStringSchema()
	result, err := schema.Max(10, "length must not exceed 10 characters").Parse("Luna")

	if result != "Luna" {
		t.Errorf("Expected result to be 'Luna', got '%s'", result)
	}
	if err != nil {
		t.Errorf("Expected error to be nil, got %v", err)
	}
}

func TestStringSchemaTestInValidMaxLengthConstraint(t *testing.T) {
	schema := z.NewStringSchema()
	result, err := schema.Max(2, "length must not exceed 2 characters").Parse("Luna")

	if result != "" {
		t.Errorf("Expected result to be empty, got '%s'", result)
	}
	if err == nil {
		t.Error("Expected an error, but got nil")
	} else if err.Error() != "length must not exceed 2 characters" {
		t.Errorf("Expected error message to be 'length must not exceed 2 characters', got '%s'", err.Error())
	}
}

func TestStringSchemaTestValidEmail(t *testing.T) {
	schema := z.NewStringSchema()
	result, err := schema.Email("custom error message").Parse("email@email.com")

	if result != "email@email.com" {
		t.Errorf("Expected result to be 'email@email.com', got '%s'", result)
	}
	if err != nil {
		t.Errorf("Expected error to be nil, got %v", err)
	}
}

func TestStringSchemaTestInvalidEmail(t *testing.T) {
	schema := z.NewStringSchema()
	result, err := schema.Email("custom error message").Parse("invalid-email.com")

	if result != "" {
		t.Errorf("Expected result to be empty, got '%s'", result)
	}
	if err == nil {
		t.Error("Expected an error, but got nil")
	} else if err.Error() != "custom error message" {
		t.Errorf("Expected error message to be 'custom error message', got '%s'", err.Error())
	}
}

func TestStringSchemaTestValidEmailAndIncorrectLength(t *testing.T) {
	schema := z.NewStringSchema()
	result, err := schema.Min(100, "length custom error").Email("email error message").Parse("me@there.com")

	if result != "" {
		t.Errorf("Expected result to be empty, got '%s'", result)
	}
	if err == nil {
		t.Error("Expected an error, but got nil")
	} else if err.Error() != "length custom error" {
		t.Errorf("Expected error message to be 'length custom error', got '%s'", err.Error())
	}
}

func TestStringSchemaTestInValidEmailAndCorrectLength(t *testing.T) {
	schema := z.NewStringSchema()
	result, err := schema.Min(1, "length custom error").Email("email error message").Parse("invalid-email.com")

	if result != "" {
		t.Errorf("Expected result to be empty, got '%s'", result)
	}
	if err == nil {
		t.Error("Expected an error, but got nil")
	} else if err.Error() != "email error message" {
		t.Errorf("Expected error message to be 'email error message', got '%s'", err.Error())
	}
}
