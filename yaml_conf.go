package gconf

import (
	"errors"
	"fmt"
	"reflect"

	"gopkg.in/yaml.v2"
)

type YamlConfSource struct {
	yamlMessage []byte
}

func NewYamlConfSource(yamlMessage []byte) *YamlConfSource {
	return &YamlConfSource{
		yamlMessage: yamlMessage,
	}
}

func (s *YamlConfSource) Build(builder ConfBuilder) ConfProvider {
	return NewConfProvider(s)
}

func (s *YamlConfSource) Load() (map[string]interface{}, error) {
	parser := newYamlConfParser(RootPath, PathDelimiter)
	err := parser.Parse(s.yamlMessage)

	if err != nil {
		return nil, err
	}

	return parser.GetDataMap(), nil
}

type yamlConfParser struct {
	dataMap      map[string]interface{}
	rootPath     string
	keyDelimiter string
}

func newYamlConfParser(rootPath string, keyDelimiter string) *yamlConfParser {
	return &yamlConfParser{
		dataMap:      make(map[string]interface{}),
		rootPath:     rootPath,
		keyDelimiter: keyDelimiter,
	}
}

func (p *yamlConfParser) Parse(stream []byte) error {
	if stream == nil || len(stream) == 0 {
		return errors.New("[yamlConfParser::Parse] invalid null argument: stream")
	}

	var data interface{}

	if err := yaml.Unmarshal(stream, &data); err != nil {
		return err
	}

	p.parse(data, p.rootPath)

	return nil
}

func (p *yamlConfParser) parse(value interface{}, path string) {
	if value == nil {
		return
	}

	rv := reflect.ValueOf(value)

	switch rv.Kind() {
	case reflect.Map:
		p.parseMap(value, path)
		break
	case reflect.Slice:
		p.parseArray(value, path)
		break
	case reflect.Array:
		p.parseArray(value, path)
		break
	default:
		p.dataMap[path] = value
		break
	}
}

func (p *yamlConfParser) parseArray(raw interface{}, parentKey string) {
	if raw == nil {
		return
	}

	data, ok := raw.([]interface{})

	if !ok {
		return
	}

	for idx, v := range data {
		newPath := PathCombine(parentKey, ArrayDelimiter+ fmt.Sprint(idx))
		p.parse(v, newPath)
	}
}

func (p *yamlConfParser) parseMap(raw interface{}, parentKey string) {
	if raw == nil {
		return
	}

	data, ok := raw.(map[interface{}]interface{})

	if !ok {
		return
	}

	for k, v := range data {
		newPath := PathCombine(parentKey, fmt.Sprint(k))
		p.parse(v, newPath)
	}
}

func (p *yamlConfParser) GetDataMap() map[string]interface{} {
	return p.dataMap
}
