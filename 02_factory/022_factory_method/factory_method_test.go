package factory_method_test

import (
	factory_method "github.com/1ikc/design-pattern/02_factory/022_factory_method"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewConfigParserFactory(t *testing.T) {
	type args struct {
		t string
	}
	cases := []struct{
		name string
		arg args
		want factory_method.IConfigParserFactory
	}{
		{
			name: "json",
			arg: args{t: "json"},
			want: factory_method.JsonConfigParserFactory{},
		},
		{
			name: "yaml",
			arg: args{t: "yaml"},
			want: factory_method.YamlConfigParserFactory{},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := factory_method.NewConfigParserFactory(c.arg.t); !assert.Equal(t, c.want, got) {
				t.Errorf("NewIRuleConfigParserFactory() = %v, want %v", got, c.want)
			}
		})
	}
}