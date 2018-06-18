package gconf

import (
	"errors"
	"encoding/json"
	"reflect"
	"fmt"
)

type JsonConfSource struct {
	jsonMessage []byte
}

func NewJsonConfSource(jsonMessage []byte) *JsonConfSource {
	return &JsonConfSource{
		jsonMessage: jsonMessage,
	}
}

func (s *JsonConfSource) Build(confBuilder ConfBuilder) ConfProvider {
	return NewConfProvider(s)
}

func (s *JsonConfSource) Load() (map[string]interface{}, error) {
	parser := newJsonConfParser(RootPath, KeyDelimiter)
	err := parser.Parse(s.jsonMessage)

	if err != nil {
		return nil, err
	}

	return parser.GetDataMap(), nil
}

type jsonConfParser struct {
	dataMap      map[string]interface{}
	rootPath     string
	keyDelimiter string
}

func newJsonConfParser(rootPath string, keyDelimiter string) *jsonConfParser {
	return &jsonConfParser{
		dataMap:      make(map[string]interface{}),
		rootPath:     rootPath,
		keyDelimiter: keyDelimiter,
	}
}

func (p *jsonConfParser) Parse(stream []byte) error {
	if stream == nil || len(stream) == 0 {
		return errors.New("[jsonConfParser::Parse] invalid null argument: stream")
	}

	var data interface{}

	if err := json.Unmarshal(stream, &data); err != nil {
		return err
	}

	p.parse(data, p.rootPath)

	return nil
}

func (p *jsonConfParser) parse(value interface{}, path string) {
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

func (p *jsonConfParser) parseArray(raw interface{}, parentKey string) {
	if raw == nil {
		return
	}

	data, ok := raw.([]interface{})

	if !ok {
		return
	}

	var path = parentKey

	if path != "" {
		path = path + p.keyDelimiter
	}

	for idx, v := range data {
		newPath := path + fmt.Sprint(idx)
		p.parse(v, newPath)
	}
}

func (p *jsonConfParser) parseMap(raw interface{}, parentKey string) {

	if raw == nil {
		return
	}

	data, ok := raw.(map[string]interface{})

	if !ok {
		return
	}

	var path = parentKey

	if path != "" {
		path = path + p.keyDelimiter
	}

	for k, v := range data {
		newPath := path + k
		p.parse(v, newPath)
	}
}

func (p *jsonConfParser) GetDataMap() map[string]interface{} {
	return p.dataMap
}
