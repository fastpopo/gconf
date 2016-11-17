package yamlconf

import (
	"errors"
	"fmt"
	"reflect"

	"gopkg.in/yaml.v2"
)

type _YamlConfParser struct {
	dataMap      map[string]interface{}
	rootPath     string
	keyDelimiter string
}

func NewYamlConfParser(rootPath string, keyDelimiter string) *_YamlConfParser {
	return &_YamlConfParser{
		dataMap:      make(map[string]interface{}),
		rootPath:     rootPath,
		keyDelimiter: keyDelimiter,
	}
}

func (p *_YamlConfParser) Parse(stream []byte) error {
	if stream == nil || len(stream) == 0 {
		return errors.New("[_YamlConfParser::Parse] invalid null argument: stream")
	}

	var data interface{}

	if err := yaml.Unmarshal(stream, &data); err != nil {
		return err
	}

	p.parse(data, p.rootPath)

	return nil
}

func (p *_YamlConfParser) parse(value interface{}, path string) {
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

func (p *_YamlConfParser) parseArray(raw interface{}, parentKey string) {
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

func (p *_YamlConfParser) parseMap(raw interface{}, parentKey string) {
	if raw == nil {
		return
	}

	data, ok := raw.(map[interface{}]interface{})

	if !ok {
		return
	}

	var path = parentKey
	if path != "" {
		path = path + p.keyDelimiter
	}

	for k, v := range data {
		newPath := path + fmt.Sprint(k)
		p.parse(v, newPath)
	}
}

func (p *_YamlConfParser) GetDataMap() map[string]interface{} {
	return p.dataMap
}
