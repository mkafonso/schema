package z_test

import (
	"testing"

	"github.com/mkafonso/go-schema/z"
)

func TestCoerceSchemaTestWithStringToString(t *testing.T) {
	schema := z.NewCoerceSchema()
	result, err := schema.String().Parse("Luna")

	if result != "Luna" {
		t.Errorf("Expected 'Luna', but got: %s", result)
	}
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestCoerceSchemaTestWithNumberToString(t *testing.T) {
	schema := z.NewCoerceSchema()
	result, err := schema.String().Parse(12)

	if result != "12" {
		t.Errorf("Expected '12', but got: %s", result)
	}
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestCoerceSchemaTestWithBoolToString(t *testing.T) {
	schema := z.NewCoerceSchema()
	result, err := schema.String().Parse(true)

	if result != "true" {
		t.Errorf("Expected 'true', but got: %s", result)
	}
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestCoerceSchemaTestWithStringToBool(t *testing.T) {
	schema := z.NewCoerceSchema()
	result, err := schema.Bool().Parse("Luna")

	if result != true {
		t.Errorf("Expected 'true', but got: %v", result)
	}
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestCoerceSchemaTestWithStringToBool1(t *testing.T) {
	schema := z.NewCoerceSchema()
	result, err := schema.Bool().Parse("true")

	if result != true {
		t.Errorf("Expected 'true', but got: %v", result)
	}
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestCoerceSchemaTestWithStringToBool2(t *testing.T) {
	schema := z.NewCoerceSchema()
	result, err := schema.Bool().Parse("FalSe")

	if result != true {
		t.Errorf("Expected 'true', but got: %v", result)
	}
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestCoerceSchemaTestWithEmptyStringToBool(t *testing.T) {
	schema := z.NewCoerceSchema()
	result, err := schema.Bool().Parse("")

	if result != false {
		t.Errorf("Expected 'false', but got: %v", result)
	}
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestCoerceSchemaTestWithNilToBool(t *testing.T) {
	schema := z.NewCoerceSchema()
	result, err := schema.Bool().Parse(nil)

	if result != false {
		t.Errorf("Expected 'false', but got: %v", result)
	}
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestCoerceSchemaTest1ToBool(t *testing.T) {
	schema := z.NewCoerceSchema()
	result, err := schema.Bool().Parse(1)

	if result != true {
		t.Errorf("Expected 'true', but got: %v", result)
	}
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestCoerceSchemaTest0ToBool(t *testing.T) {
	schema := z.NewCoerceSchema()
	result, err := schema.Bool().Parse(0)

	if result != false {
		t.Errorf("Expected 'false', but got: %v", result)
	}
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestCoerceSchemaTestFalseToBool(t *testing.T) {
	schema := z.NewCoerceSchema()
	result, err := schema.Bool().Parse(false)

	if result != false {
		t.Errorf("Expected 'false', but got: %v", result)
	}
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestCoerceSchemaTestTrueToBool(t *testing.T) {
	schema := z.NewCoerceSchema()
	result, err := schema.Bool().Parse(true)

	if result != true {
		t.Errorf("Expected 'true', but got: %v", result)
	}
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}
