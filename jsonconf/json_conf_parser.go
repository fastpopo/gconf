package jsonconf

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

type _JsonConfParser struct {
	dataMap      map[string]interface{}
	rootPath     string
	keyDelimiter string
}

func NewJsonConfParser(rootPath string, keyDelimiter string) *_JsonConfParser {
	return &_JsonConfParser{
		dataMap:      make(map[string]interface{}),
		rootPath:     rootPath,
		keyDelimiter: keyDelimiter,
	}
}

func (p *_JsonConfParser) Parse(stream []byte) error {
	if stream == nil || len(stream) == 0 {
		return errors.New("[_JsonConfParser::Parse] invalid null argument: stream")
	}

	var data interface{}

	if err := json.Unmarshal(stream, &data); err != nil {
		return err
	}

	p.parse(data, p.rootPath)

	return nil
}

func (p *_JsonConfParser) parse(value interface{}, path string) {
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

func (p *_JsonConfParser) parseArray(raw interface{}, parentKey string) {
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

func (p *_JsonConfParser) parseMap(raw interface{}, parentKey string) {

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

func (p *_JsonConfParser) GetDataMap() map[string]interface{} {
	return p.dataMap
}
