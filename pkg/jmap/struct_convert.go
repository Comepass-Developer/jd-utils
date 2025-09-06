package jmap

import (
	"reflect"
)

// Function StructToMap convert a struct to map and ignore nil value field
func StructToMap[T any](req T) map[string]interface{} {
	// Get value slice
	v := reflect.ValueOf(req)
	// Get type slice
	t := v.Type()
	// Init result map
	result := make(map[string]interface{})
	// Loop through slice fields
	for i := 0; i < v.NumField(); i++ {
		// Get field name and value
		fieldName := t.Field(i).Name
		fieldValue := v.Field(i)
		// Check if field value is not zero and not string or bool
		if !fieldValue.IsZero() && fieldValue.Kind() != reflect.String && fieldValue.Kind() != reflect.Bool {
			// Check if field value is pointer
			if fieldValue.Kind() == reflect.Ptr {
				result[fieldName] = fieldValue.Elem().Interface()
			} else {
				result[fieldName] = fieldValue.Interface()
			}
		}
	}
	return result
}
