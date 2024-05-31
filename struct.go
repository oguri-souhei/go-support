package support

import (
	"reflect"
)

func StructFields(v any) []string {
	rt := reflect.TypeOf(v)
	if rt.Kind() == reflect.Pointer {
		rt = rt.Elem()
	}
	if rt.Kind() != reflect.Struct {
		panic("arg must be struct")
	}

	var fields []string
	for i := 0; i < rt.NumField(); i++ {
		fields = append(fields, rt.Field(i).Name)
	}
	return fields
}
