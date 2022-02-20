package abstract_factory_test

import (
	abstract_factory "github.com/1ikc/design-pattern/02_factory/023_abstract_factory"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJsonConfigParserFactory_CreateRuleConfigParser(t *testing.T) {
	cases := []struct{
		name string
		want abstract_factory.IRuleConfigParser
	}{
		{
			name: "json",
			want: abstract_factory.JsonRuleConfigParser{},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			f := abstract_factory.JsonConfigParserFactory{}
			if got := f.CreateRuleConfigParser(); !assert.Equal(t, c.want, got) {
				t.Errorf("CreateRuleParser() = %v, want %v", got, c.want)
			}
		})
	}
}

func TestJsonConfigParserFactory_CreateSystemConfigParser(t *testing.T) {
	cases := []struct{
		name string
		want abstract_factory.ISystemConfigParser
	}{
		{
			name: "json",
			want: abstract_factory.JsonSystemConfigParser{},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			f := abstract_factory.JsonConfigParserFactory{}
			if got := f.CreateSystemConfigParser(); !assert.Equal(t, c.want, got) {
				t.Errorf("CreateSystemConfigParser() = %v, want %v", got, c.want)
			}
		})
	}
}