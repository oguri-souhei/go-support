package support

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStructFields(t *testing.T) {
	type args struct {
		v any
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "exported field",
			args: args{
				v: struct{ Hoge, Huga string }{},
			},
			want: []string{"Hoge", "Huga"},
		},
		{
			name: "unexported field",
			args: args{
				v: struct{ hoge, huga string }{},
			},
			want: []string{"hoge", "huga"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StructFields(tt.args.v)
			assert.Equal(t, tt.want, got)
		})
	}
}
