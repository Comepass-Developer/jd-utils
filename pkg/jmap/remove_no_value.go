package jmap

import (
	"reflect"
	"strings"
)

func StructToMapNonZero[T any](s T) map[string]interface{} {
	result := make(map[string]interface{})
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		tag := t.Field(i).Tag.Get("json")
		if tag == "" || tag == "-" {
			continue
		}
		key := strings.Split(tag, ",")[0]

		if field.Kind() == reflect.Pointer {
			if field.IsNil() {
				continue
			}
			result[key] = field.Elem().Interface()
			continue
		}

		zero := reflect.Zero(field.Type()).Interface()
		if !reflect.DeepEqual(field.Interface(), zero) {
			result[key] = field.Interface()
		}
	}
	return result
}
