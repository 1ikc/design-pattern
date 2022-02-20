package factory

type IConfigParser interface {
	Parse(data []byte)
}

type JsonConfigParser struct {}
func (j JsonConfigParser) Parse(data []byte) {
	panic("implement me")
}

type YamlConfigParser struct {}
func (y YamlConfigParser) Parse(data []byte) {
	panic("implement me")
}

func NewConfigParser(t string) IConfigParser {
	switch t {
	case "json":
		return JsonConfigParser{}
	case "yaml":
		return YamlConfigParser{}
	}

	return nil
}