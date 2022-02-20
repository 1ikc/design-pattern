package abstract_factory

type IRuleConfigParser interface {
	ParseRule(data []byte)
}

type JsonRuleConfigParser struct {}
func (j JsonRuleConfigParser) ParseRule(data []byte) {
	panic("implement me")
}

type ISystemConfigParser interface {
	ParseSystem(data []byte)
}

type JsonSystemConfigParser struct {}
func (j JsonSystemConfigParser) ParseSystem(data []byte) {
	panic("implement me")
}

type JsonConfigParserFactory struct {}
func (j JsonConfigParserFactory) CreateRuleConfigParser() IRuleConfigParser {
	return JsonRuleConfigParser{}
}
func (j JsonConfigParserFactory) CreateSystemConfigParser() ISystemConfigParser {
	return JsonSystemConfigParser{}
}