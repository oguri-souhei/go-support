package support

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexBy(t *testing.T) {
	type user struct {
		id   int
		name string
		age  int
	}
	users := []user{
		{id: 1, name: "tanaka", age: 40},
		{id: 2, name: "kato", age: 40},
		{id: 3, name: "tanaka", age: 20},
	}
	t.Run("index by integer", func(t *testing.T) {
		got := IndexBy(users, func(u user) int { return u.id })
		want := map[int]user{
			1: {id: 1, name: "tanaka", age: 40},
			2: {id: 2, name: "kato", age: 40},
			3: {id: 3, name: "tanaka", age: 20},
		}
		assert.Equal(t, want, got)
	})

	t.Run("index by integer (duplicated key exsits)", func(t *testing.T) {
		got := IndexBy(users, func(u user) int { return u.age })
		want := map[int]user{
			40: {id: 2, name: "kato", age: 40},
			20: {id: 3, name: "tanaka", age: 20},
		}
		assert.Equal(t, want, got)
	})

	t.Run("index by struct", func(t *testing.T) {
		type key struct{ nm string }
		got := IndexBy(users, func(u user) key { return key{nm: u.name} })
		want := map[key]user{
			{nm: "tanaka"}: {id: 1, name: "tanaka", age: 40},
			{nm: "kato"}:   {id: 2, name: "kato", age: 40},
			{nm: "tanaka"}: {id: 3, name: "tanaka", age: 20},
		}
		assert.Equal(t, want, got)
	})

	t.Run("index by struct (duplicated key exists)", func(t *testing.T) {
		type key struct{ age int }
		got := IndexBy(users, func(u user) key { return key{age: u.age} })
		want := map[key]user{
			{age: 40}: {id: 2, name: "kato", age: 40},
			{age: 20}: {id: 3, name: "tanaka", age: 20},
		}
		assert.Equal(t, want, got)
	})
}

func TestGroupBy(t *testing.T) {
	type user struct {
		id   int
		name string
		age  int
	}
	users := []user{
		{id: 1, name: "tanaka", age: 40},
		{id: 2, name: "kato", age: 40},
		{id: 3, name: "tanaka", age: 20},
	}
	t.Run("group by id", func(t *testing.T) {
		got := GroupBy(users, func(u user) int {
			return u.id
		})
		want := map[int]user{
			1: {id: 1, name: "tanaka", age: 40},
			2: {id: 2, name: "kato", age: 40},
			3: {id: 3, name: "tanaka", age: 20},
		}
		assert.Equal(t, want, got)
	})
}

func TestDig(t *testing.T) {
	type args struct {
		in   any
		keys []any
	}
	type want struct {
		result any
		found  bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "simple array",
			args: args{in: []int{1, 2, 3}, keys: []any{1}},
			want: want{
				result: 2,
				found:  true,
			},
		},
		{
			name: "simple array (not found)",
			args: args{
				in:   []int{1, 2, 3},
				keys: []any{3},
			},
			want: want{
				result: nil,
				found:  false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, found := Dig(tt.args.in, tt.args.keys...)
			assert.Equal(t, tt.want.result, got)
			assert.Equal(t, tt.want.found, found)
		})
	}
}
