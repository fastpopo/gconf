package gconf

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/BurntSushi/toml"
)

type TomlConfSource struct {
	tomlMessage []byte
}

func NewTomlConfSource(tomlMessage []byte) *TomlConfSource {
	return &TomlConfSource{
		tomlMessage: tomlMessage,
	}
}

func (s *TomlConfSource) Build(confBuilder ConfBuilder) ConfProvider {
	return NewConfProvider(s)
}

func (s *TomlConfSource) Load() (map[string]interface{}, error) {
	parser := newJsonConfParser(RootPath, PathDelimiter)
	err := parser.Parse(s.tomlMessage)

	if err != nil {
		return nil, err
	}

	return parser.GetDataMap(), nil
}

type tomlConfParser struct {
	dataMap      map[string]interface{}
	rootPath     string
	keyDelimiter string
}

func newTomlConfParser(rootPath string, keyDelimiter string) *tomlConfParser {
	return &tomlConfParser{
		dataMap:      make(map[string]interface{}),
		rootPath:     rootPath,
		keyDelimiter: keyDelimiter,
	}
}

func (p *tomlConfParser) Parse(stream []byte) error {
	if stream == nil || len(stream) == 0 {
		return errors.New("[jsonConfParser::Parse] invalid null argument: stream")
	}

	var data interface{}

	if _, err := toml.Decode(string(stream), &data); err != nil {
		return err
	}

	p.parse(data, p.rootPath)

	return nil
}

func (p *tomlConfParser) parse(value interface{}, path string) {
	if value == nil {
		return
	}

	rv := reflect.ValueOf(value)

	switch rv.Kind() {
	case reflect.Map:
		p.parseMap(value, path)
		break
	case reflect.Slice, reflect.Array:
		p.parseArray(value, path)
		break
	default:
		p.dataMap[path] = value
		break
	}
}

func (p *tomlConfParser) parseArray(raw interface{}, parentKey string) {
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

func (p *tomlConfParser) parseMap(raw interface{}, parentKey string) {

	if raw == nil {
		return
	}

	data, ok := raw.(map[string]interface{})

	if !ok {
		return
	}

	for k, v := range data {
		newPath := PathCombine(parentKey, k)
		p.parse(v, newPath)
	}
}

func (p *tomlConfParser) GetDataMap() map[string]interface{} {
	return p.dataMap
}
