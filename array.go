package support

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/thoas/go-funk"
)

/*
任意の型のリストとキーの文字列を受け取り、指定されたフィールドの値をキーとしたマップを返す。

注意点:

  - キーのフィールドはエクスポートされている必要がある
  - キーのフィールドが存在しない場合はpanicになる
*/
func IndexBy[T any](list []T, fieldName string) map[any][]T {
	if reflect.TypeOf(list).Elem().Kind() != reflect.Struct {
		panic("T must be struct")
	}

	result := make(map[any][]T)

	for _, item := range list {
		field := reflect.ValueOf(item).FieldByName(fieldName)
		if !field.IsValid() {
			panic(fmt.Sprintf("no such field %s", fieldName))
		}

		fieldValue := field.Interface()
		result[fieldValue] = append(result[fieldValue], item)
	}

	return result
}

var directions = []string{"asc", "desc"}

/*
OrderByは、任意の型のリストを指定されたフィールドと方向（"asc" または "desc"）でソートします。
フィールドの型は int, uint, float, string のいずれかである必要があります。
方向が "asc" の場合は昇順、"desc" の場合は降順にソートされます。
*/
func OrderBy[T any](list []T, fieldName, direction string) []T {
	if !funk.Contains(directions, direction) {
		panic("direction must be \"asc\" or \"desc\"")
	}

	sort.SliceStable(list, func(i, j int) bool {
		valI := reflect.ValueOf(list[i])
		valJ := reflect.ValueOf(list[j])

		fieldI := valI.FieldByName(fieldName)
		fieldJ := valJ.FieldByName(fieldName)

		if !fieldI.IsValid() || !fieldJ.IsValid() {
			panic(fmt.Sprintf("no such field %s", fieldName))
		}

		var less bool
		switch fieldI.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			less = fieldI.Int() < fieldJ.Int()
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			less = fieldI.Uint() < fieldJ.Uint()
		case reflect.Float32, reflect.Float64:
			less = fieldI.Float() < fieldJ.Float()
		case reflect.String:
			less = fieldI.String() < fieldJ.String()
		default:
			panic("Unsupported field type")
		}

		if direction == "desc" {
			return !less
		}
		return less
	})

	return list
}
