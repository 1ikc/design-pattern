package factory_method

import factory "github.com/1ikc/design-pattern/02_factory/021_simple_factory"

type IConfigParserFactory interface {
	CreateParser() factory.IConfigParser
}

type JsonConfigParserFactory struct {}
func (j JsonConfigParserFactory) CreateParser() factory.IConfigParser {
	return factory.JsonConfigParser{}
}

type YamlConfigParserFactory struct {}
func (y YamlConfigParserFactory) CreateParser() factory.IConfigParser {
	return factory.YamlConfigParser{}
}

func NewConfigParserFactory(t string) IConfigParserFactory {
	switch t {
	case "json":
		return JsonConfigParserFactory{}
	case "yaml":
		return YamlConfigParserFactory{}
	}

	return nil
}