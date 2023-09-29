package z

import (
	"errors"
	"fmt"
	"reflect"
)

// Parse function parses data from a map into a target struct based on JSON field tags.
//
// Parameters:
//   - data: A map[string]interface{} containing the data to be parsed.
//   - target: A pointer to a struct where the parsed values will be set.
//
// Returns:
//   - An error if there are issues with parsing or setting values.
//
// The Parse function iterates through the fields of the target struct, matches them
// with field names in the provided data, and assigns the values based on their types.
//
// Example usage:
//
//	data := map[string]interface{}{
//	  "name": "Alice",
//	  "age":  30,
//	}
//	person := Person{}
//	err := Parse(data, &person)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	// Now, the 'person' struct will contain the parsed values.
func ParseStruct(data map[string]interface{}, target interface{}) error {
	// get the type of the target object.
	targetType := reflect.TypeOf(target)

	// ensure that the target is a pointer to a struct.
	if targetType.Kind() != reflect.Ptr || targetType.Elem().Kind() != reflect.Struct {
		return errors.New("target must be a pointer to a struct")
	}

	// create a new instance of the target object.
	targetValue := reflect.New(targetType.Elem()).Elem()

	// iterate through the fields of the target struct.
	for i := 0; i < targetType.Elem().NumField(); i++ {
		field := targetType.Elem().Field(i)
		fieldName := field.Tag.Get("json")

		// check if the field name exists in the data.
		value, ok := data[fieldName]
		if !ok {
			return fmt.Errorf("missing field: %s", fieldName)
		}

		// set the field value in the target object.
		fieldValue := targetValue.FieldByName(field.Name)
		if !fieldValue.IsValid() {
			return fmt.Errorf("invalid field: %s", fieldName)
		}

		if !fieldValue.CanSet() {
			return fmt.Errorf("cannot set field: %s", fieldName)
		}

		// convert and set the value based on its type.
		if value != nil {
			valueType := reflect.TypeOf(value)
			fieldType := fieldValue.Type()
			if valueType.ConvertibleTo(fieldType) {
				fieldValue.Set(reflect.ValueOf(value).Convert(fieldType))
			} else {
				return fmt.Errorf("field type mismatch: %s", fieldName)
			}
		}
	}

	// Set the target object to the parsed values.
	reflect.ValueOf(target).Elem().Set(targetValue)

	return nil
}
