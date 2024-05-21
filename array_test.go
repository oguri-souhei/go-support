package support

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexBy(t *testing.T) {
	type user struct {
		ID    int
		Name  string
		email string
	}

	type args struct {
		list []user
		key  string
	}
	tests := []struct {
		name  string
		args  args
		want  map[any][]user
		panic bool
	}{
		{
			name: "index by name",
			args: args{
				list: []user{
					{ID: 1, Name: "test1"},
					{ID: 2, Name: "test2"},
					{ID: 3, Name: "test1"},
				},
				key: "Name",
			},
			want: map[any][]user{
				"test1": {
					{ID: 1, Name: "test1"},
					{ID: 3, Name: "test1"},
				},
				"test2": {
					{ID: 2, Name: "test2"},
				},
			},
		},
		{
			name: "index by ID",
			args: args{
				list: []user{
					{ID: 1, Name: "test1"},
					{ID: 2, Name: "test2"},
					{ID: 1, Name: "test3"},
				},
				key: "ID",
			},
			want: map[any][]user{
				1: {
					{ID: 1, Name: "test1"},
					{ID: 1, Name: "test3"},
				},
				2: {
					{ID: 2, Name: "test2"},
				},
			},
		},
		{
			name: "fieldName not found",
			args: args{
				list: []user{
					{ID: 1, Name: "test1"},
					{ID: 2, Name: "test2"},
					{ID: 3, Name: "test1"},
				},
				key: "Age",
			},
			panic: true,
		},
		{
			name: "field not exported",
			args: args{
				list: []user{
					{ID: 1, Name: "test1", email: "test-1@test.com"},
					{ID: 2, Name: "test2", email: "test-2@test.com"},
					{ID: 3, Name: "test1", email: "test-1@test.com"},
				},
				key: "email",
			},
			panic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.panic {
				assert.Panics(t, func() {
					IndexBy(tt.args.list, tt.args.key)
				})
			} else {
				got := IndexBy(tt.args.list, tt.args.key)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func ExampleIndexBy() {
	type Person struct {
		Name string
		Age  int
	}

	people := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Alice", Age: 28},
	}

	fmt.Println(IndexBy(people, "Name"))
	// Output: map[Alice:[{Alice 30} {Alice 28}] Bob:[{Bob 25}]]
}

func TestOrderBy(t *testing.T) {
	type User struct {
		ID     int
		Name   string
		Age    uint
		Height float32
		weight float32
	}
	type args struct {
		list      []User
		fieldName string
		direction string
	}
	tests := []struct {
		name  string
		args  args
		want  []User
		panic bool
	}{
		{
			name: "order by int field",
			args: args{
				list: []User{
					{ID: 3},
					{ID: 1},
					{ID: 2},
				},
				fieldName: "ID",
				direction: "asc",
			},
			want: []User{
				{ID: 1},
				{ID: 2},
				{ID: 3},
			},
		},
		{
			name: "order by float field",
			args: args{
				list: []User{
					{Height: 180.1},
					{Height: 160.2},
					{Height: 170.8},
				},
				fieldName: "Height",
				direction: "desc",
			},
			want: []User{
				{Height: 180.1},
				{Height: 170.8},
				{Height: 160.2},
			},
		},
		{
			name: "order by uint field",
			args: args{
				list: []User{
					{Age: 30},
					{Age: 40},
					{Age: 20},
				},
				fieldName: "Age",
				direction: "asc",
			},
			want: []User{
				{Age: 20},
				{Age: 30},
				{Age: 40},
			},
		},
		{
			name: "order by string field",
			args: args{
				list: []User{
					{Name: "test-2"},
					{Name: "test-1"},
					{Name: "test-3"},
				},
				fieldName: "Name",
				direction: "asc",
			},
			want: []User{
				{Name: "test-1"},
				{Name: "test-2"},
				{Name: "test-3"},
			},
		},
		{
			name: "field not exported",
			args: args{
				list: []User{
					{weight: 60.1},
					{weight: 40.1},
					{weight: 70.1},
				},
				fieldName: "weight",
				direction: "asc",
			},
			want: []User{
				{weight: 40.1},
				{weight: 60.1},
				{weight: 70.1},
			},
		},
		{
			name: "field not found",
			args: args{
				list: []User{
					{},
					{},
				},
				fieldName: "Hoge",
				direction: "asc",
			},
			panic: true,
		},
		{
			name: "direction is not \"asc\" or \"desc\"",
			args: args{
				list: []User{
					{ID: 1},
					{ID: 2},
				},
				fieldName: "ID",
				direction: "hoge",
			},
			panic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.panic {
				assert.Panics(t, func() {
					OrderBy(tt.args.list, tt.args.fieldName, tt.args.direction)
				})
			} else {
				got := OrderBy(tt.args.list, tt.args.fieldName, tt.args.direction)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func ExampleOrderBy() {
	type Person struct {
		Name string
		Age  int
	}

	people := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 28},
	}

	fmt.Println(OrderBy(people, "Age", "asc"))
	// Output: [{Bob 25} {Charlie 28} {Alice 30}]
}
