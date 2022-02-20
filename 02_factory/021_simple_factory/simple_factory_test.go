package factory_test

import (
	factory "github.com/1ikc/design-pattern/02_factory/021_simple_factory"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewConfigParser(t *testing.T) {
	type args struct {
		t string
	}
	cases := []struct{
		name string
		arg args
		want factory.IConfigParser
	}{
		{
			name: "json",
			arg: args{t: "json"},
			want: factory.JsonConfigParser{},
		},
		{
			name: "yaml",
			arg: args{t: "yaml"},
			want: factory.YamlConfigParser{},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := factory.NewConfigParser(c.arg.t); !assert.Equal(t, c.want, got) {
				t.Errorf("NewConfigParser() = %v, want %v", got, c.want)
			}
		})
	}
}