package support

import (
	"reflect"

	"github.com/thoas/go-funk"
)

func IndexBy[T any, S comparable](in []T, fn func(T) S) map[S]T {
	result := make(map[S]T, len(in))
	for _, elem := range in {
		key := fn(elem)
		result[key] = elem
	}
	return result
}

func GroupBy[T any, S comparable](in []T, fn func(elem T) S) map[S][]T {
	result := make(map[S][]T)
	for _, elem := range in {
		key := fn(elem)
		result[key] = append(result[key], elem)
	}
	return result
}

// func OrderBy[T any](list []T, fn func(a, b T) bool) []T {
// }

func Dig(in any, keys ...any) (result any, found bool) {
	current := reflect.ValueOf(in)

	for _, key := range keys {
		indexValue := reflect.ValueOf(key)
		switch current.Kind() {
		case reflect.Array, reflect.Slice:
			// 配列のIndexを求める（Int想定）
			if !indexValue.CanInt() {
				return nil, false
			}
			index := int(indexValue.Int())
			if index >= current.Len() {
				return nil, false
			}
			current = current.Index(index)

		case reflect.Map:
			if !funk.Contains(current.MapKeys(), indexValue) {
				return nil, false
			}
			current = current.MapIndex(indexValue)

		case reflect.Struct:
			fields := StructFields(current)
			fieldName := indexValue.String()
			if !funk.Contains(fields, fieldName) {
				return nil, false
			}

			current = current.FieldByName(fieldName)

		default:
			return nil, false
		}
	}

	return current.Interface(), true
}
